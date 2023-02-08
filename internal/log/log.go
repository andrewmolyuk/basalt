// Package log provides a simple logger implementation with debug mode support and secrets masking
package log

// ILogger is the simple logger interface
type ILogger interface {
	Debug(...interface{})
	Info(...interface{})
	Warn(...interface{})
	Error(...interface{})
}
