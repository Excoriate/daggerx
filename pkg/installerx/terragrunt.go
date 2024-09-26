package installerx

import (
	"fmt"
	"path/filepath"
	"strings"
)

// TerragruntInstallParams represents the parameters for installing Terragrunt
type TerragruntInstallParams struct {
	// Version of Terragrunt to install (e.g., "0.38.0")
	Version string
	// InstallDir is the directory to install Terragrunt. If empty, defaults to "/usr/local/bin"
	InstallDir string
}

// GetTerragruntInstallCommand returns a string representing the command
// to install Terragrunt of a specific version.
//
// Parameters:
// - params: TerragruntInstallParams struct containing installation parameters
//
// Returns:
// - A string representing the installation command
func GetTerragruntInstallCommand(params TerragruntInstallParams) string {
	if params.InstallDir == "" {
		params.InstallDir = "/usr/local/bin"
	}

	installPath := filepath.Join(params.InstallDir, "terragrunt")

	command := fmt.Sprintf(`set -ex
curl -L https://github.com/gruntwork-io/terragrunt/releases/download/v%s/terragrunt_linux_amd64 -o %s
chmod +x %s`, params.Version, installPath, installPath)

	return strings.TrimSpace(command)
}
