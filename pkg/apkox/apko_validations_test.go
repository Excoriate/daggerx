package apkox

import (
	"testing"
)

func TestIsKeyringFormatValid(t *testing.T) {
	tests := []struct {
		name         string
		keyrings     []string
		enforceHTTPS bool
		wantErr      bool
		errMsg       string
	}{
		{
			name:         "Valid keyring with path",
			keyrings:     []string{"/etc/apk/keys/foo=https://example.com/key.pub"},
			enforceHTTPS: true,
			wantErr:      false,
		},
		{
			name:         "Valid keyring without path",
			keyrings:     []string{"https://example.com/key.pub"},
			enforceHTTPS: true,
			wantErr:      false,
		},
		{
			name:         "Invalid format - extra '='",
			keyrings:     []string{"/etc/apk/keys/foo=https://example.com/key.pub=extra"},
			enforceHTTPS: true,
			wantErr:      true,
			errMsg:       "invalid keyring format: /etc/apk/keys/foo=https://example.com/key.pub=extra",
		},
		{
			name:         "Invalid path",
			keyrings:     []string{"/wrong/path/foo=https://example.com/key.pub"},
			enforceHTTPS: true,
			wantErr:      true,
			errMsg:       "invalid keyring path: /wrong/path/foo",
		},
		{
			name:         "Invalid URL - HTTP when HTTPS required",
			keyrings:     []string{"/etc/apk/keys/foo=http://example.com/key.pub"},
			enforceHTTPS: true,
			wantErr:      true,
			errMsg:       "invalid keyring URL: http://example.com/key.pub, error: HTTPS is required",
		},
		{
			name:         "Valid URL - HTTP when HTTPS not required",
			keyrings:     []string{"/etc/apk/keys/foo=http://example.com/key.pub"},
			enforceHTTPS: false,
			wantErr:      false,
		},
		{
			name: "Multiple keyrings with one invalid",
			keyrings: []string{
				"/etc/apk/keys/foo=https://example.com/key.pub",
				"/etc/apk/keys/bar=http://example.com/key.pub",
			},
			enforceHTTPS: true,
			wantErr:      true,
			errMsg:       "invalid keyring URL: http://example.com/key.pub, error: HTTPS is required",
		},
		{
			name: "Multiple valid keyrings",
			keyrings: []string{
				"/etc/apk/keys/foo=https://example.com/key.pub",
				"https://example.com/key2.pub",
			},
			enforceHTTPS: true,
			wantErr:      false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := IsKeyringFormatValid(tt.keyrings, tt.enforceHTTPS)
			if (err != nil) != tt.wantErr {
				t.Errorf("IsKeyringFormatValid() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil && tt.errMsg != "" && err.Error() != tt.errMsg {
				t.Errorf("IsKeyringFormatValid() error = %v, wantErrMsg %v", err.Error(), tt.errMsg)
			}
		})
	}
}
