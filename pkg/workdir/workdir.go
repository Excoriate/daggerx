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
