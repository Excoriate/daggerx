// Package workdir provides utilities for managing and validating working directories
// within the context of the Excoriate DaggerX project. This package includes functions
// to get default working directory paths, set working directories with default fallbacks,
// and validate the correctness of provided working directory paths.
//
// The primary functions in this package are:
// - GetDefault: Returns the default mount prefix from fixtures.
// - SetOrDefault: Returns the provided workdir if it is not empty; otherwise, it returns the default workdir.
// - IsValid: Checks if the provided workdir is valid.
//
// This package is designed to ensure that working directories are correctly set and validated
// according to the requirements of the Excoriate DaggerX project, providing a consistent and
// reliable way to manage working directories.
//
// Example usage:
//
//	import (
//	    "fmt"
//	    "github.com/Excoriate/daggerx/pkg/workdir"
//	)
//
//	func main() {
//	    defaultWorkdir := workdir.GetDefault()
//	    fmt.Println("Default Workdir:", defaultWorkdir)
//
//	    customWorkdir := "/custom/path"
//	    validatedWorkdir := workdir.SetOrDefault(customWorkdir)
//	    fmt.Println("Validated Workdir:", validatedWorkdir)
//
//	    err := workdir.IsValid(validatedWorkdir)
//	    if err != nil {
//	        fmt.Println("Invalid Workdir:", err)
//	    } else {
//	        fmt.Println("Workdir is valid")
//	    }
//	}
package workdir

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/Excoriate/daggerx/pkg/fixtures"
)

// GetDefault returns the default mount prefix from fixtures.
func GetDefault() string {
	return fixtures.MntPrefix
}

// SetOrDefault returns the provided workdir if it is not empty;
// otherwise, it returns the default workdir.
//
// Parameters:
//   - workdir: The workdir path to validate.
//
// Returns:
//   - The provided workdir if it is not empty and starts with '/'.
//   - The default workdir if the provided workdir is empty.
func SetOrDefault(workdir string) string {
	if workdir == "" {
		return GetDefault()
	}

	// Ensure the workdir starts with a '/'
	if !strings.HasPrefix(workdir, "/") {
		workdir = "/" + workdir
	}

	return workdir
}

// IsValid checks if the provided workdir is valid.
//
// A valid workdir must:
// - Be an absolute path.
// - Start with the mount prefix from fixtures.
//
// Parameters:
//   - workdir: The workdir path to validate.
//
// Returns:
//   - An error if the workdir is invalid; otherwise, nil.
func IsValid(workdir string) error {
	if !filepath.IsAbs(workdir) {
		return fmt.Errorf("workdir must be an absolute path: %s", workdir)
	}

	if !strings.HasPrefix(workdir, fixtures.MntPrefix) {
		return fmt.Errorf("workdir must start with %s: %s", fixtures.MntPrefix, workdir)
	}

	return nil
}
