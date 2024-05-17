package containerx

import "dagger.io/dagger"

// WithEnvVariable sets an environment variable in the provided container.
//
// Parameters:
//   - ctr: The base container to which the environment variable will be added.
//   - name: The name of the environment variable (e.g., "HOST").
//   - value: The value of the environment variable (e.g., "localhost").
//   - expand: If true, replaces `${VAR}` or `$VAR` in the value according to the current environment variables defined in the container.
//
// Returns:
//   - A pointer to the Dagger container with the environment variable set.
//
// Example:
//
//	// Create a new Dagger container
//	ctr := dagger.NewContainer()
//
//	// Set an environment variable
//	updatedCtr := WithEnvVariable(ctr, "HOST", "localhost", false)
//	// The container now has the environment variable "HOST" set to "localhost"
//
//	// Set an environment variable with expansion
//	updatedCtr = WithEnvVariable(ctr, "PATH", "/usr/local/bin:$PATH", true)
//	// The container now has the environment variable "PATH" with expansion
func WithEnvVariable(
	ctr *dagger.Container, // The base container to use.
	name string, // The name of the environment variable.
	value string, // The value of the environment variable.
	expand bool, // Whether to expand the value according to current environment variables.
) *dagger.Container {
	return ctr.WithEnvVariable(name, value, dagger.ContainerWithEnvVariableOpts{
		Expand: expand,
	})
}
