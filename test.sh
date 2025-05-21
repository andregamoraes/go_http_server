#!/bin/bash

# Start the server in the background
./run.sh &
SERVER_PID=$!

# Wait for the server to start
sleep 1

# Run the tests with TEST_DIR environment variable
TEST_DIR="$TEST_DIR" go test -v ./app/tests/...

# Kill the server
kill $SERVER_PID

# Clean up the temporary directory
rm -rf "$TEST_DIR" 