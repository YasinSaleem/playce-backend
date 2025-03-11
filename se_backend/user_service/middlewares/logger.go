package middlewares

import (
	"time"
	"user_service/utils"

	"github.com/labstack/echo/v4"
)

var logger = utils.NewLogger()

// LoggerMiddleware logs information about each request
func LoggerMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Start timer
			startTime := time.Now()

			// Process request
			err := next(c)

			// Calculate latency
			latency := time.Since(startTime)

			// Log request details
			logger.LogRequest(
				c.Request().Method,
				c.Path(),
				c.RealIP(),
				c.Request().UserAgent(),
				c.Response().Status,
				latency,
			)

			// Log any errors if they occurred
			if err != nil {
				logger.Error("%v", err)
			}

			return err
		}
	}
}
