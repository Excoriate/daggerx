package apkox

import (
	"reflect"
	"testing"
)

func TestParseKeyring(t *testing.T) {
	tests := []struct {
		name    string
		keyring string
		want    KeyringSkeleton
		wantErr bool
	}{
		{
			name:    "Valid keyring with path",
			keyring: "/etc/apk/keys/foo=https://example.com/key.pub",
			want:    KeyringSkeleton{Path: "/etc/apk/keys/foo", URL: "https://example.com/key.pub"},
			wantErr: false,
		},
		{
			name:    "Valid keyring without path",
			keyring: "https://example.com/key.pub",
			want:    KeyringSkeleton{URL: "https://example.com/key.pub"},
			wantErr: false,
		},
		{
			name:    "Invalid path",
			keyring: "/wrong/path/foo=https://example.com/key.pub",
			want:    KeyringSkeleton{},
			wantErr: true,
		},
		{
			name:    "Invalid format",
			keyring: "invalid=format=keyring",
			want:    KeyringSkeleton{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseKeyring(tt.keyring)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseKeyring() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseKeyring() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidateKeyring(t *testing.T) {
	tests := []struct {
		name         string
		keyring      string
		enforceHTTPS bool
		wantErr      bool
	}{
		{
			name:         "Valid keyring with HTTPS",
			keyring:      "/etc/apk/keys/foo=https://example.com/key.pub",
			enforceHTTPS: true,
			wantErr:      false,
		},
		{
			name:         "Valid keyring with HTTP when not enforcing HTTPS",
			keyring:      "/etc/apk/keys/foo=http://example.com/key.pub",
			enforceHTTPS: false,
			wantErr:      false,
		},
		{
			name:         "Invalid keyring with HTTP when enforcing HTTPS",
			keyring:      "/etc/apk/keys/foo=http://example.com/key.pub",
			enforceHTTPS: true,
			wantErr:      true,
		},
		{
			name:         "Invalid path",
			keyring:      "/wrong/path/foo=https://example.com/key.pub",
			enforceHTTPS: true,
			wantErr:      true,
		},
		{
			name:         "Invalid URL",
			keyring:      "/etc/apk/keys/foo=not-a-valid-url",
			enforceHTTPS: true,
			wantErr:      true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateKeyring(tt.keyring, tt.enforceHTTPS)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateKeyring() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
