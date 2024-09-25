// Package installerx provides functionality for installing various binaries and tools.
package installerx

import (
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
	installDir string
}

// NewBaseInstaller creates a new BaseInstaller instance.
func NewBaseInstaller(version, releaseURL, binaryName, fileExt, installDir string) *BaseInstaller {
	if installDir == "" {
		installDir = "$HOME/bin"
	}
	return &BaseInstaller{
		version:    strings.TrimPrefix(version, "v"),
		releaseURL: releaseURL,
		binaryName: binaryName,
		fileExt:    fileExt,
		installDir: installDir,
	}
}

// GetInstallCommands returns the commands to install the binary.
func (bi *BaseInstaller) GetInstallCommands(url string) [][]string {
	commands := [][]string{
		{"mkdir", "-p", "$HOME/bin"},
		{"curl", "-L", "-o", "/tmp/" + bi.binaryName + "." + bi.fileExt, url},
	}
	if bi.fileExt == "zip" {
		commands = append(commands, []string{"unzip", "-d", "$HOME/bin", "/tmp/" + bi.binaryName + "." + bi.fileExt})
	} else {
		commands = append(commands, []string{"mv", "/tmp/" + bi.binaryName + "." + bi.fileExt, "$HOME/bin/" + bi.binaryName})
	}
	commands = append(commands,
		[]string{"chmod", "+x", "$HOME/bin/" + bi.binaryName},
		[]string{"$HOME/bin/" + bi.binaryName, "--version"},
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
	return container.WithEnvVariable("PATH", "$HOME/bin:$PATH")
}
