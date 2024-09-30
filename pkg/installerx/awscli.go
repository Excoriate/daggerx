// Package installerx provides utilities for installing various software packages
// in a streamlined and automated manner. This package is particularly useful for
// setting up development environments, CI/CD pipelines, and other automated workflows
// that require the installation of specific software components.
//
// The installerx package includes functions to generate shell commands for downloading
// and installing software packages, ensuring compatibility with different system
// architectures and handling any necessary conversions or adjustments.
//
// One of the key features of this package is the ability to generate installation
// commands for the AWS CLI, a powerful tool for managing AWS services from the command
// line. The package supports different system architectures, including "x86_64" and
// "aarch64" (which is converted to "arm64" for compatibility).
//
// Example usage:
//
//	package main
//
//	import (
//		"fmt"
//		"installerx"
//	)
//
//	func main() {
//		// Generate the installation command for the AWS CLI for the default architecture (x86_64)
//		command := installerx.GetAwsCliInstallCommand("")
//		fmt.Println(command)
//
//		// Generate the installation command for the AWS CLI for the "aarch64" architecture
//		command = installerx.GetAwsCliInstallCommand("aarch64")
//		fmt.Println(command)
//	}
//
// This package is designed to be easy to use and integrate into existing Go projects,
// providing a simple and efficient way to automate software installations.
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
