package installerx

import "dagger.io/dagger"

// Installer defines the interface for all installers
type Installer interface {
	GetInstallCommands() [][]string
	Install(container *dagger.Container) *dagger.Container
	GetLatestVersion() (string, error)
}
