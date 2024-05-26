package containerx

import (
	"fmt"

	"github.com/Excoriate/daggerx/pkg/fixtures"
)

// SetDefaultImageVersionIfEmpty returns a default version if the provided version is empty.
// If both the provided version and the fallback version are empty, it returns the default version "latest".
//
// Parameters:
//   - version: The primary version string.
//   - fallbackVersion: The fallback version string.
//
// Returns:
//   - The primary version if it is not empty.
//   - The fallback version if the primary version is empty and the fallback version is not empty.
//   - "latest" if both the primary and fallback versions are empty.
func SetDefaultImageVersionIfEmpty(version, fallbackVersion string) string {
	if version != "" {
		return version
	}

	if fallbackVersion != "" {
		return fallbackVersion
	}

	return fixtures.ImageVersion
}

// SetDefaultImageNameIfEmpty returns a fallback image if the provided image is empty.
// If both the provided image and the fallback image are empty, it returns the default image "alpine".
//
// Parameters:
//   - image: The primary image string.
//   - fallbackImage: The fallback image string.
//
// Returns:
//   - The primary image if it is not empty.
//   - The fallback image if the primary image is empty and the fallback image is not empty.
//   - "alpine" if both the primary and fallback images are empty.
func SetDefaultImageNameIfEmpty(image, fallbackImage string) string {
	if image != "" {
		return image
	}

	if fallbackImage != "" {
		return fallbackImage
	}

	return fixtures.Image
}

type NewBaseContainerOpts struct {
	// Image is the name of the image to use.
	Image string
	// Version is the version of the image to use.
	Version string
	// FallbackImage is the name of the fallback image to use if the primary image is empty.
	FallbackImage string
	// FallBackVersion is the version of the fallback image to use if the primary image is empty.
	FallBackVersion string
}

// GetImageURL constructs the full image URL from the provided options.
// It returns an error if the options are nil or if both the image and fallback image names are empty.
//
// Parameters:
//   - opts: A pointer to NewBaseContainerOpts containing the image name, version, fallback image, and fallback version.
//
// Returns:
//   - A string containing the full image URL in the format "image:version".
//   - An error if the options are nil or if both the image and fallback image names are empty.
//
// Example:
//
//	opts := &NewBaseContainerOpts{
//	    Image:         "golang",
//	    Version:       "1.16",
//	    FallbackImage: "fallback-golang",
//	    FallBackVersion: "1.15",
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

	if opts.Image == "" && opts.FallbackImage == "" {
		return "", fmt.Errorf("failed to create base container: both image and fallback image are empty")
	}

	opts.Image = SetDefaultImageNameIfEmpty(opts.Image, opts.FallbackImage)
	opts.Version = SetDefaultImageVersionIfEmpty(opts.Version, opts.FallBackVersion)

	return fmt.Sprintf("%s:%s", opts.Image, opts.Version), nil
}
