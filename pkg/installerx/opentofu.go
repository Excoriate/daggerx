package installerx

import (
	"fmt"
	"path/filepath"
	"strings"
)

// OpenTofuInstallParams represents the parameters for installing OpenTofu
type OpenTofuInstallParams struct {
	// Version of OpenTofu to install (e.g., "1.6.0")
	Version string
	// InstallDir is the directory to install OpenTofu. If empty, defaults to DefaultInstallDir
	InstallDir string
}

// GetOpenTofuInstallCommand returns a string representing the command
// to install OpenTofu of a specific version.
//
// Parameters:
// - params: OpenTofuInstallParams struct containing installation parameters
//
// Returns:
// - A string representing the installation command
func GetOpenTofuInstallCommand(params OpenTofuInstallParams) string {
	if params.InstallDir == "" {
		params.InstallDir = DefaultInstallDir
	}

	installPath := filepath.Join(params.InstallDir, "opentofu")

	command := fmt.Sprintf(`set -ex
echo "Downloading OpenTofu..."
curl -L "https://github.com/opentofu/opentofu/releases/download/v%[1]s/tofu_%[1]s_linux_amd64.zip" -o /tmp/opentofu.zip
unzip /tmp/opentofu.zip -d /tmp
mv /tmp/tofu %[2]s
chmod +x %[2]s
rm /tmp/opentofu.zip
echo "OpenTofu installation completed successfully"
%[2]s version`, params.Version, installPath)

	return strings.TrimSpace(command)
}
