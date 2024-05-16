package types

import "dagger.io/dagger"

// DaggerCMD is a type for an expected Dagger command.
// The expected command format in Dagger is a slice of strings.
type DaggerCMD []string

// ContainerCommand is a type for a container command.
// It contains a DaggerCMD and a boolean to enable focus.
// If the focus is enabled, the command will add the WithFocus() method to the container's command.
type ContainerCommand struct {
	// CMD is the command to run in the container.
	CMD DaggerCMD
	// EnableFocus is a boolean to enable focus.
	EnableFocus bool
	// ContainerCMDOptions is a wrapper on top of the Dagger ContainerWithExecOpts
	ContainerCMDOptions dagger.ContainerWithExecOpts
}
