package middleware

import (
	"strconv"
	"time"

	"user-service/pkg/metrics"

	"github.com/labstack/echo/v4"
)

// MetricsMiddleware provides comprehensive metrics collection
func MetricsMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		start := time.Now()
		method := c.Request().Method
		endpoint := c.Path()

		// Track in-flight requests
		metrics.HttpRequestsInFlight.Inc()
		defer metrics.HttpRequestsInFlight.Dec()

		// Process request
		err := next(c)

		// Calculate duration
		duration := time.Since(start).Seconds()
		status := strconv.Itoa(c.Response().Status)

		// Record metrics
		metrics.HttpRequestsTotal.WithLabelValues(method, endpoint, status).Inc()
		metrics.HttpRequestDuration.WithLabelValues(method, endpoint, status).Observe(duration)

		return err
	}
}
