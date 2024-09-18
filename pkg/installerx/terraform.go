package installerx

import (
	"fmt"
)

const terraformReleaseURL = "https://releases.hashicorp.com/terraform"

type TerraformInstaller struct {
	*BaseInstaller
}

func NewTerraformInstaller(version string) *TerraformInstaller {
	return &TerraformInstaller{
		BaseInstaller: NewBaseInstaller(version, terraformReleaseURL, "terraform", "zip"),
	}
}

func (ti *TerraformInstaller) GetInstallCommands() [][]string {
	url := fmt.Sprintf("%s/%s/terraform_%s_linux_amd64.zip", ti.releaseURL, ti.version, ti.version)
	return ti.BaseInstaller.GetInstallCommands(url)
}

func (ti *TerraformInstaller) GetLatestVersion() (string, error) {
	// TODO: Implement logic to fetch the latest version from Terraform's releases page or API
	return "1.9.4", nil
}
