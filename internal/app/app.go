// Package app provides the application entry point and the main application logic implementation
package app

import (
	"fmt"
	"github.com/andrewmolyuk/basalt/internal/log"
	"github.com/jessevdk/go-flags"
	"os"
)

// Basalt is the main application structure that holds all the configuration and the application state
type Basalt struct {
	Debug              bool   `short:"d" long:"debug" description:"Debug mode"`
	ShowVersion        bool   `short:"v" long:"version" description:"Show Basalt version info"`
	AwsAccessKeyId     string `short:"k" long:"key" description:"AWS Access Key ID"`
	AwsSecretAccessKey string `short:"s" long:"secret" description:"AWS Access Key Secret"`
	AwsRegion          string `short:"r" long:"region" description:"AWS Region" default:"us-east-1"`
	ConfigFile         string `short:"c" long:"config" description:"Path to the configuration file" default:"basalt.yml"`
	BuildInfo          BuildInfo
}

// Run is the main process where the application is running
func (b *Basalt) Run() {

	b.AwsAccessKeyId = os.Getenv("AWS_ACCESS_KEY_ID")
	b.AwsSecretAccessKey = os.Getenv("AWS_SECRET_ACCESS_KEY")
	b.AwsRegion = os.Getenv("AWS_DEFAULT_REGION")

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

	if b.AwsAccessKeyId == "" || b.AwsSecretAccessKey == "" {
		log.Error("AWS Access Key ID and AWS Access Key Secret must be provided")
	}

	log.Init(b.Debug, []string{b.AwsRegion, b.AwsAccessKeyId, b.AwsSecretAccessKey})

	// TODO: implement the main logic here
}
