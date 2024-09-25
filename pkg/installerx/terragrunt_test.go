package installerx

import (
	"reflect"
	"testing"
)

func TestGetTerragruntInstallCommand(t *testing.T) {
	tests := []struct {
		name       string
		version    string
		entryPoint string
		want       []string
	}{
		{
			name:       "Default entry point",
			version:    "0.38.0",
			entryPoint: "",
			want: []string{"sh -c", `
set -ex
curl -L https://github.com/gruntwork-io/terragrunt/releases/download/v0.38.0/terragrunt_linux_amd64 -o /usr/local/bin/terragrunt
chmod +x /usr/local/bin/terragrunt
`},
		},
		{
			name:       "Custom entry point",
			version:    "0.39.0",
			entryPoint: "bash -c",
			want: []string{"bash -c", `
set -ex
curl -L https://github.com/gruntwork-io/terragrunt/releases/download/v0.39.0/terragrunt_linux_amd64 -o /usr/local/bin/terragrunt
chmod +x /usr/local/bin/terragrunt
`},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetTerragruntInstallCommand(tt.version, tt.entryPoint); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTerragruntInstallCommand() = %v, want %v", got, tt.want)
			}
		})
	}
}
