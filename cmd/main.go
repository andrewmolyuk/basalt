// Package main provides the entry point for the basalt command line tool.
package main

import (
	"github.com/andrewmolyuk/basalt/internal/app"
)

var (
	version = "0.1.2.DEVELOPMENT"
	commit  = "UNKNOWN"
)

// main is the entry point for the basalt command line tool.
func main() {
	basaltApp := app.Basalt{
		BuildInfo: app.BuildInfo{
			Version: version,
			Commit:  commit,
		},
	}

	basaltApp.Run()
}
