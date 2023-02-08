package main

import (
	"fmt"
	"os"
)

var (
	version = "0.1.2.DEVELOPMENT"
	commit  = "UNKNOWN"
)

func main() {
	fmt.Printf("Basalt %s\n", fmt.Sprintf("%s (git: %s)", version, commit[:7]))
	os.Exit(0)
}
