package exitor

import "os"

// Ensure defaultExitor implements IExitor interface
var _ IExitor = (*defaultExitor)(nil)

type defaultExitor struct{}

// Exit calls os.Exit
func (e *defaultExitor) Exit(code int) {
	os.Exit(code)
}

func New() IExitor {
	return &defaultExitor{}
}
