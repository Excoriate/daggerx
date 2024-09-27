package cmdx

import (
	"strings"
	"testing"
	"time"
)

func TestBuildCurlCommand(t *testing.T) {
	tests := []struct {
		name            string
		baseURL         string
		headers         map[string]string
		timeout         time.Duration
		authType        string
		authCredentials string
		want            []string
	}{
		{
			name:    "Basic command",
			baseURL: "https://api.example.com",
			timeout: 30 * time.Second,
			want:    []string{"curl -m 30 'https://api.example.com'"},
		},
		{
			name:    "With headers",
			baseURL: "https://api.example.com",
			headers: map[string]string{
				"Content-Type": "application/json",
				"Accept":       "application/json",
			},
			timeout: 60 * time.Second,
			want: []string{
				"curl -m 60",
				"-H 'Content-Type: application/json'",
				"-H 'Accept: application/json'",
				"'https://api.example.com'",
			},
		},
		{
			name:            "With basic auth",
			baseURL:         "https://api.example.com",
			timeout:         45 * time.Second,
			authType:        "basic",
			authCredentials: "user:pass",
			want:            []string{"curl -m 45 -u 'user:pass' 'https://api.example.com'"},
		},
		{
			name:            "With bearer auth",
			baseURL:         "https://api.example.com",
			timeout:         90 * time.Second,
			authType:        "bearer",
			authCredentials: "token123",
			want:            []string{"curl -m 90 -H 'Authorization: Bearer token123' 'https://api.example.com'"},
		},
		{
			name:    "With all options",
			baseURL: "https://api.example.com",
			headers: map[string]string{
				"Content-Type": "application/json",
				"User-Agent":   "TestApp/1.0",
			},
			timeout:         120 * time.Second,
			authType:        "bearer",
			authCredentials: "token456",
			want: []string{
				"curl -m 120",
				"-H 'Content-Type: application/json'",
				"-H 'User-Agent: TestApp/1.0'",
				"-H 'Authorization: Bearer token456'",
				"'https://api.example.com'",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := BuildCurlCommand(tt.baseURL, tt.headers, tt.timeout, tt.authType, tt.authCredentials)
			for _, want := range tt.want {
				if !strings.Contains(got, want) {
					t.Errorf("BuildCurlCommand() = %v, want to contain %v", got, want)
				}
			}
		})
	}
}
