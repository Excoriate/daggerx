package envvars

import (
	"dagger.io/dagger"
	"github.com/Excoriate/daggerx/pkg/types"
)

// AddToContainers adds a list of environment variables to a Dagger container.
//
// Parameters:
//   - ctr: A pointer to the Dagger container to which the environment variables will be added.
//   - envVars: A slice of DaggerEnvVars, each containing the name, value, and expansion option for an environment variable.
//
// Returns:
//   - A pointer to the Dagger container with the added environment variables.
//
// Example:
//
//	// Create a new Dagger container
//	ctr := dagger.NewContainer()
//
//	// Define environment variables
//	envVars := []types.DaggerEnvVars{
//	    {Name: "VAR1", Value: "value1", Expand: false},
//	    {Name: "VAR2", Value: "value2", Expand: true},
//	}
//
//	// Add environment variables to the container
//	updatedCtr := AddToContainers(ctr, envVars)
//
//	// Now, updatedCtr has the environment variables VAR1 and VAR2
func AddToContainers(ctr *dagger.Container, envVars []types.DaggerEnvVars) *dagger.Container {
	for _, envVar := range envVars {
		ctr = ctr.WithEnvVariable(envVar.Name, envVar.Value, dagger.ContainerWithEnvVariableOpts{
			Expand: envVar.Expand,
		})
	}
	return ctr
}
