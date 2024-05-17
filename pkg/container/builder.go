package container

import (
	"dagger.io/dagger"
	"dagger.io/dagger/dag"
	"fmt"
	"github.com/Excoriate/daggerx/pkg/fixtures"
)

// NewBaseContainerOpts contains options for creating a new base container.
type NewBaseContainerOpts struct {
	Image   string // The name of the base image.
	Version string // The version of the base image. If empty, a default version is used.
}

// setDefaultIfEmpty returns a default version if the provided version is empty.
func setDefaultIfEmpty(version string) string {
	if version == "" {
		return fixtures.ImageVersion
	}
	return version
}

// GetImageURL constructs the full image URL from the provided options.
// It returns an error if the options are nil or the image name is empty.
//
// Parameters:
//   - opts: A pointer to NewBaseContainerOpts containing the image name and version.
//
// Returns:
//   - A string containing the full image URL in the format "image:version".
//   - An error if the options are nil or the image name is empty.
//
// Example:
//
//	opts := &NewBaseContainerOpts{
//	    Image:   "golang",
//	    Version: "1.16",
//	}
//	imageURL, err := GetImageURL(opts)
//	if err != nil {
//	    // handle error
//	}
//	fmt.Println(imageURL) // Output: "golang:1.16"
func GetImageURL(opts *NewBaseContainerOpts) (string, error) {
	if opts == nil {
		return "", fmt.Errorf("failed to create base container: opts is nil")
	}

	// fail if image is empty
	if opts.Image == "" {
		return "", fmt.Errorf("failed to create base container: image is empty")
	}

	opts.Version = setDefaultIfEmpty(opts.Version)

	return fmt.Sprintf("%s:%s", opts.Image, opts.Version), nil
}

// NewBase creates a new base container from the specified image and version.
// It returns an error if the options are nil or the image name is empty.
//
// Parameters:
//   - opts: A pointer to NewBaseContainerOpts containing the image name and version.
//
// Returns:
//   - A pointer to the Dagger container created from the specified image and version.
//   - An error if the options are nil or the image name is empty.
//
// Example:
//
//	opts := &NewBaseContainerOpts{
//	    Image:   "golang",
//	    Version: "1.16",
//	}
//	container, err := NewBase(opts)
//	if err != nil {
//	    // handle error
//	}
//	// Use the container, e.g., fmt.Println(container)
func NewBase(opts *NewBaseContainerOpts) (*dagger.Container, error) {
	imageURL, err := GetImageURL(opts)
	if err != nil {
		return nil, err
	}

	return dag.Container().From(imageURL), nil
}
