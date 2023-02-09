// Package app provides the application entry point and the main application logic implementation
package app

// BuildInfo contains a minimal build information embedded into binary during version release
type BuildInfo struct {
	Version string
	Commit  string
}

// String returns a string representation of BuildInfo
func (i BuildInfo) String() string {
	return i.Version + " (" + i.Commit + ")"
}
