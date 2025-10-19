package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// метрики
var (
	httpRequestsTotal = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "Total number of HTTP requests",
	}, []string{"endpoint", "method", "status"})

	httpRequestDuration = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "http_request_duration_seconds",
		Help:    "HTTP request duration in seconds",
		Buckets: prometheus.DefBuckets,
	}, []string{"endpoint", "method", "status"})

	httpRequestsInFlight = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "http_requests_in_flight",
		Help: "Current number of HTTP requests in flight",
	})

	usersCountTotal = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "users_count_total",
		Help: "Total number of users in the database",
	})

	databaseConnections = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "database_connections_total",
		Help: "Current number of database connections",
	})

	userOperationsTotal = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "user_operations_total",
		Help: "Total number of user operations",
	}, []string{"operation", "status"})
)

// User represents a user model
type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

func initDB() *sql.DB {
	connStr := getEnv("DATABASE_URL", "postgres://postgres:password@postgres:5432/userdb?sslmode=disable")
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	// Try to connect with retries
	for i := 0; i < 10; i++ {
		err = db.Ping()
		if err == nil {
			break
		}
		log.Printf("Database connection attempt %d failed: %v", i+1, err)
		time.Sleep(2 * time.Second)
	}

	if err != nil {
		log.Fatal("Failed to connect to database after retries:", err)
	}

	// Create users table if not exists
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			name VARCHAR(100) NOT NULL,
			email VARCHAR(100) UNIQUE NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)
	`)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("✅ Успешное подключение к PostgreSQL!")
	return db
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func updateDatabaseMetrics(db *sql.DB) {
	// Update users count
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM users").Scan(&count)
	if err != nil {
		log.Printf("Error counting users: %v", err)
	} else {
		usersCountTotal.Set(float64(count))
	}

	// Update database connections stats
	stats := db.Stats()
	databaseConnections.Set(float64(stats.OpenConnections))
}

func prometheusMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Start timer and increment in-flight requests
		start := time.Now()
		httpRequestsInFlight.Inc()
		defer httpRequestsInFlight.Dec()

		// Process request
		err := next(c)

		// Get endpoint name
		endpoint := c.Path()
		if endpoint == "" {
			endpoint = c.Request().URL.Path
		}

		// Record metrics
		status := fmt.Sprintf("%d", c.Response().Status)
		httpRequestsTotal.WithLabelValues(endpoint, c.Request().Method, status).Inc()
		httpRequestDuration.WithLabelValues(endpoint, c.Request().Method, status).Observe(time.Since(start).Seconds())

		return err
	}
}

func healthHandler(c echo.Context) error {
	userOperationsTotal.WithLabelValues("health_check", "200").Inc()
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":    "healthy",
		"database":  "connected",
		"timestamp": time.Now().UTC(),
	})
}

func getUsersHandler(c echo.Context) error {
	db := c.Get("db").(*sql.DB)
	
	rows, err := db.Query("SELECT id, name, email, created_at FROM users ORDER BY id")
	if err != nil {
		userOperationsTotal.WithLabelValues("get_users", "500").Inc()
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt)
		if err != nil {
			userOperationsTotal.WithLabelValues("get_users", "500").Inc()
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}
		users = append(users, user)
	}

	userOperationsTotal.WithLabelValues("get_users", "200").Inc()
	return c.JSON(http.StatusOK, map[string]interface{}{
		"total": len(users),
		"users": users,
	})
}

func createUserHandler(c echo.Context) error {
	db := c.Get("db").(*sql.DB)
	
	var user struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}
	
	if err := c.Bind(&user); err != nil {
		userOperationsTotal.WithLabelValues("create_user", "400").Inc()
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	_, err := db.Exec("INSERT INTO users (name, email) VALUES ($1, $2)", user.Name, user.Email)
	if err != nil {
		userOperationsTotal.WithLabelValues("create_user", "500").Inc()
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	userOperationsTotal.WithLabelValues("create_user", "201").Inc()
	return c.JSON(http.StatusCreated, map[string]string{"message": "User created successfully"})
}

func main() {
	// Initialize database
	db := initDB()
	defer db.Close()

	// Create Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(prometheusMiddleware)
	
	// Add database to context
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("db", db)
			return next(c)
		}
	})

	// Routes
	e.GET("/health", healthHandler)
	e.GET("/api/users", getUsersHandler)
	e.POST("/api/users", createUserHandler)
	e.GET("/metrics", echo.WrapHandler(promhttp.Handler()))

	// Start metrics updater
	go func() {
		for {
			updateDatabaseMetrics(db)
			time.Sleep(30 * time.Second)
		}
	}()

	// Start server
	address := getEnv("SERVER_ADDRESS", "0.0.0.0:8091")
	fmt.Printf("🚀 Запуск HTTP сервера на %s\n", address)
	fmt.Printf("📊 Метрики доступны по: http://%s/metrics\n", address)
	fmt.Printf("❤️  Health check: http://%s/health\n", address)
	
	if err := e.Start(address); err != nil {
		log.Fatal(err)
	}
}
