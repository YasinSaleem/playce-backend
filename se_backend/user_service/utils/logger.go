package utils

import (
	"fmt"
	"log"
	"os"
	"time"
)

// Logger is a custom logger for the application
type Logger struct {
	InfoLogger  *log.Logger
	ErrorLogger *log.Logger
}

// NewLogger creates a new logger instance
func NewLogger() *Logger {
	return &Logger{
		InfoLogger:  log.New(os.Stdout, "[INFO] ", log.Ldate|log.Ltime),
		ErrorLogger: log.New(os.Stderr, "[ERROR] ", log.Ldate|log.Ltime),
	}
}

// Info logs an info message
func (l *Logger) Info(format string, v ...interface{}) {
	l.InfoLogger.Printf(format, v...)
}

// Error logs an error message
func (l *Logger) Error(format string, v ...interface{}) {
	l.ErrorLogger.Printf(format, v...)
}

// LogRequest logs information about an HTTP request
func (l *Logger) LogRequest(method, path, clientIP, userAgent string, statusCode int, latency time.Duration) {
	message := fmt.Sprintf("Request: %s %s | Status: %d | IP: %s | User-Agent: %s | Latency: %v",
		method, path, statusCode, clientIP, userAgent, latency)
	l.Info("%s", message)
}
