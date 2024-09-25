package installerx

import (
	"fmt"
	"path/filepath"
)

// TerraformInstallParams represents the parameters for installing Terraform
type TerraformInstallParams struct {
	// Version of Terraform to install (e.g., "1.0.0")
	Version string
	// EntryPoint for the command. If empty, defaults to "sh -c"
	EntryPoint string
	// InstallDir is the directory to install Terraform. If empty, defaults to "/usr/local/bin"
	InstallDir string
}

// GetTerraformInstallCommand returns a slice of strings representing the command
// to install Terraform of a specific version.
//
// Parameters:
// - params: TerraformInstallParams struct containing installation parameters
//
// Returns:
// - A slice of strings representing the installation command
func GetTerraformInstallCommand(params TerraformInstallParams) []string {
	if params.EntryPoint == "" {
		params.EntryPoint = "sh -c"
	}

	if params.InstallDir == "" {
		params.InstallDir = "/usr/local/bin"
	}

	installPath := filepath.Join(params.InstallDir, "terraform")

	command := fmt.Sprintf(`
set -ex
curl -L https://releases.hashicorp.com/terraform/%[1]s/terraform_%[1]s_linux_amd64.zip -o /tmp/terraform.zip
unzip /tmp/terraform.zip -d %[2]s
chmod +x %[3]s
rm /tmp/terraform.zip
`, params.Version, params.InstallDir, installPath)

	return []string{params.EntryPoint, command}
}
