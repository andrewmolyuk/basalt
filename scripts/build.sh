#!/usr/bin/env bash

go build -ldflags "-s -w" -o ./bin/basalt ./cmd/main.go
