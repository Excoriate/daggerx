package installerx

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTerragruntInstaller_GetInstallCommands(t *testing.T) {
	tests := []struct {
		name     string
		version  string
		expected [][]string
	}{
		{
			name:    "Specific version",
			version: "0.67.4",
			expected: [][]string{
				{"mkdir", "-p", "/usr/local/bin"},
				{"curl", "-L", "-o", "/usr/local/bin/terragrunt", "https://github.com/gruntwork-io/terragrunt/releases/download/v0.67.4/terragrunt_linux_amd64"},
				{"chmod", "+x", "/usr/local/bin/terragrunt"},
				{"terragrunt", "--version"},
			},
		},
		{
			name:    "Version with 'v' prefix",
			version: "v0.67.4",
			expected: [][]string{
				{"mkdir", "-p", "/usr/local/bin"},
				{"curl", "-L", "-o", "/usr/local/bin/terragrunt", "https://github.com/gruntwork-io/terragrunt/releases/download/v0.67.4/terragrunt_linux_amd64"},
				{"chmod", "+x", "/usr/local/bin/terragrunt"},
				{"terragrunt", "--version"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			installer := NewTerragruntInstaller(tt.version)
			commands := installer.GetInstallCommands()
			assert.Equal(t, tt.expected, commands)
		})
	}
}

func TestTerragruntInstaller_GetLatestVersion(t *testing.T) {
	installer := NewTerragruntInstaller("latest")
	version, err := installer.GetLatestVersion()
	require.NoError(t, err)
	assert.NotEmpty(t, version)
	// Note: This test assumes that GetLatestVersion always returns "0.67.4" as per the current implementation
	assert.Equal(t, "0.67.4", version)
}
