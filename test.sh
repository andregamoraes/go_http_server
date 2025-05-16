#!/bin/bash

# Start the server in the background
./run.sh &
SERVER_PID=$!

# Wait for the server to start
sleep 1

# Run the tests
go test -v ./app/tests/...

# Kill the server
kill $SERVER_PID 