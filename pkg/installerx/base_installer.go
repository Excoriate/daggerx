package installerx

import (
	"strings"

	"dagger.io/dagger"
)

// BaseInstaller provides common functionality for installers.
type BaseInstaller struct {
	version     string
	releaseURL  string
	binaryName  string
	archiveType string
}

// NewBaseInstaller creates a new BaseInstaller instance.
func NewBaseInstaller(version, releaseURL, binaryName, archiveType string) BaseInstaller {
	return BaseInstaller{
		version:     strings.TrimPrefix(version, "v"),
		releaseURL:  releaseURL,
		binaryName:  binaryName,
		archiveType: archiveType,
	}
}

// GetInstallCommands returns the commands to install the binary.
func (bi *BaseInstaller) GetInstallCommands(url string) [][]string {
	commands := [][]string{
		{"mkdir", "-p", "/usr/local/bin"},
		{"curl", "-L", "-o", "/tmp/" + bi.binaryName + "." + bi.archiveType, url},
	}

	if bi.archiveType == "zip" {
		commands = append(commands, []string{"unzip", "-d", "/usr/local/bin", "/tmp/" + bi.binaryName + "." + bi.archiveType})
	} else {
		commands = append(commands, []string{"mv", "/tmp/" + bi.binaryName + "." + bi.archiveType, "/usr/local/bin/" + bi.binaryName})
	}

	commands = append(commands,
		[]string{"chmod", "+x", "/usr/local/bin/" + bi.binaryName},
		[]string{bi.binaryName, "--version"},
	)

	if bi.archiveType == "zip" {
		commands = append(commands, []string{"rm", "/tmp/" + bi.binaryName + "." + bi.archiveType})
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
