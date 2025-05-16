package tests

import (
	"io"
	"net/http"
	"testing"
)

func TestUserAgentEndpoint(t *testing.T) {
	tests := []struct {
		name       string
		userAgent  string
		expectBody string
	}{
		{
			name:       "Custom user agent",
			userAgent:  "MyCustomAgent/1.0",
			expectBody: "MyCustomAgent/1.0",
		},
		{
			name:       "Empty user agent",
			userAgent:  "",
			expectBody: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &http.Client{}
			req, err := http.NewRequest("GET", "http://localhost:4221/user-agent", nil)
			if err != nil {
				t.Fatalf("Failed to create request: %v", err)
			}

			req.Header.Set("User-Agent", tt.userAgent)
			resp, err := client.Do(req)
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

			if string(body) != tt.expectBody {
				t.Errorf("Expected body %q, got %q", tt.expectBody, string(body))
			}

			if resp.Header.Get("Content-Type") != "text/plain" {
				t.Errorf("Expected Content-Type text/plain, got %s", resp.Header.Get("Content-Type"))
			}

			if resp.Header.Get("Content-Length") != string(len(tt.expectBody)) {
				t.Errorf("Expected Content-Length %d, got %s", len(tt.expectBody), resp.Header.Get("Content-Length"))
			}
		})
	}
} 