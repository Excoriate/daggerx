package installerx

import (
	"testing"
)

func TestGetTerragruntInstallCommand(t *testing.T) {
	tests := []struct {
		name   string
		params TerragruntInstallParams
		want   string
	}{
		{
			name: "Default parameters",
			params: TerragruntInstallParams{
				Version: "0.38.0",
			},
			want: `set -ex
curl -L https://github.com/gruntwork-io/terragrunt/releases/download/v0.38.0/terragrunt_linux_amd64 -o /usr/local/bin/terragrunt
chmod +x /usr/local/bin/terragrunt`,
		},
		{
			name: "Custom install directory",
			params: TerragruntInstallParams{
				Version:    "0.39.0",
				InstallDir: "/custom/bin",
			},
			want: `set -ex
curl -L https://github.com/gruntwork-io/terragrunt/releases/download/v0.39.0/terragrunt_linux_amd64 -o /custom/bin/terragrunt
chmod +x /custom/bin/terragrunt`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetTerragruntInstallCommand(tt.params)
			if got != tt.want {
				t.Errorf("GetTerragruntInstallCommand() = %v, want %v", got, tt.want)
			}
		})
	}
}
