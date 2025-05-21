# Go HTTP/1.1 Server Implementation

A lightweight HTTP/1.1 server implemented in Go, featuring support for multiple endpoints, file operations, and persistent connections.

## Features

- **Basic HTTP/1.1 Protocol Support**
  - Handles GET and POST methods
  - Supports persistent connections
  - Proper header handling
  - Status code responses (200, 201, 404)

- **Multiple Endpoints**
  - `/` - Root endpoint returning 200 OK
  - `/echo/{string}` - Returns the provided string with proper headers
  - `/user-agent` - Returns the client's User-Agent
  - `/files/{filename}` - File operations (GET and POST)

- **File Operations**
  - GET `/files/{filename}` - Retrieve file contents
  - POST `/files/{filename}` - Create new files
  - Directory-based file storage
  - Security measures against path traversal

## Project Structure

```
app/
├── main.go           # Server initialization and configuration
├── handler/
│   └── handler.go    # HTTP response handlers
└── router/
    └── router.go     # Request routing and processing
└── tests/
    └── echo_test.go      # Tests for the echo endpoint
    └── files_test.go     # Tests for file operations
    └── user_agent_test.go # Tests for the user-agent endpoint
```

## Requirements

- Go 1.24 or higher
- No external dependencies required

## Installation

1. Clone the repository:
   ```bash
   git clone <your-repo-url>
   cd <repo-directory>
   ```

2. Build the server:
   ```bash
   go build -o http-server app/main.go
   ```

## Usage

Run the server with a specified directory for file operations:

```bash
./http-server --directory /path/to/files
```

The server will start listening on port 4221.

### Example Requests

1. Echo Endpoint:
   ```bash
   curl -v http://localhost:4221/echo/hello
   ```

2. User-Agent Endpoint:
   ```bash
   curl -v http://localhost:4221/user-agent
   ```

3. File Operations:
   ```bash
   # Create a file
   curl -v -X POST -H "Content-Type: application/octet-stream" --data "Hello World" http://localhost:4221/files/hello.txt

   # Retrieve a file
   curl -v http://localhost:4221/files/hello.txt
   ```

## Testing

The project includes a comprehensive test suite located in the `app/tests` directory. Tests cover all endpoints and functionality:

- Echo endpoint tests
- User-Agent endpoint tests
- File operations tests (GET and POST)

### Running Tests

You can run the tests in two ways:

1. Using the provided test script (recommended):
   ```bash
   ./test.sh
   ```
   This script will:
   - Start the server in the background
   - Run all tests
   - Kill the server after tests complete

2. Running tests directly with Go:
   ```bash
   go test -v ./app/tests/...
   ```
   Note: When running tests directly, make sure the server is running with the correct directory:
   ```bash
   ./http-server --directory /tmp
   ```

## Implementation Details

### HTTP/1.1 Features

- **Persistent Connections**: Supports multiple requests over the same TCP connection
- **Connection Header**: Proper handling of `Connection: close` header
- **Content Headers**: Correct implementation of Content-Type and Content-Length

### Security

- Path traversal prevention in file operations
- Proper request parsing and validation
- Safe file handling

## Contributing

Feel free to submit issues and enhancement requests!

## License

[Your chosen license]

## Acknowledgments

This project was initially developed as part of the CodeCrafters "Build Your Own HTTP Server" challenge.
