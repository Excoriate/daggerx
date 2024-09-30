package builderx

import (
	"testing"
)

func TestIsKeyringFormatValid(t *testing.T) {
	tests := []struct {
		name     string
		keyrings []string
		wantErr  bool
		errMsg   string
	}{
		{
			name:     "Valid keyring",
			keyrings: []string{"/etc/apk/keys/foo=https://example.com/key.pub"},
			wantErr:  false,
		},
		{
			name:     "Invalid format - missing '='",
			keyrings: []string{"/etc/apk/keys/foo"},
			wantErr:  true,
			errMsg:   "invalid keyring format: /etc/apk/keys/foo",
		},
		{
			name:     "Invalid format - extra '='",
			keyrings: []string{"/etc/apk/keys/foo=https://example.com/key.pub=extra"},
			wantErr:  true,
			errMsg:   "invalid keyring format: /etc/apk/keys/foo=https://example.com/key.pub=extra",
		},
		{
			name:     "Invalid path",
			keyrings: []string{"/wrong/path/foo=https://example.com/key.pub"},
			wantErr:  true,
			errMsg:   "invalid keyring path: /wrong/path/foo",
		},
		{
			name:     "Invalid URL",
			keyrings: []string{"/etc/apk/keys/foo=http://example.com/key.pub"},
			wantErr:  true,
			errMsg:   "invalid keyring URL: http://example.com/key.pub",
		},
		{
			name: "Multiple keyrings with one invalid",
			keyrings: []string{
				"/etc/apk/keys/foo=https://example.com/key.pub",
				"/etc/apk/keys/bar=http://example.com/key.pub",
			},
			wantErr: true,
			errMsg:  "invalid keyring URL: http://example.com/key.pub",
		},
		{
			name: "Multiple valid keyrings",
			keyrings: []string{
				"/etc/apk/keys/foo=https://example.com/key.pub",
				"/etc/apk/keys/bar=https://example.com/key2.pub",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := IsKeyringFormatValid(tt.keyrings)
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
