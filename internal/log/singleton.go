// Package log provides a simple logger implementation with debug mode support and secrets masking
package log

import (
	"github.com/andrewmolyuk/basalt/internal/exitor"
)

var instanceLogger ILogger

func getLogger() ILogger {
	return instanceLogger
}

// Init initializes default logger
func Init(debugMode bool, secrets []string) {
	instanceLogger = New(debugMode, secrets, exitor.New())
	Info("Logger initialized successfully (debug mode: %v, secrets: %v)", debugMode, len(secrets))
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
