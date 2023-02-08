#!/usr/bin/env bash

brew list golangci-lint || brew install golangci-lint
brew list staticcheck || brew install staticcheck
go install github.com/fzipp/gocyclo/cmd/gocyclo@latest
