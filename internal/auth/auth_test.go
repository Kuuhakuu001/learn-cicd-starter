package auth_test

import (
	"net/http"
	"testing"

	"github.com/bootdotdev/learn-cicd-starter/internal/auth"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name        string
		headers     http.Header
		want        string
		expectError string
	}{
		{
			name: "valid API key",
			headers: http.Header{
				"Authorization": []string{"ApiKey test-api-key"},
			},
			want:        "test-api-key",
			expectError: "",
		},
		{
			name:        "missing authorization header",
			headers:     http.Header{},
			want:        "",
			expectError: "no authorization header included",
		},
		{
			name: "malformed header - wrong prefix",
			headers: http.Header{
				"Authorization": []string{"Bearer test-api-key"},
			},
			want:        "",
			expectError: "malformed authorization header",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := auth.GetAPIKey(tt.headers)

			if tt.expectError != "" {
				if err == nil {
					t.Errorf("GetAPIKey() expected error %v, got nil", tt.expectError)
					return
				}
				if err.Error() != tt.expectError {
					t.Errorf("GetAPIKey() expected error %v, got %v", tt.expectError, err)
					return
				}
			} else if err != nil {
				t.Errorf("GetAPIKey() unexpected error: %v", err)
				return
			}

			if got != tt.want {
				t.Errorf("GetAPIKey() = %v, want %v", got, tt.want)
			}
		})
	}
}
