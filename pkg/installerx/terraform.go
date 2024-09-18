package installerx

import (
	"fmt"

	"dagger.io/dagger"
)

const terraformReleaseURL = "https://releases.hashicorp.com/terraform"

// TerraformInstaller handles the installation of Terraform.
type TerraformInstaller struct {
	version string
}

// NewTerraformInstaller creates a new TerraformInstaller instance.
func NewTerraformInstaller(version string) *TerraformInstaller {
	return &TerraformInstaller{
		version: strings.TrimPrefix(version, "v"),
	}
}

// GetInstallCommands returns the commands to install Terraform.
func (ti *TerraformInstaller) GetInstallCommands() [][]string {
	url := fmt.Sprintf("%s/%s/terraform_%s_linux_amd64.zip", terraformReleaseURL, ti.version, ti.version)
	return [][]string{
		{"mkdir", "-p", "/usr/local/bin"},
		{"curl", "-L", "-o", "/tmp/terraform.zip", url},
		{"unzip", "-d", "/usr/local/bin", "/tmp/terraform.zip"},
		{"chmod", "+x", "/usr/local/bin/terraform"},
		{"rm", "/tmp/terraform.zip"},
		{"terraform", "--version"},
	}
}

// Install performs the Terraform installation on a Dagger container.
func (ti *TerraformInstaller) Install(container *dagger.Container) *dagger.Container {
	commands := ti.GetInstallCommands()
	for _, cmd := range commands {
		container = container.WithExec(cmd)
	}
	return container
}

// GetLatestVersion fetches the latest version for Terraform.
func (ti *TerraformInstaller) GetLatestVersion() (string, error) {
	// TODO: Implement logic to fetch the latest version from Terraform's releases page or API
	return "1.9.4", nil
}
