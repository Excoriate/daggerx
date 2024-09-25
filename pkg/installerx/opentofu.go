package installerx

import (
	"fmt"
)

// openTofuReleaseURL is the base URL for OpenTofu releases on GitHub.
const openTofuReleaseURL = "https://github.com/opentofu/opentofu/releases/download"

// OpenTofuInstaller represents an installer for OpenTofu.
// It embeds BaseInstaller to inherit common installation functionality.
type OpenTofuInstaller struct {
	*BaseInstaller
}

// NewOpenTofuInstaller creates and returns a new OpenTofuInstaller instance.
// It initializes the embedded BaseInstaller with OpenTofu-specific parameters.
//
// Parameters:
//   - version: The version of OpenTofu to install.
//
// Returns:
//   - *OpenTofuInstaller: A pointer to the newly created OpenTofuInstaller.
func NewOpenTofuInstaller(version string) *OpenTofuInstaller {
	return &OpenTofuInstaller{
		BaseInstaller: NewBaseInstaller(version, openTofuReleaseURL, "tofu", "zip", "$HOME/bin"),
	}
}

// GetInstallCommands returns a slice of command slices needed to install OpenTofu.
// It generates the download URL for the specific version and architecture,
// then delegates to the BaseInstaller to generate the actual install commands.
//
// Returns:
//   - [][]string: A slice of command slices, where each inner slice represents
//     a command with its arguments.
func (oti *OpenTofuInstaller) GetInstallCommands() [][]string {
	url := fmt.Sprintf("%s/v%s/tofu_%s_linux_amd64.zip", oti.releaseURL, oti.version, oti.version)
	return oti.BaseInstaller.GetInstallCommands(url)
}

// GetLatestVersion retrieves the latest version of OpenTofu available.
//
// Returns:
//   - string: The latest version number.
//   - error: An error if the version retrieval fails.
//
// Note: This is currently a placeholder implementation.
func (oti *OpenTofuInstaller) GetLatestVersion() (string, error) {
	// TODO: Implement logic to fetch the latest version from OpenTofu's releases page or API
	return "1.5.0", nil
}
