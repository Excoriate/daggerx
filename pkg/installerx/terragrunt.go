package installerx

import (
	"fmt"
	"strings"

	"dagger.io/dagger"
)

const terragruntReleaseURL = "https://github.com/gruntwork-io/terragrunt/releases/download"

// TerragruntInstaller handles the installation of Terragrunt.
type TerragruntInstaller struct {
	version string
}

// NewTerragruntInstaller creates a new TerragruntInstaller instance.
func NewTerragruntInstaller(version string) *TerragruntInstaller {
	return &TerragruntInstaller{
		version: strings.TrimPrefix(version, "v"),
	}
}

// GetInstallCommands returns the commands to install Terragrunt.
func (tgi *TerragruntInstaller) GetInstallCommands() [][]string {
	url := fmt.Sprintf("%s/v%s/terragrunt_linux_amd64", terragruntReleaseURL, tgi.version)
	return [][]string{
		{"mkdir", "-p", "/usr/local/bin"},
		{"curl", "-L", "-o", "/usr/local/bin/terragrunt", url},
		{"chmod", "+x", "/usr/local/bin/terragrunt"},
		{"terragrunt", "--version"},
	}
}

// Install performs the Terragrunt installation on a Dagger container.
func (tgi *TerragruntInstaller) Install(container *dagger.Container) *dagger.Container {
	commands := tgi.GetInstallCommands()
	for _, cmd := range commands {
		container = container.WithExec(cmd)
	}
	return container
}

// GetLatestVersion fetches the latest version for Terragrunt.
func (tgi *TerragruntInstaller) GetLatestVersion() (string, error) {
	// TODO: Implement logic to fetch the latest version from Terragrunt's releases page or API
	return "0.67.4", nil
}
