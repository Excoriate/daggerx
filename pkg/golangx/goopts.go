package golangx

import (
	"github.com/Excoriate/daggerx/pkg/types"
	"github.com/containerd/containerd/platforms"
)

// WithGoCgoEnabled returns a DaggerEnvVars struct that sets the CGO_ENABLED
// environment variable to "1", enabling CGO for Go builds.
//
// Returns:
//   - A DaggerEnvVars struct with CGO_ENABLED set to "1".
//
// Example:
//
//	envVar := WithGoCgoEnabled()
//	fmt.Println(envVar) // Output: {Name: "CGO_ENABLED", Value: "1", Expand: false}
func WithGoCgoEnabled() types.DaggerEnvVars {
	return types.DaggerEnvVars{
		Name:   "CGO_ENABLED",
		Value:  "1",
		Expand: false,
	}
}

// WithGoCgoDisabled returns a DaggerEnvVars struct that sets the CGO_ENABLED
// environment variable to "0", disabling CGO for Go builds.
//
// Returns:
//   - A DaggerEnvVars struct with CGO_ENABLED set to "0".
//
// Example:
//
//	envVar := WithGoCgoDisabled()
//	fmt.Println(envVar) // Output: {Name: "CGO_ENABLED", Value: "0", Expand: false}
func WithGoCgoDisabled() types.DaggerEnvVars {
	return types.DaggerEnvVars{
		Name:   "CGO_ENABLED",
		Value:  "0",
		Expand: false,
	}
}

// WithGoPlatform returns a slice of DaggerEnvVars structs that set the GOOS, GOARCH,
// and optionally GOARM environment variables based on the provided platform string.
//
// Parameters:
//   - platform: A string representing the target platform in the format "os/arch[/variant]".
//
// Returns:
//   - A slice of DaggerEnvVars structs with GOOS, GOARCH, and optionally GOARM set.
//
// Example:
//
//	envVars := WithGoPlatform("linux/amd64")
//	for _, envVar := range envVars {
//	    fmt.Println(envVar)
//	}
//	// Output:
//	// {Name: "GOOS", Value: "linux", Expand: false}
//	// {Name: "GOARCH", Value: "amd64", Expand: false}
//
//	envVars = WithGoPlatform("linux/arm/v7")
//	for _, envVar := range envVars {
//	    fmt.Println(envVar)
//	}
//	// Output:
//	// {Name: "GOOS", Value: "linux", Expand: false}
//	// {Name: "GOARCH", Value: "arm", Expand: false}
//	// {Name: "GOARM", Value: "v7", Expand: false}
func WithGoPlatform(platform string) []types.DaggerEnvVars {
	p := platforms.MustParse(platform)

	envVars := []types.DaggerEnvVars{
		{
			Name:   "GOOS",
			Value:  p.OS,
			Expand: false,
		},
		{
			Name:   "GOARCH",
			Value:  p.Architecture,
			Expand: false,
		},
	}

	if p.Variant != "" {
		envVars = append(envVars, types.DaggerEnvVars{
			Name:   "GOARM",
			Value:  p.Variant,
			Expand: false,
		})
	}

	return envVars
}
