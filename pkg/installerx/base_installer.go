// Package installerx provides functionality for installing various binaries and tools.
package installerx

import (
	"fmt"
	"strings"

	"dagger.io/dagger"
)

// BaseInstaller implements the Installer interface and provides common functionality
// for installing binaries across different platforms and package formats.
type BaseInstaller struct {
	version    string
	releaseURL string
	binaryName string
	fileExt    string
}

// NewBaseInstaller creates a new BaseInstaller instance.
func NewBaseInstaller(version, releaseURL, binaryName, fileExt string) *BaseInstaller {
	return &BaseInstaller{
		version:    strings.TrimPrefix(version, "v"),
		releaseURL: releaseURL,
		binaryName: binaryName,
		fileExt:    fileExt,
	}
}

// GetInstallCommands returns the commands to install the binary.
func (bi *BaseInstaller) GetInstallCommands(url string) [][]string {
	commands := [][]string{
		{"mkdir", "-p", "/usr/local/bin"},
		{"curl", "-L", "-o", "/tmp/" + bi.binaryName + "." + bi.fileExt, url},
	}

	if bi.fileExt == "zip" {
		commands = append(commands, []string{"unzip", "-d", "/usr/local/bin", "/tmp/" + bi.binaryName + "." + bi.fileExt})
	} else {
		commands = append(commands, []string{"mv", "/tmp/" + bi.binaryName + "." + bi.fileExt, "/usr/local/bin/" + bi.binaryName})
	}

	commands = append(commands,
		[]string{"chmod", "+x", "/usr/local/bin/" + bi.binaryName},
		[]string{bi.binaryName, "--version"},
	)

	if bi.fileExt == "zip" {
		commands = append(commands, []string{"rm", "/tmp/" + bi.binaryName + "." + bi.fileExt})
	}

	return commands
}

// Install performs the installation on a Dagger container.
func (bi *BaseInstaller) Install(container *dagger.Container, commands [][]string) *dagger.Container {
	for _, cmd := range commands {
		container = container.WithExec(cmd)
	}
	return container
}

// GetLatestVersion returns the latest version of the binary.
func (bi *BaseInstaller) GetLatestVersion() (string, error) {
	// Default implementation, should be overridden by specific installers
	return "", fmt.Errorf("GetLatestVersion not implemented for BaseInstaller")
}
