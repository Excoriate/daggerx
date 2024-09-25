package installerx

import (
	"reflect"
	"testing"
)

func TestGetTerraformInstallCommand(t *testing.T) {
	tests := []struct {
		name       string
		version    string
		entryPoint string
		want       []string
	}{
		{
			name:       "Default entry point",
			version:    "1.0.0",
			entryPoint: "",
			want: []string{"sh -c", `
set -ex
curl -L https://releases.hashicorp.com/terraform/1.0.0/terraform_1.0.0_linux_amd64.zip -o /tmp/terraform.zip
unzip /tmp/terraform.zip -d /usr/local/bin
chmod +x /usr/local/bin/terraform
rm /tmp/terraform.zip
`},
		},
		{
			name:       "Custom entry point",
			version:    "1.1.0",
			entryPoint: "bash -c",
			want: []string{"bash -c", `
set -ex
curl -L https://releases.hashicorp.com/terraform/1.1.0/terraform_1.1.0_linux_amd64.zip -o /tmp/terraform.zip
unzip /tmp/terraform.zip -d /usr/local/bin
chmod +x /usr/local/bin/terraform
rm /tmp/terraform.zip
`},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetTerraformInstallCommand(tt.version, tt.entryPoint); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTerraformInstallCommand() = %v, want %v", got, tt.want)
			}
		})
	}
}
