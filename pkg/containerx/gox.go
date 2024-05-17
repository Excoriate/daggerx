package containerx

import (
	"dagger.io/dagger"
	"github.com/containerd/containerd/platforms"
)

// WithGoPlatform sets the GOOS, GOARCH, and GOARM environment variables based on the target platform.
//
// Parameters:
//   - platform: The target platform in "[os]/[platform]/[version]" format (e.g., "darwin/arm64/v7", "windows/amd64", "linux/arm64").
//   - ctr: The base container to which the environment variables will be added.
//
// Returns:
//   - A pointer to the Dagger container with the platform environment variables set.
//
// Example:
//
//	// Create a new Dagger container
//	ctr := dagger.NewContainer()
//
//	// Set the platform environment variables
//	updatedCtr := WithPlatform("linux/amd64", ctr)
//	// The container now has the environment variables "GOOS", "GOARCH", and "GOARM" set accordingly
func WithGoPlatform(
	ctr *dagger.Container, // The base container to use.
	platform dagger.Platform, // The target platform in "[os]/[platform]/[version]" format.
) interface{} {
	if platform == "" {
		return ctr
	}

	p := platforms.MustParse(string(platform))

	ctr = ctr.
		WithEnvVariable("GOOS", p.OS).
		WithEnvVariable("GOARCH", p.Architecture)

	if p.Variant != "" {
		ctr = ctr.WithEnvVariable("GOARM", p.Variant)
	}

	return ctr
}

// WithGoCgoEnabled enables CGO in the provided container by setting the "CGO_ENABLED" environment variable to "1".
//
// Parameters:
//   - ctr: A pointer to the Dagger container in which CGO will be enabled.
//
// Returns:
//   - A pointer to the Dagger container with CGO enabled.
//
// Example:
//
//	// Create a new Dagger container
//	ctr := dagger.NewContainer()
//
//	// Enable CGO in the container
//	updatedCtr := WithGoCgoEnabled(ctr)
//	// The container now has the environment variable "CGO_ENABLED" set to "1"
func WithGoCgoEnabled(ctr *dagger.Container) *dagger.Container {
	return ctr.WithEnvVariable("CGO_ENABLED", "1")
}

// WithGoCgoDisabled disables CGO in the provided container by setting the "CGO_ENABLED" environment variable to "0".
//
// Parameters:
//   - ctr: A pointer to the Dagger container in which CGO will be disabled.
//
// Returns:
//   - A pointer to the Dagger container with CGO disabled.
//
// Example:
//
//	// Create a new Dagger container
//	ctr := dagger.NewContainer()
//
//	// Disable CGO in the container
//	updatedCtr := WithGoCgoDisabled(ctr)
//	// The container now has the environment variable "CGO_ENABLED" set to "0"
func WithGoCgoDisabled(ctr *dagger.Container) *dagger.Container {
	return ctr.WithEnvVariable("CGO_ENABLED", "0")
}
