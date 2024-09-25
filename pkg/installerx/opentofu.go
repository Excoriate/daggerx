package installerx

import (
	"fmt"
)

// GetOpenTofuInstallCommand returns a slice of strings representing the command
// to install OpenTofu of a specific version.
//
// Parameters:
// - version: The version of OpenTofu to install (e.g., "1.6.0")
// - entryPoint: Optional entry point for the command. If empty, defaults to "sh -c"
//
// Returns:
// - A slice of strings representing the installation command
func GetOpenTofuInstallCommand(version string, entryPoint string) []string {
	if entryPoint == "" {
		entryPoint = "sh -c"
	}

	command := fmt.Sprintf(`
set -ex
echo "Downloading OpenTofu..."
curl -L https://github.com/opentofu/opentofu/releases/download/v%[1]s/tofu_%[1]s_linux_amd64.zip -o /tmp/opentofu.zip
unzip /tmp/opentofu.zip -d /usr/local/bin
mv /usr/local/bin/tofu /usr/local/bin/opentofu
chmod +x /usr/local/bin/opentofu
rm /tmp/opentofu.zip
echo "OpenTofu installation completed successfully"
opentofu version
`, version)

	return []string{entryPoint, command}
}
