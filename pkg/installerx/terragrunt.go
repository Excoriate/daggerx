package installerx

import (
	"fmt"
)

const terragruntReleaseURL = "https://github.com/gruntwork-io/terragrunt/releases/download"

type TerragruntInstaller struct {
	*BaseInstaller
}

func NewTerragruntInstaller(version string) *TerragruntInstaller {
	return &TerragruntInstaller{
		BaseInstaller: NewBaseInstaller(version, terragruntReleaseURL, "terragrunt", ""),
	}
}

func (tgi *TerragruntInstaller) GetInstallCommands() [][]string {
	url := fmt.Sprintf("%s/v%s/terragrunt_linux_amd64", tgi.releaseURL, tgi.version)
	return tgi.BaseInstaller.GetInstallCommands(url)
}

func (tgi *TerragruntInstaller) GetLatestVersion() (string, error) {
	// TODO: Implement logic to fetch the latest version from Terragrunt's releases page or API
	return "0.67.4", nil
}
