package installerx

import (
	"testing"
)

func TestGetTerraformInstallCommand(t *testing.T) {
	tests := []struct {
		name   string
		params TerraformInstallParams
		want   string
	}{
		{
			name: "Default parameters",
			params: TerraformInstallParams{
				Version: "1.0.0",
			},
			want: `set -ex
curl -L https://releases.hashicorp.com/terraform/1.0.0/terraform_1.0.0_linux_amd64.zip -o /tmp/terraform.zip
unzip /tmp/terraform.zip -d /tmp
mv /tmp/terraform /usr/local/bin/terraform
chmod +x /usr/local/bin/terraform
rm /tmp/terraform.zip`,
		},
		{
			name: "Custom install directory",
			params: TerraformInstallParams{
				Version:    "1.1.0",
				InstallDir: "/custom/bin",
			},
			want: `set -ex
curl -L https://releases.hashicorp.com/terraform/1.1.0/terraform_1.1.0_linux_amd64.zip -o /tmp/terraform.zip
unzip /tmp/terraform.zip -d /tmp
mv /tmp/terraform /custom/bin/terraform
chmod +x /custom/bin/terraform
rm /tmp/terraform.zip`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetTerraformInstallCommand(tt.params)
			if got != tt.want {
				t.Errorf("GetTerraformInstallCommand() = %v, want %v", got, tt.want)
			}
		})
	}
}
