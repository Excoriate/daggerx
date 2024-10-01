package apkox

import (
	"testing"
)

func TestKeyringSkeleton(t *testing.T) {
	tests := []struct {
		name string
		ks   KeyringSkeleton
		want KeyringSkeleton
	}{
		{
			name: "KeyringSkeleton with path and URL",
			ks:   KeyringSkeleton{Path: "/etc/apk/keys/foo", URL: "https://example.com/key.pub"},
			want: KeyringSkeleton{Path: "/etc/apk/keys/foo", URL: "https://example.com/key.pub"},
		},
		{
			name: "KeyringSkeleton with only URL",
			ks:   KeyringSkeleton{URL: "https://example.com/key.pub"},
			want: KeyringSkeleton{URL: "https://example.com/key.pub"},
		},
		{
			name: "Empty KeyringSkeleton",
			ks:   KeyringSkeleton{},
			want: KeyringSkeleton{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.ks.Path != tt.want.Path || tt.ks.URL != tt.want.URL {
				t.Errorf("KeyringSkeleton = %v, want %v", tt.ks, tt.want)
			}
		})
	}
}
