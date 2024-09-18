package containerx

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/Excoriate/daggerx/pkg/fixtures"
)

// SetDefaultImageNameIfEmpty returns a fallback image if the provided image is empty.
// If both the provided image and the fallback image are empty, it returns the default image from fixtures.
func SetDefaultImageNameIfEmpty(image, fallbackImage string) string {
	if image != "" {
		return image
	}
	if fallbackImage != "" {
		return fallbackImage
	}
	return fixtures.Image
}

// SetDefaultImageVersionIfEmpty returns a fallback version if the provided version is empty.
// If both the provided version and the fallback version are empty, it returns "latest".
func SetDefaultImageVersionIfEmpty(version, fallbackVersion string) string {
	if version != "" {
		return version
	}
	if fallbackVersion != "" {
		return fallbackVersion
	}
	return "latest"
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
//
// This function takes a pointer to NewBaseContainerOpts and constructs a Docker image URL
// by combining the image name and version. If the provided image name or version is empty,
// it uses fallback values. If both the provided and fallback values are empty, it uses
// default values.
//
// Parameters:
//   - opts (*NewBaseContainerOpts): A pointer to a NewBaseContainerOpts struct containing
//     the image name, version, fallback image name, and fallback version.
//
// Returns:
//   - (string, error): Returns the constructed image URL as a string and an error if the
//     options pointer is nil.
//
// Example:
//
//	opts := &NewBaseContainerOpts{
//	    Image:         "myimage",
//	    Version:       "1.0",
//	    FallbackImage: "fallbackimage",
//	    FallBackVersion: "latest",
//	}
//	imageURL, err := GetImageURL(opts)
//	if err != nil {
//	    log.Fatalf("Error: %v", err)
//	}
//	fmt.Println(imageURL) // Output: myimage:1.0
func GetImageURL(opts *NewBaseContainerOpts) (string, error) {
	if opts == nil {
		return "", fmt.Errorf("failed to create base container: opts is nil")
	}

	image := SetDefaultImageNameIfEmpty(opts.Image, opts.FallbackImage)
	version := SetDefaultImageVersionIfEmpty(opts.Version, opts.FallBackVersion)

	return fmt.Sprintf("%s:%s", image, version), nil
}

// ValidateImageURL verifies the correctness and validity of a given Docker image URL.
//
// This function performs several validation steps to ensure the image URL adheres to expected formats
// and does not contain any invalid characters or components. The validation includes:
//
//  1. **Empty Check**: Ensures that the `imageURL` is not an empty string.
//  2. **Structure Validation**: Splits the `imageURL` into segments using the '/' delimiter and checks
//     that the number of components does not exceed four, which could indicate an improperly formatted URL.
//  3. **Repository Name Validation**: Iterates through all repository components except the last one to
//     verify that they do not contain the '@' character, which is reserved for specifying digests.
//  4. **Registry Validation**: If a registry is present in the URL, it validates the registry's format
//     to ensure it conforms to expected naming conventions.
//  5. **Namespace and Repository Validation**: Checks that each namespace and repository name within the URL
//     adheres to valid naming rules, preventing the inclusion of invalid characters or formats.
//  6. **Tag and Digest Validation**: Examines the last part of the URL to ensure that any specified
//     tag or digest is correctly formatted. Specifically, if a digest is present, it must start with
//     a supported hashing algorithm prefix like `sha256:` or `sha512:`.
//
// **Parameters:**
// - `imageURL` (string): The Docker image URL to be validated.
//
// **Returns:**
// - `bool`: Returns `true` if the `imageURL` passes all validation checks; otherwise, returns `false`.
// - `error`: Provides an error detailing the reason for validation failure, if any.
func ValidateImageURL(imageURL string) (bool, error) {
	if imageURL == "" {
		return false, fmt.Errorf("image URL cannot be empty")
	}

	parts := strings.Split(imageURL, "/")

	// Check for '@' in repository components except the last part
	for _, part := range parts[:len(parts)-1] {
		if strings.Contains(part, "@") {
			return false, fmt.Errorf("invalid '@' character in repository name: %s", part)
		}
	}

	if len(parts) > 4 {
		return false, fmt.Errorf("too many components in image URL")
	}

	if err := validateRegistryIfPresent(parts); err != nil {
		return false, err
	}

	if err := validateNamespaceAndRepo(parts); err != nil {
		return false, err
	}

	return validateLastPart(parts[len(parts)-1])
}

func validateRegistryIfPresent(parts []string) error {
	if len(parts) >= 2 && strings.Contains(parts[0], ".") {
		if !validateRegistry(parts[0]) {
			return fmt.Errorf("invalid registry: %s", parts[0])
		}
	}
	return nil
}

func validateRegistry(registry string) bool {
	// Disallow consecutive dots, dashes, or underscores
	return regexp.MustCompile(`^[a-zA-Z0-9]+(?:[.-][a-zA-Z0-9]+)*(:\d+)?$`).MatchString(registry)
}

func validateNamespaceAndRepo(parts []string) error {
	startIndex := 0
	if len(parts) >= 2 && strings.Contains(parts[0], ".") {
		startIndex = 1
	}
	for i := startIndex; i < len(parts)-1; i++ {
		if !validateNamespaceOrRepo(parts[i]) {
			return fmt.Errorf("invalid %s: %s", getComponentName(i, len(parts)), parts[i])
		}
	}
	return nil
}

//nolint:cyclop // TODO: Refactor this function to reduce complexity
func validateLastPart(lastPart string) (bool, error) {
	// Check if '@' is present in the repository name
	if strings.Contains(lastPart, "@") {
		// Split on '@' to see if it's a valid digest separator
		parts := strings.SplitN(lastPart, "@", 2)
		if len(parts) != 2 || !strings.HasPrefix(parts[1], "sha256:") && !strings.HasPrefix(parts[1], "sha512:") {
			return false, fmt.Errorf("invalid repository name: %s", lastPart)
		}
	}

	// Split on '@' to separate digest if present
	tagAndDigest := strings.SplitN(lastPart, "@", 2)

	tagPart := tagAndDigest[0]
	digest := ""
	if len(tagAndDigest) == 2 {
		digest = tagAndDigest[1]
	}

	// Split tagPart into repoName and tag if present
	tagParts := strings.SplitN(tagPart, ":", 2)
	repoName := tagParts[0]
	tag := ""
	if len(tagParts) == 2 {
		tag = tagParts[1]
	}

	// Validate repoName
	if !validateNamespaceOrRepo(repoName) {
		return false, fmt.Errorf("invalid repository name: %s", repoName)
	}

	// Check for empty tag when ':' is present
	if len(tagParts) == 2 {
		if tag == "" {
			return false, fmt.Errorf("tag cannot be empty")
		}
		if !validateTag(tag) {
			return false, fmt.Errorf("invalid tag: %s", tag)
		}
	}

	// Validate digest if present
	if digest != "" {
		if !validateDigest(digest) {
			return false, fmt.Errorf("invalid digest: %s", digest)
		}
	}

	return true, nil
}

func validateNamespaceOrRepo(name string) bool {
	// This regex disallows '@' characters
	return regexp.MustCompile(`^[a-zA-Z0-9]+(?:[._-][a-zA-Z0-9]+)*$`).MatchString(name)
}

func validateTag(tag string) bool {
	return regexp.MustCompile(`^[a-zA-Z0-9_.-]+$`).MatchString(tag)
}

func validateDigest(digest string) bool {
	// Accept algorithms like sha256, sha512, etc., and hexadecimal hashes of exact 64 characters
	return regexp.MustCompile(`^[A-Za-z0-9_+.-]+:[a-fA-F0-9]{64}$`).MatchString(digest)
}

func getComponentName(index, totalParts int) string {
	if index == totalParts-2 {
		return "repository"
	}
	return "namespace"
}
