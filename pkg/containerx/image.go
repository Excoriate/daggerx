package containerx

import (
	"fmt"
	"github.com/Excoriate/daggerx/pkg/fixtures"
)

// setDefaultIfEmpty returns a default version if the provided version is empty.
func setDefaultIfEmpty(version string) string {
	if version == "" {
		return fixtures.ImageVersion
	}
	return version
}

type NewBaseContainerOpts struct {
	// Image is the name of the image to use.
	Image string
	// Version is the version of the image to use.
	Version string
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
