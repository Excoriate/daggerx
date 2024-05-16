package envvars

import (
	"dagger.io/dagger"
	"github.com/Excoriate/daggerx/pkg/types"
)

// AddToContainers adds environment variables to a container.
func AddToContainers(ctr *dagger.Container, envVars []types.DaggerEnvVars) *dagger.Container {
	for _, envVar := range envVars {
		ctr = ctr.WithEnvVariable(envVar.Name, envVar.Value, dagger.ContainerWithEnvVariableOpts{
			Expand: envVar.Expand,
		})
	}
	return ctr
}
