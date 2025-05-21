package router

import (
	"fmt"
	"net"
	"go-http-server/app/handler"
	"bufio"
	"strings"
	"io"
	"net/url"
)

func parseHeaders(reader *bufio.Reader) map[string]string {
	headers := make(map[string]string)
	
	for {
		line, err := reader.ReadString('\n')
		if err != nil || line == "\r\n" {
			break
		}
		
		line = strings.TrimRight(line, "\r\n")
		parts := strings.SplitN(line, ": ", 2)
		if len(parts) == 2 {
			headers[parts[0]] = parts[1]
		}
	}
	
	return headers
}

func readBody(reader *bufio.Reader, headers map[string]string) []byte {
	contentLength := 0
	if lenStr, ok := headers["Content-Length"]; ok {
		fmt.Sscanf(lenStr, "%d", &contentLength)
	}

	if contentLength > 0 {
		body := make([]byte, contentLength)
		io.ReadFull(reader, body)
		return body
	}
	return nil
}

func shouldCloseConnection(headers map[string]string) bool {
	if connection, exists := headers["Connection"]; exists {
		return strings.ToLower(connection) == "close"
	}
	return false
}

func HandleConnection(conn net.Conn, directory string) {
	defer conn.Close()
	reader := bufio.NewReader(conn)

	for {
		requestLine, err := reader.ReadString('\n')
		if err != nil {
			return
		}

		parts := strings.Split(strings.TrimSpace(requestLine), " ")
		if len(parts) < 2 {
			return
		}

		method := parts[0]
		path := parts[1]

		headers := parseHeaders(reader)

		body := readBody(reader, headers)

		handleRequest(conn, method, path, headers, body, directory)

		if shouldCloseConnection(headers) {
			return
		}
	}
}

func handleRequest(conn net.Conn, method string, path string, headers map[string]string, body []byte, directory string) {
	if path == "/" {
		handler.Handle200(conn)
	} else if path == "/user-agent" {
		userAgent := headers["User-Agent"]
		handler.HandleUserAgent(conn, userAgent)
	} else if strings.HasPrefix(path, "/echo/") {
		echoStr := strings.TrimPrefix(path, "/echo/")
		if echoStr == "" {
			handler.Handle404(conn)
			return
		}
		// URL decode the echo string
		decodedStr, err := url.QueryUnescape(echoStr)
		if err != nil {
			handler.Handle404(conn)
			return
		}
		handler.HandleEcho(conn, decodedStr)
	} else if strings.HasPrefix(path, "/files/") {
		filename := strings.TrimPrefix(path, "/files/")
		if method == "POST" {
			handler.HandleFilePost(conn, directory, filename, body)
		} else {
			handler.HandleFile(conn, directory, filename)
		}
	} else {
		handler.Handle404(conn)
	}
}
