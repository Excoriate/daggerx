package installerx

import (
	"fmt"
)

const openTofuReleaseURL = "https://github.com/opentofu/opentofu/releases/download"

type OpenTofuInstaller struct {
	*BaseInstaller
}

func NewOpenTofuInstaller(version string) *OpenTofuInstaller {
	return &OpenTofuInstaller{
		BaseInstaller: NewBaseInstaller(version, openTofuReleaseURL, "tofu", "zip"),
	}
}

func (oti *OpenTofuInstaller) GetInstallCommands() [][]string {
	url := fmt.Sprintf("%s/v%s/tofu_%s_linux_amd64.zip", oti.releaseURL, oti.version, oti.version)
	return oti.BaseInstaller.GetInstallCommands(url)
}

func (oti *OpenTofuInstaller) GetLatestVersion() (string, error) {
	// TODO: Implement logic to fetch the latest version from OpenTofu's releases page or API
	return "1.5.0", nil
}
