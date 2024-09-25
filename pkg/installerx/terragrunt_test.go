package installerx

import (
	"reflect"
	"testing"
)

func TestGetTerragruntInstallCommand(t *testing.T) {
	tests := []struct {
		name   string
		params TerragruntInstallParams
		want   []string
	}{
		{
			name: "Default parameters",
			params: TerragruntInstallParams{
				Version: "0.38.0",
			},
			want: []string{"sh -c", `
set -ex
curl -L https://github.com/gruntwork-io/terragrunt/releases/download/v0.38.0/terragrunt_linux_amd64 -o /usr/local/bin/terragrunt
chmod +x /usr/local/bin/terragrunt
`},
		},
		{
			name: "Custom entry point and install directory",
			params: TerragruntInstallParams{
				Version:    "0.39.0",
				EntryPoint: "bash -c",
				InstallDir: "/custom/bin",
			},
			want: []string{"bash -c", `
set -ex
curl -L https://github.com/gruntwork-io/terragrunt/releases/download/v0.39.0/terragrunt_linux_amd64 -o /custom/bin/terragrunt
chmod +x /custom/bin/terragrunt
`},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetTerragruntInstallCommand(tt.params); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTerragruntInstallCommand() = %v, want %v", got, tt.want)
			}
		})
	}
}
