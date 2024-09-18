package installerx

import (
	"fmt"
	"strings"

	"dagger.io/dagger"
)

const openTofuReleaseURL = "https://github.com/opentofu/opentofu/releases/download"

// OpenTofuInstaller handles the installation of OpenTofu.
type OpenTofuInstaller struct {
	version string
}

// NewOpenTofuInstaller creates a new OpenTofuInstaller instance.
func NewOpenTofuInstaller(version string) *OpenTofuInstaller {
	return &OpenTofuInstaller{
		version: strings.TrimPrefix(version, "v"),
	}
}

// GetInstallCommands returns the commands to install OpenTofu.
func (oti *OpenTofuInstaller) GetInstallCommands() [][]string {
	url := fmt.Sprintf("%s/v%s/tofu_%s_linux_amd64.zip", openTofuReleaseURL, oti.version, oti.version)
	return [][]string{
		{"mkdir", "-p", "/usr/local/bin"},
		{"curl", "-L", "-o", "/tmp/tofu.zip", url},
		{"unzip", "-d", "/usr/local/bin", "/tmp/tofu.zip"},
		{"chmod", "+x", "/usr/local/bin/tofu"},
		{"rm", "/tmp/tofu.zip"},
		{"tofu", "--version"},
	}
}

// Install performs the OpenTofu installation on a Dagger container.
func (oti *OpenTofuInstaller) Install(container *dagger.Container) *dagger.Container {
	commands := oti.GetInstallCommands()
	for _, cmd := range commands {
		container = container.WithExec(cmd)
	}
	return container
}

// GetLatestVersion fetches the latest version for OpenTofu.
func (oti *OpenTofuInstaller) GetLatestVersion() (string, error) {
	// TODO: Implement logic to fetch the latest version from OpenTofu's releases page or API
	return "1.5.0", nil
}
