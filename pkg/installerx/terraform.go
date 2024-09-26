package installerx

import (
	"fmt"
	"path/filepath"
	"strings"
)

// TerraformInstallParams represents the parameters for installing Terraform
type TerraformInstallParams struct {
	// Version of Terraform to install (e.g., "1.0.0")
	Version string
	// InstallDir is the directory to install Terraform. If empty, defaults to DefaultInstallDir
	InstallDir string
}

// GetTerraformInstallCommand returns a string representing the command
// to install Terraform of a specific version.
//
// Parameters:
// - params: TerraformInstallParams struct containing installation parameters
//
// Returns:
// - A string representing the installation command
func GetTerraformInstallCommand(params TerraformInstallParams) string {
	if params.InstallDir == "" {
		params.InstallDir = DefaultInstallDir
	}

	installPath := filepath.Join(params.InstallDir, "terraform")

	command := fmt.Sprintf(`set -ex
curl -L https://releases.hashicorp.com/terraform/%[1]s/terraform_%[1]s_linux_amd64.zip -o /tmp/terraform.zip
unzip /tmp/terraform.zip -d /tmp
mv /tmp/terraform %[2]s
chmod +x %[2]s
rm /tmp/terraform.zip`, params.Version, installPath)

	return strings.TrimSpace(command)
}
