// Package golangx provides utility functions for working with Go modules and
// validating Go project structures. This package includes functions to check
// if a given path is a Go module, ensuring that the necessary Go module files
// are present in the specified directory.
//
// The primary function in this package is IsGoModule, which checks if a given
// path contains a Go module by looking for the presence of a go.mod file.
//
// Example usage:
//
//	package main
//
//	import (
//	    "fmt"
//	    "path/to/golangx"
//	)
//
//	func main() {
//	    err := golangx.IsGoModule("path/to/go/module")
//	    if err != nil {
//	        fmt.Println(err)
//	    } else {
//	        fmt.Println("This is a valid Go module.")
//	    }
//	}
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
