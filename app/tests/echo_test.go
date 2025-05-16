package tests

import (
	"fmt"
	"net/http"
	"testing"
	"io"
)

func TestEchoEndpoint(t *testing.T) {
	tests := []struct {
		name           string
		path           string
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "Simple echo",
			path:           "/echo/hello",
			expectedStatus: http.StatusOK,
			expectedBody:   "hello",
		},
		{
			name:           "Echo with spaces",
			path:           "/echo/hello world",
			expectedStatus: http.StatusOK,
			expectedBody:   "hello world",
		},
		{
			name:           "Empty echo",
			path:           "/echo/",
			expectedStatus: http.StatusNotFound,
			expectedBody:   "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := fmt.Sprintf("http://localhost:4221%s", tt.path)
			resp, err := http.Get(url)
			if err != nil {
				t.Fatalf("Failed to make request: %v", err)
			}
			defer resp.Body.Close()

			if resp.StatusCode != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, resp.StatusCode)
			}

			if tt.expectedStatus == http.StatusOK {
				body, err := io.ReadAll(resp.Body)
				if err != nil {
					t.Fatalf("Failed to read response body: %v", err)
				}

				if string(body) != tt.expectedBody {
					t.Errorf("Expected body %q, got %q", tt.expectedBody, string(body))
				}

				if resp.Header.Get("Content-Type") != "text/plain" {
					t.Errorf("Expected Content-Type text/plain, got %s", resp.Header.Get("Content-Type"))
				}

				if resp.Header.Get("Content-Length") != fmt.Sprintf("%d", len(tt.expectedBody)) {
					t.Errorf("Expected Content-Length %d, got %s", len(tt.expectedBody), resp.Header.Get("Content-Length"))
				}
			}
		})
	}
} 