package apis

import (
	"dagger.io/dagger"
	"fmt"
	"github.com/Excoriate/daggerx/pkg/fixtures"
	"path/filepath"
)

// WithSource sets the source directory and mounts it to the container.
// If the source directory is not provided (nil), it returns an error.
// The workdir is set to the provided workdir path or defaults to a pre-defined mount prefix.
//
// Parameters:
//   - ctr: A pointer to the Dagger container to which the source directory will be mounted.
//   - src: A pointer to the Dagger directory representing the source directory to be mounted.
//   - workdir: An optional string representing the working directory inside the container.
//
// Returns:
//   - A pointer to the Dagger container with the source directory mounted and workdir set.
//   - An error if the source directory is not provided.
//
// Example:
//
//	// Create a new Dagger container
//	ctr := dagger.NewContainer()
//
//	// Create a new Dagger directory representing the source directory
//	src := dagger.NewDirectory()
//
//	// Set the source directory and workdir
//	updatedCtr, err := WithSource(ctr, src, "app")
//	if err != nil {
//	    // handle error
//	}
//	// The container now has the source directory mounted at the mount prefix
//	// and the workdir set to the combined mount prefix and "app".
func WithSource(
	ctr *dagger.Container,
	src *dagger.Directory,
	workdir string,
) (*dagger.Container, error) {
	if src == nil {
		return nil, fmt.Errorf("failed to set source directory: src is nil")
	}

	var workDirPath string
	if workdir == "" {
		workDirPath = fixtures.MntPrefix
	} else {
		workDirPath = filepath.Join(fixtures.MntPrefix, workdir)
	}

	ctr = ctr.
		WithMountedDirectory(fixtures.MntPrefix, src).
		WithWorkdir(workDirPath)

	return ctr, nil
}
