// Package logger provides a structured logging interface for Go applications.
//
// This package defines a Log interface with methods for logging messages at different levels:
// Info, Warn, Error, and Debug. It also includes an implementation of this interface, SlogAdapter,
// which uses the slog.Logger for logging.
//
// The logger package is designed to be flexible and easy to use, allowing developers to log messages
// with varying levels of severity. The package also includes a utility function, GetLevelFromEnv,
// which reads the LOG_LEVEL environment variable and returns the corresponding slog.Level.
//
// Example usage:
//
//	// Create a new logger instance
//	logger := logger.NewLogger()
//
//	// Log messages at different levels
//	logger.Info("This is an info message")
//	logger.Warn("This is a warning message")
//	logger.Error("This is an error message")
//	logger.Debug("This is a debug message")
//
// The logger package is intended to be used in applications where structured logging is required.
// It provides a consistent and easy-to-use interface for logging messages, making it easier to
// manage and analyze log data.
//
// For more information, visit the Go.pkg.dev documentation page for this package.
package logger

import (
	"log/slog"
	"os"
)

// Log is an interface that defines the methods for logging.
//
// The methods are used to log messages at different levels.
// The levels are Info, Warn, Error, and Debug.
//
// Example:
//
//	logger := logger.NewLogger()
//	logger.Info("This is an info message")
//	logger.Warn("This is a warning message")
//	logger.Error("This is an error message")
//	logger.Debug("This is a debug message")
type Log interface {
	Info(msg string, args ...any)
	Warn(msg string, args ...any)
	Error(msg string, args ...any)
	Debug(msg string, args ...any)
}

// SlogAdapter is a struct that implements the Log interface.
//
// It contains a slog.Logger instance.
// The slog.Logger instance is used for logging.
type SlogAdapter struct {
	Logger Log
}

// GetLevelFromEnv
//
// This function reads the LOG_LEVEL environment variable and returns the corresponding slog.Level.
// If the LOG_LEVEL environment variable is not set, it returns slog.LevelInfo.
//
// Returns:
//   - The slog.Level corresponding to the LOG_LEVEL environment variable.
//
// Example:
//
//	level := GetLevelFromEnv()
func GetLevelFromEnv() slog.Level {
	level := os.Getenv("LOG_LEVEL")

	switch level {
	case "DEBUG":
		return slog.LevelDebug
	case "INFO":
		return slog.LevelInfo
	case "WARN":
		return slog.LevelWarn
	case "ERROR":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}

// NewLogger creates a new SlogAdapter instance with a slog.Logger.
//
// It reads the LOG_FORMAT environment variable to determine the log format (text or json).
// If the LOG_FORMAT environment variable is not set, it defaults to text.
//
// Returns:
//   - A new SlogAdapter instance with a slog.Logger.
//
// Example:
//
//	logger := logger.NewLogger()
//	logger.Info("This is an info message")
//	logger.Warn("This is a warning message")
//	logger.Error("This is an error message")
//	logger.Debug("This is a debug message")
func NewLogger() *SlogAdapter {
	logFormat := os.Getenv("LOG_FORMAT")

	var handler slog.Handler
	if logFormat == "json" {
		handler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: GetLevelFromEnv(),
		})
	} else {
		handler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: GetLevelFromEnv(),
		})
	}

	logger := slog.New(handler)

	return &SlogAdapter{
		Logger: logger,
	}
}
