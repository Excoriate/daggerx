package installerx

import (
	"fmt"
	"path/filepath"
)

// OpenTofuInstallParams represents the parameters for installing OpenTofu
type OpenTofuInstallParams struct {
	// Version of OpenTofu to install (e.g., "1.6.0")
	Version string
	// EntryPoint for the command. If empty, defaults to "sh -c"
	EntryPoint string
	// InstallDir is the directory to install OpenTofu. If empty, defaults to "/usr/local/bin"
	InstallDir string
}

// GetOpenTofuInstallCommand returns a slice of strings representing the command
// to install OpenTofu of a specific version.
//
// Parameters:
// - params: OpenTofuInstallParams struct containing installation parameters
//
// Returns:
// - A slice of strings representing the installation command
func GetOpenTofuInstallCommand(params OpenTofuInstallParams) []string {
	if params.EntryPoint == "" {
		params.EntryPoint = "sh -c"
	}

	if params.InstallDir == "" {
		params.InstallDir = "/usr/local/bin"
	}

	installPath := filepath.Join(params.InstallDir, "opentofu")

	command := fmt.Sprintf(`
set -ex
echo "Downloading OpenTofu..."
curl -L https://github.com/opentofu/opentofu/releases/download/v%[1]s/tofu_%[1]s_linux_amd64.zip -o /tmp/opentofu.zip
unzip /tmp/opentofu.zip -d %[2]s
mv %[2]s/tofu %[3]s
chmod +x %[3]s
rm /tmp/opentofu.zip
echo "OpenTofu installation completed successfully"
%[3]s version
`, params.Version, params.InstallDir, installPath)

	return []string{params.EntryPoint, command}
}
