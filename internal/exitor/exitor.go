// Package exitor provides an interface for exiting the program
package exitor

// IExitor is the simple interface for os.Exit
type IExitor interface {
	Exit(int)
}
