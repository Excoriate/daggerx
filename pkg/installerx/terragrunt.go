// Package installerx provides functionality for installing various tools, including Terragrunt.
package installerx

import (
	"fmt"
)

// terragruntReleaseURL is the base URL for downloading Terragrunt releases.
const terragruntReleaseURL = "https://github.com/gruntwork-io/terragrunt/releases/download"

// TerragruntInstaller represents an installer for Terragrunt.
// It embeds BaseInstaller to inherit common installation functionality.
type TerragruntInstaller struct {
	*BaseInstaller
}

// NewTerragruntInstaller creates and returns a new TerragruntInstaller.
// It initializes the embedded BaseInstaller with Terragrunt-specific parameters.
//
// Parameters:
//   - version: The version of Terragrunt to install.
//
// Returns:
//   - *TerragruntInstaller: A pointer to the newly created TerragruntInstaller.
func NewTerragruntInstaller(version string) *TerragruntInstaller {
	return &TerragruntInstaller{
		BaseInstaller: NewBaseInstaller(version, terragruntReleaseURL, "terragrunt", "", "/app/bin"),
	}
}

// GetInstallCommands returns the commands needed to install Terragrunt.
// It generates the download URL for the specified version and delegates to BaseInstaller.
//
// Returns:
//   - [][]string: A slice of command arguments, where each inner slice represents a single command.
func (tgi *TerragruntInstaller) GetInstallCommands() [][]string {
	url := fmt.Sprintf("%s/v%s/terragrunt_linux_amd64", tgi.releaseURL, tgi.version)
	return tgi.BaseInstaller.GetInstallCommands(url)
}

// GetLatestVersion retrieves the latest version of Terragrunt available.
//
// Returns:
//   - string: The latest version of Terragrunt.
//   - error: An error if the version retrieval fails.
//
// Note: This is currently a placeholder implementation.
func (tgi *TerragruntInstaller) GetLatestVersion() (string, error) {
	// TODO: Implement logic to fetch the latest version from Terragrunt's releases page or API
	return "0.67.4", nil
}
