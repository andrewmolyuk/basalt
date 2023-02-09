// Package app provides the application entry point and the main application logic implementation
package app

import (
	"fmt"
	"github.com/jessevdk/go-flags"
	"os"
)

// Basalt is the main application structure that holds all the configuration and the application state
type Basalt struct {
	Debug       bool `short:"d" long:"debug" description:"Debug mode"`
	ShowVersion bool `short:"v" long:"version" description:"Show Basalt version info"`
	BuildInfo   BuildInfo
}

// Run is the main process where the application is running
func (b *Basalt) Run() {

	parser := flags.NewParser(b, flags.Default)
	parser.ShortDescription = "Basalt is command line for managing your bastion and AWS infrastructure access."
	parser.LongDescription = "Basalt is a command-line tool for managing user access to AWS infrastructure through bastion, distributing and rotating SSH keys."

	if _, err := parser.Parse(); err != nil {
		code := 1
		if fe, ok := err.(*flags.Error); ok {
			if fe.Type == flags.ErrHelp {
				code = 0
			}
		}
		os.Exit(code)
	}

	if b.ShowVersion {
		fmt.Printf("Basalt %s\n", b.BuildInfo)
	}

	// TODO: implement the main logic here
}
