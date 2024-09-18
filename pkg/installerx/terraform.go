package installerx

import (
	"fmt"
)

// terraformReleaseURL is the base URL for Terraform releases
const terraformReleaseURL = "https://releases.hashicorp.com/terraform"

// TerraformInstaller represents an installer for Terraform
type TerraformInstaller struct {
	*BaseInstaller
}

// NewTerraformInstaller creates a new TerraformInstaller instance
//
// Parameters:
//   - version: The version of Terraform to install
//
// Returns:
//   - *TerraformInstaller: A pointer to the newly created TerraformInstaller
func NewTerraformInstaller(version string) *TerraformInstaller {
	return &TerraformInstaller{
		BaseInstaller: NewBaseInstaller(version, terraformReleaseURL, "terraform", "zip"),
	}
}

// GetInstallCommands returns the commands needed to install Terraform
//
// This method generates the download URL for the specified Terraform version
// and returns the installation commands using the BaseInstaller.
//
// Returns:
//   - [][]string: A slice of string slices representing the installation commands
func (ti *TerraformInstaller) GetInstallCommands() [][]string {
	url := fmt.Sprintf("%s/%s/terraform_%s_linux_amd64.zip", ti.releaseURL, ti.version, ti.version)
	return ti.BaseInstaller.GetInstallCommands(url)
}

// GetLatestVersion retrieves the latest version of Terraform
//
// This method is currently a placeholder and returns a hardcoded version.
// TODO: Implement logic to fetch the latest version from Terraform's releases page or API
//
// Returns:
//   - string: The latest version of Terraform (currently hardcoded)
//   - error: An error if the retrieval fails (currently always nil)
func (ti *TerraformInstaller) GetLatestVersion() (string, error) {
	// TODO: Implement logic to fetch the latest version from Terraform's releases page or API
	return "1.9.4", nil
}
