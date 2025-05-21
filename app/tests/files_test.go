package tests

import (
	"bytes"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"testing"
)

func TestFilesEndpoint(t *testing.T) {
	// Use /tmp directory
	tempDir := "/tmp"

	// Test file creation (POST)
	t.Run("POST new file", func(t *testing.T) {
		content := []byte("Hello, World!")
		req, err := http.NewRequest(
			"POST",
			"http://localhost:4221/files/test.txt",
			bytes.NewBuffer(content),
		)
		if err != nil {
			t.Fatalf("Failed to create request: %v", err)
		}

		req.Header.Set("Content-Type", "application/octet-stream")
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatalf("Failed to make request: %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusCreated {
			t.Errorf("Expected status 201, got %d", resp.StatusCode)
		}

		// Verify file was created
		content, err = os.ReadFile(filepath.Join(tempDir, "test.txt"))
		if err != nil {
			t.Fatalf("Failed to read created file: %v", err)
		}
		if string(content) != "Hello, World!" {
			t.Errorf("Expected file content %q, got %q", "Hello, World!", string(content))
		}
	})

	// Test file retrieval (GET)
	t.Run("GET existing file", func(t *testing.T) {
		// Create a test file
		testContent := []byte("Test content")
		testFile := filepath.Join(tempDir, "get-test.txt")
		if err := os.WriteFile(testFile, testContent, 0644); err != nil {
			t.Fatalf("Failed to create test file: %v", err)
		}

		resp, err := http.Get("http://localhost:4221/files/get-test.txt")
		if err != nil {
			t.Fatalf("Failed to make request: %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			t.Errorf("Expected status 200, got %d", resp.StatusCode)
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Fatalf("Failed to read response body: %v", err)
		}

		if string(body) != string(testContent) {
			t.Errorf("Expected body %q, got %q", string(testContent), string(body))
		}

		if resp.Header.Get("Content-Type") != "application/octet-stream" {
			t.Errorf("Expected Content-Type application/octet-stream, got %s", resp.Header.Get("Content-Type"))
		}
	})

	// Test non-existent file
	t.Run("GET non-existent file", func(t *testing.T) {
		resp, err := http.Get("http://localhost:4221/files/non-existent.txt")
		if err != nil {
			t.Fatalf("Failed to make request: %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusNotFound {
			t.Errorf("Expected status 404, got %d", resp.StatusCode)
		}
	})
} 