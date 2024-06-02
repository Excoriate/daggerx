package golangx

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestIsGoModule(t *testing.T) {
	// Create a temporary directory.
	tempDir, err := ioutil.TempDir("", "gomodule")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	defer func(path string) {
		_ = os.RemoveAll(path)
	}(tempDir) // clean up

	// Create a go.mod file in the temporary directory.
	goModPath := filepath.Join(tempDir, "go.mod")
	if err := ioutil.WriteFile(goModPath, []byte("module temp"), 0644); err != nil {
		t.Fatalf("Failed to write go.mod file: %v", err)
	}

	tests := []struct {
		name    string
		path    string
		wantErr bool
	}{
		{
			name:    "Path is a Go module",
			path:    tempDir,
			wantErr: false,
		},
		{
			name:    "Path is not a Go module",
			path:    "/nonexistent",
			wantErr: true,
		},
		{
			name:    "Path is empty",
			path:    "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := IsGoModule(tt.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("IsGoModule() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
