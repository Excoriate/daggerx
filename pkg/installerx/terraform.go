package installerx

import (
	"fmt"
)

// GetTerraformInstallCommand returns a slice of strings representing the command
// to install Terraform of a specific version.
//
// Parameters:
// - version: The version of Terraform to install (e.g., "1.0.0")
// - entryPoint: Optional entry point for the command. If empty, defaults to "sh -c"
//
// Returns:
// - A slice of strings representing the installation command
func GetTerraformInstallCommand(version string, entryPoint string) []string {
	if entryPoint == "" {
		entryPoint = "sh -c"
	}

	command := fmt.Sprintf(`
set -ex
curl -L https://releases.hashicorp.com/terraform/%[1]s/terraform_%[1]s_linux_amd64.zip -o /tmp/terraform.zip
unzip /tmp/terraform.zip -d /usr/local/bin
chmod +x /usr/local/bin/terraform
rm /tmp/terraform.zip
`, version)

	return []string{entryPoint, command}
}
