package installerx

import (
	"fmt"
	"path/filepath"
)

// TerragruntInstallParams represents the parameters for installing Terragrunt
type TerragruntInstallParams struct {
	// Version of Terragrunt to install (e.g., "0.38.0")
	Version string
	// EntryPoint for the command. If empty, defaults to "sh -c"
	EntryPoint string
	// InstallDir is the directory to install Terragrunt. If empty, defaults to "/usr/local/bin"
	InstallDir string
}

// GetTerragruntInstallCommand returns a slice of strings representing the command
// to install Terragrunt of a specific version.
//
// Parameters:
// - params: TerragruntInstallParams struct containing installation parameters
//
// Returns:
// - A slice of strings representing the installation command
func GetTerragruntInstallCommand(params TerragruntInstallParams) []string {
	if params.EntryPoint == "" {
		params.EntryPoint = "sh -c"
	}

	if params.InstallDir == "" {
		params.InstallDir = "/usr/local/bin"
	}

	installPath := filepath.Join(params.InstallDir, "terragrunt")

	command := fmt.Sprintf(`
set -ex
curl -L https://github.com/gruntwork-io/terragrunt/releases/download/v%[1]s/terragrunt_linux_amd64 -o %[2]s
chmod +x %[2]s
`, params.Version, installPath)

	return []string{params.EntryPoint, command}
}
