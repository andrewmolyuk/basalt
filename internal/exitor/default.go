// Package exitor provides an interface for exiting the program
package exitor

import "os"

// Ensure defaultExitor implements IExitor interface
var _ IExitor = (*defaultExitor)(nil)

type defaultExitor struct{}

// Exit provides a simple wrapper for os.Exit
func (e *defaultExitor) Exit(code int) {
	os.Exit(code)
}

// New creates a new instance of exitor
func New() IExitor {
	return &defaultExitor{}
}
