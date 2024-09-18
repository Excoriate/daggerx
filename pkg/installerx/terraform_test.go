package installerx

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTerraformInstaller_GetInstallCommands(t *testing.T) {
	tests := []struct {
		name     string
		version  string
		expected [][]string
	}{
		{
			name:    "Specific version",
			version: "1.9.4",
			expected: [][]string{
				{"mkdir", "-p", "/usr/local/bin"},
				{"curl", "-L", "-o", "/tmp/terraform.zip", "https://releases.hashicorp.com/terraform/1.9.4/terraform_1.9.4_linux_amd64.zip"},
				{"unzip", "-d", "/usr/local/bin", "/tmp/terraform.zip"},
				{"chmod", "+x", "/usr/local/bin/terraform"},
				{"terraform", "--version"},
				{"rm", "/tmp/terraform.zip"},
			},
		},
		{
			name:    "Version with 'v' prefix",
			version: "v1.9.4",
			expected: [][]string{
				{"mkdir", "-p", "/usr/local/bin"},
				{"curl", "-L", "-o", "/tmp/terraform.zip", "https://releases.hashicorp.com/terraform/1.9.4/terraform_1.9.4_linux_amd64.zip"},
				{"unzip", "-d", "/usr/local/bin", "/tmp/terraform.zip"},
				{"chmod", "+x", "/usr/local/bin/terraform"},
				{"terraform", "--version"},
				{"rm", "/tmp/terraform.zip"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			installer := NewTerraformInstaller(tt.version)
			commands := installer.GetInstallCommands()
			assert.Equal(t, tt.expected, commands)
		})
	}
}

func TestTerraformInstaller_GetLatestVersion(t *testing.T) {
	installer := NewTerraformInstaller("latest")
	version, err := installer.GetLatestVersion()
	require.NoError(t, err)
	assert.NotEmpty(t, version)
	// Note: This test assumes that GetLatestVersion always returns "1.9.4" as per the current implementation
	assert.Equal(t, "1.9.4", version)
}
