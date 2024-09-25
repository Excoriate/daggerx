package installerx

import (
	"fmt"
)

// GetTerragruntInstallCommand returns a slice of strings representing the command
// to install Terragrunt of a specific version.
//
// Parameters:
// - version: The version of Terragrunt to install (e.g., "0.38.0")
// - entryPoint: Optional entry point for the command. If empty, defaults to "sh -c"
//
// Returns:
// - A slice of strings representing the installation command
func GetTerragruntInstallCommand(version string, entryPoint string) []string {
	if entryPoint == "" {
		entryPoint = "sh -c"
	}

	command := fmt.Sprintf(`
set -ex
curl -L https://github.com/gruntwork-io/terragrunt/releases/download/v%[1]s/terragrunt_linux_amd64 -o /usr/local/bin/terragrunt
chmod +x /usr/local/bin/terragrunt
`, version)

	return []string{entryPoint, command}
}
