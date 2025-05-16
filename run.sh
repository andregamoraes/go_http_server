#!/bin/bash

# Build the server
go build -o http-server app/main.go

# Run the server with the provided directory
# If no directory is provided, use /tmp as default
DIR=${1:-/tmp}
./http-server --directory "$DIR" 