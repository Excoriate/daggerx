package execmd

import (
	"dagger.io/dagger"
	"github.com/Excoriate/daggerx/pkg/types"
)

// AddToContainer adds commands to a container.
// This function is useful when you want to inject commands into a container
// before running them.
func AddToContainer(ctr *dagger.Container, execCMDs []types.ContainerCommand) *dagger.Container {
	for _, cmd := range execCMDs {
		if cmd.EnableFocus {
			ctr = ctr.WithFocus()
		}

		ctr = ctr.WithExec(cmd.CMD, cmd.ContainerCMDOptions)
	}
	return ctr
}
