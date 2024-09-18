package installerx

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestOpenTofuInstaller_GetInstallCommands(t *testing.T) {
	tests := []struct {
		name     string
		version  string
		expected [][]string
	}{
		{
			name:    "Specific version",
			version: "1.5.0",
			expected: [][]string{
				{"mkdir", "-p", "/usr/local/bin"},
				{"curl", "-L", "-o", "/tmp/tofu.zip", "https://github.com/opentofu/opentofu/releases/download/v1.5.0/tofu_1.5.0_linux_amd64.zip"},
				{"unzip", "-d", "/usr/local/bin", "/tmp/tofu.zip"},
				{"chmod", "+x", "/usr/local/bin/tofu"},
				{"tofu", "--version"},
				{"rm", "/tmp/tofu.zip"},
			},
		},
		{
			name:    "Version with 'v' prefix",
			version: "v1.5.0",
			expected: [][]string{
				{"mkdir", "-p", "/usr/local/bin"},
				{"curl", "-L", "-o", "/tmp/tofu.zip", "https://github.com/opentofu/opentofu/releases/download/v1.5.0/tofu_1.5.0_linux_amd64.zip"},
				{"unzip", "-d", "/usr/local/bin", "/tmp/tofu.zip"},
				{"chmod", "+x", "/usr/local/bin/tofu"},
				{"tofu", "--version"},
				{"rm", "/tmp/tofu.zip"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			installer := NewOpenTofuInstaller(tt.version)
			commands := installer.GetInstallCommands()
			assert.Equal(t, tt.expected, commands)
		})
	}
}

func TestOpenTofuInstaller_GetLatestVersion(t *testing.T) {
	installer := NewOpenTofuInstaller("latest")
	version, err := installer.GetLatestVersion()
	require.NoError(t, err)
	assert.NotEmpty(t, version)
	assert.Equal(t, "1.5.0", version)
}
