package installerx

import "dagger.io/dagger"

// Installer defines the interface for all installers. It provides methods to get installation commands,
// perform the installation on a given container, and retrieve the latest version of the software.
type Installer interface {
	// GetInstallCommands returns a slice of slices of strings, where each inner slice represents
	// a set of commands to be executed for installation. Each command set is executed in sequence.
	GetInstallCommands() [][]string

	// Install takes a pointer to a dagger.Container and performs the installation process on it.
	// It returns the modified container after the installation is complete.
	Install(container *dagger.Container) *dagger.Container

	// GetLatestVersion fetches the latest version of the software to be installed.
	// It returns the version as a string and an error if there is any issue in fetching the version.
	GetLatestVersion() (string, error)
}
