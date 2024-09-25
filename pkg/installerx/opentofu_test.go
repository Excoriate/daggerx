package installerx

import (
	"reflect"
	"testing"
)

func TestGetOpenTofuInstallCommand(t *testing.T) {
	tests := []struct {
		name       string
		version    string
		entryPoint string
		want       []string
	}{
		{
			name:       "Default entry point",
			version:    "1.6.0",
			entryPoint: "",
			want: []string{"sh -c", `
set -ex
echo "Downloading OpenTofu..."
curl -L https://github.com/opentofu/opentofu/releases/download/v1.6.0/tofu_1.6.0_linux_amd64.zip -o /tmp/opentofu.zip
unzip /tmp/opentofu.zip -d /usr/local/bin
mv /usr/local/bin/tofu /usr/local/bin/opentofu
chmod +x /usr/local/bin/opentofu
rm /tmp/opentofu.zip
echo "OpenTofu installation completed successfully"
opentofu version
`},
		},
		{
			name:       "Custom entry point",
			version:    "1.7.0",
			entryPoint: "bash -c",
			want: []string{"bash -c", `
set -ex
echo "Downloading OpenTofu..."
curl -L https://github.com/opentofu/opentofu/releases/download/v1.7.0/tofu_1.7.0_linux_amd64.zip -o /tmp/opentofu.zip
unzip /tmp/opentofu.zip -d /usr/local/bin
mv /usr/local/bin/tofu /usr/local/bin/opentofu
chmod +x /usr/local/bin/opentofu
rm /tmp/opentofu.zip
echo "OpenTofu installation completed successfully"
opentofu version
`},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetOpenTofuInstallCommand(tt.version, tt.entryPoint); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetOpenTofuInstallCommand() = %v, want %v", got, tt.want)
			}
		})
	}
}
