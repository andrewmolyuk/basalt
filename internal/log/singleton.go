package log

import (
	"github.com/andrewmolyuk/basalt/internal/exitor"
	"os"
)

var instanceLogger ILogger

func init() {
	debugMode := os.Getenv("BASALT_DEBUG") == "true"
	secrets := getSecrets()
	instanceLogger = New(debugMode, secrets, exitor.New())
	Info("Logger initialized successfully (debug mode: %v, secrets: %v)", debugMode, len(secrets))
}

func getSecrets() []string {
	var secrets []string

	var secretNames = []string{
		"",
	}

	for _, name := range secretNames {
		if value := os.Getenv(name); value != "" {
			secrets = append(secrets, value)
		}
	}

	return secrets
}

func getLogger() ILogger {
	return instanceLogger
}

// Debug prints debug message if debug mode is enabled
func Debug(args ...interface{}) {
	getLogger().Debug(args...)
}

// Info prints info message
func Info(args ...interface{}) {
	getLogger().Info(args...)
}

// Warn prints warning message
func Warn(args ...interface{}) {
	getLogger().Warn(args...)
}

// Error prints error message and exits with code 1
func Error(args ...interface{}) {
	getLogger().Error(args...)
}
