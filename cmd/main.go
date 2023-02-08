package main

import (
	"fmt"
	"github.com/andrewmolyuk/basalt/internal/log"
)

var (
	version = "0.1.2.DEVELOPMENT"
	commit  = "UNKNOWN"
)

func main() {
	log.Info("Basalt %s\n", fmt.Sprintf("%s (git: %s)", version, commit[:7]))
}
