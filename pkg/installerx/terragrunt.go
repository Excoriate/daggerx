package installerx

import (
	"fmt"
	"strings"

	"dagger.io/dagger"
)

const terragruntReleaseURL = "https://github.com/gruntwork-io/terragrunt/releases/download"

// TerragruntInstaller handles the installation of Terragrunt.
type TerragruntInstaller struct {
	BaseInstaller
}

// NewTerragruntInstaller creates a new TerragruntInstaller instance.
func NewTerragruntInstaller(version string) *TerragruntInstaller {
	return &TerragruntInstaller{
		BaseInstaller: NewBaseInstaller(version, terragruntReleaseURL, "terragrunt", ""),
	}
}

// GetInstallCommands returns the commands to install Terragrunt.
func (tgi *TerragruntInstaller) GetInstallCommands() [][]string {
	url := fmt.Sprintf("%s/v%s/terragrunt_linux_amd64", tgi.releaseURL, tgi.version)
	return tgi.BaseInstaller.GetInstallCommands(url)
}

// Install performs the Terragrunt installation on a Dagger container.
func (tgi *TerragruntInstaller) Install(container *dagger.Container) *dagger.Container {
	return tgi.BaseInstaller.Install(container, tgi.GetInstallCommands())
}

// GetLatestVersion fetches the latest version for Terragrunt.
func (tgi *TerragruntInstaller) GetLatestVersion() (string, error) {
	// TODO: Implement logic to fetch the latest version from Terragrunt's releases page or API
	return "0.67.4", nil
}
