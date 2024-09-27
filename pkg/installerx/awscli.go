package installerx

import (
	"fmt"
	"strings"
)

// GetAwsCliInstallCommand generates a shell command to download and install the AWS CLI
// for the specified architecture. If the architecture is not provided, it defaults to "x86_64".
// The function also handles the conversion of "aarch64" to "arm64" for compatibility.
//
// Parameters:
//
//	architecture - a string representing the system architecture (e.g., "x86_64", "aarch64")
//
// Returns:
//
//	A string containing the shell command to install the AWS CLI.
func GetAwsCliInstallCommand(architecture string) string {
	if architecture == "" {
		architecture = "x86_64"
	}

	if architecture == "aarch64" {
		architecture = "arm64"
	}

	url := fmt.Sprintf("https://awscli.amazonaws.com/awscli-exe-linux-%s.zip", architecture)

	command := fmt.Sprintf(`set -ex
curl -L %[1]s -o awscliv2.zip
unzip awscliv2.zip
sudo ./aws/install
rm -rf awscliv2.zip aws
`, url)

	return strings.TrimSpace(command)
}
