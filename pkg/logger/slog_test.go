package logger

import (
	"log/slog"
	"os"
	"reflect"
	"testing"
)

func TestGetLogLevelFromEnv(t *testing.T) {
	// Define test cases for different environment variable settings
	testCases := []struct {
		envValue string
		expected slog.Level
	}{
		{"DEBUG", slog.LevelDebug},
		{"INFO", slog.LevelInfo},
		{"WARN", slog.LevelWarn},
		{"ERROR", slog.LevelError},
		{"", slog.LevelInfo}, // Default case
	}

	for _, tc := range testCases {
		t.Run(tc.envValue, func(t *testing.T) {
			// Set the environment variable for the test case
			_ = os.Setenv("LOG_LEVEL", tc.envValue)
			defer func() {
				_ = os.Unsetenv("LOG_LEVEL")
			}()

			// Call the function under test
			level := GetLevelFromEnv()

			// Assert that the result matches the expected outcome
			if level != tc.expected {
				t.Errorf("Expected log level %v for env value '%s', but got %v", tc.expected, tc.envValue, level)
			}
		})
	}
}

func TestNewLogger(t *testing.T) {
	// Define test cases for different log formats
	testCases := []struct {
		format   string
		expected string
	}{
		{"json", "json"},
		{"", "text"}, // Default case should create a text logger
	}

	for _, tc := range testCases {
		t.Run(tc.format, func(t *testing.T) {
			// Set the environment variable for the test case
			_ = os.Setenv("LOG_FORMAT", tc.format)
			defer func() {
				_ = os.Unsetenv("LOG_FORMAT")
			}()

			// Create the logger
			logger := NewLogger()

			// Type assert logger.Logger to *slog.Logger to access its Handler for assertion
			// NOTE: This might require making the test part of the same package or
			// adjusting visibility to inspect the handler's type or properties.
			slogLogger, ok := logger.Logger.(*slog.Logger)
			if !ok {
				t.Fatalf("Expected SlogAdapter.Logger to be *slog.Logger, but it's not")
			}

			// Check if the logger's handler is of the expected format.
			if tc.expected == "json" {
				if reflect.TypeOf(slogLogger.Handler()) != reflect.TypeOf(slog.NewJSONHandler(nil, nil)) {
					t.Errorf("Expected logger format to be %s, but it was not", tc.expected)
				}
			} else if tc.expected == "text" {
				if reflect.TypeOf(slogLogger.Handler()) != reflect.TypeOf(slog.NewTextHandler(nil, nil)) {
					t.Errorf("Expected logger format to be %s, but it was not", tc.expected)
				}
			}
		})
	}
}
