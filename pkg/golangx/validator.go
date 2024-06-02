package golangx

import (
	"fmt"
	"os"
	"path/filepath"
)

// IsGoModule checks if the given path is a Go module.
//
// Parameters:
//   - path: The path to check.
//
// Returns:
//   - An error if the path is not a Go module.
//
// Example:
//
//	err := IsGoModule("path/to/go/module")
//	if err != nil {
//	    // handle error
//	}
//	// Use err, e.g., fmt.Println(err) // Output: path is not a Go module
func IsGoModule(path string) error {
	if path == "" {
		return fmt.Errorf("path is empty")
	}

	goMod := "go.mod"
	if _, err := os.Stat(filepath.Join(path, goMod)); err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("path is not a Go module")
		}
	}

	return nil
}
