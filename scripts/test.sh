#!/usr/bin/env bash

go test -timeout=60s -race -coverprofile=.coverage.out ./...