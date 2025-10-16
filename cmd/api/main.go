package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
)

type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

var db *sql.DB

func initDB() error {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		"postgres", "5432", "postgres", "password", "userdb")

	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		return err
	}

	// Create table if not exists
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			name VARCHAR(100) NOT NULL,
			email VARCHAR(100) UNIQUE NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)
	`)
	if err != nil {
		return fmt.Errorf("failed to create table: %v", err)
	}

	return db.Ping()
}

func healthCheck(c echo.Context) error {
	if db != nil {
		if err := db.Ping(); err != nil {
			return c.JSON(http.StatusServiceUnavailable, map[string]string{
				"status":  "unhealthy", 
				"message": "Database connection failed",
			})
		}
		return c.JSON(http.StatusOK, map[string]string{
			"status":    "healthy",
			"timestamp": time.Now().Format(time.RFC3339),
			"database":  "connected",
		})
	}
	
	return c.JSON(http.StatusOK, map[string]string{
		"status":    "healthy",
		"timestamp": time.Now().Format(time.RFC3339),
		"database":  "mock_mode",
	})
}

func getMetrics(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Metrics endpoint - Prometheus metrics will be added later",
	})
}

func getUsers(c echo.Context) error {
	if db == nil {
		// Mock данные если БД не подключена
		users := []User{
			{ID: 1, Name: "John Doe", Email: "john@example.com", CreatedAt: time.Now()},
			{ID: 2, Name: "Jane Smith", Email: "jane@example.com", CreatedAt: time.Now()},
			{ID: 3, Name: "Bob Johnson", Email: "bob@example.com", CreatedAt: time.Now()},
		}
		return c.JSON(http.StatusOK, users)
	}

	// Реальные данные из БД
	rows, err := db.Query("SELECT id, name, email, created_at FROM users ORDER BY created_at DESC")
	if err != nil {
		log.Printf("ERROR: Failed to query users: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to fetch users",
		})
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt); err != nil {
			log.Printf("ERROR: Failed to scan user: %v", err)
			continue
		}
		users = append(users, user)
	}

	if len(users) == 0 {
		return c.JSON(http.StatusOK, []User{})
	}

	return c.JSON(http.StatusOK, users)
}

func createUser(c echo.Context) error {
	var user struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request body",
		})
	}

	if user.Name == "" || user.Email == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Name and email are required",
		})
	}

	if db == nil {
		// Mock response если БД не подключена
		return c.JSON(http.StatusCreated, map[string]interface{}{
			"id":      time.Now().Unix(),
			"name":    user.Name,
			"email":   user.Email,
			"message": "User created successfully (mock - no DB connection)",
		})
	}

	// Сохраняем в реальную БД
	var id int
	err := db.QueryRow(
		"INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id",
		user.Name, user.Email,
	).Scan(&id)

	if err != nil {
		log.Printf("ERROR: Failed to create user: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to create user - may be duplicate email",
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"id":    id,
		"name":  user.Name,
		"email": user.Email,
		"message": "User created successfully in database",
	})
}

func main() {
	// Инициализация базы данных
	if err := initDB(); err != nil {
		log.Printf("WARNING: Database connection failed: %v", err)
		log.Printf("INFO: Starting with mock data mode")
	} else {
		log.Printf("INFO: Database connected successfully")
	}

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// Routes
	e.GET("/health", healthCheck)
	e.GET("/metrics", getMetrics) 
	e.GET("/api/users", getUsers)
	e.POST("/api/users", createUser)

	// Запуск сервера
	port := "8090"

	log.Printf("🚀 Starting User Service on port %s", port)
	log.Printf("📊 Health check: http://localhost:%s/health", port)
	log.Printf("👥 API Users: http://localhost:%s/api/users", port)
	log.Printf("📈 Metrics: http://localhost:%s/metrics", port)
	
	if err := e.Start(":" + port); err != nil {
		log.Fatalf("❌ Failed to start server: %v", err)
	}
}
