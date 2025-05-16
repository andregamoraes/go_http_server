package handler

import (
	"net"
	"fmt"
	"os"
	"path/filepath"
)

func Handle200(conn net.Conn) {
	response := "HTTP/1.1 200 OK\r\nConnection: close\r\n\r\n"
	conn.Write([]byte(response))
}

func Handle404(conn net.Conn) {
	response := "HTTP/1.1 404 Not Found\r\nConnection: close\r\n\r\n"
	conn.Write([]byte(response))
}

func HandleEcho(conn net.Conn, content string) {
	contentLength := len(content)
	response := fmt.Sprintf(
		"HTTP/1.1 200 OK\r\n" +
		"Content-Type: text/plain\r\n" +
		"Content-Length: %d\r\n" +
		"Connection: close\r\n" +
		"\r\n" +
		"%s",
		contentLength,
		content,
	)
	conn.Write([]byte(response))
}

func HandleUserAgent(conn net.Conn, userAgent string) {
	contentLength := len(userAgent)
	response := fmt.Sprintf(
		"HTTP/1.1 200 OK\r\n" +
		"Content-Type: text/plain\r\n" +
		"Content-Length: %d\r\n" +
		"Connection: close\r\n" +
		"\r\n" +
		"%s",
		contentLength,
		userAgent,
	)
	conn.Write([]byte(response))
}

func HandleFile(conn net.Conn, directory string, filename string) {
	// Construct the full file path and ensure it's within the directory
	fullPath := filepath.Join(directory, filename)
	if !filepath.HasPrefix(fullPath, directory) {
		Handle404(conn)
		return
	}

	// Try to open and read the file
	content, err := os.ReadFile(fullPath)
	if err != nil {
		Handle404(conn)
		return
	}

	// Send the file contents with appropriate headers
	response := fmt.Sprintf(
		"HTTP/1.1 200 OK\r\n" +
		"Content-Type: application/octet-stream\r\n" +
		"Content-Length: %d\r\n" +
		"Connection: close\r\n" +
		"\r\n",
		len(content),
	)
	
	conn.Write([]byte(response))
	conn.Write(content)
}

func HandleFilePost(conn net.Conn, directory string, filename string, content []byte) {
	// Construct the full file path and ensure it's within the directory
	fullPath := filepath.Join(directory, filename)
	if !filepath.HasPrefix(fullPath, directory) {
		Handle404(conn)
		return
	}

	// Write the file
	err := os.WriteFile(fullPath, content, 0644)
	if err != nil {
		Handle404(conn)
		return
	}

	// Send 201 Created response
	response := "HTTP/1.1 201 Created\r\nConnection: close\r\n\r\n"
	conn.Write([]byte(response))
}
