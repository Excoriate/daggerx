package installerx

import (
	"testing"
)

func TestGetOpenTofuInstallCommand(t *testing.T) {
	tests := []struct {
		name   string
		params OpenTofuInstallParams
		want   string
	}{
		{
			name: "Default parameters",
			params: OpenTofuInstallParams{
				Version: "1.6.0",
			},
			want: `set -ex
echo "Downloading OpenTofu..."
curl -L "https://github.com/opentofu/opentofu/releases/download/v1.6.0/tofu_1.6.0_linux_amd64.zip" -o /tmp/opentofu.zip
unzip /tmp/opentofu.zip -d /tmp
mv /tmp/tofu /usr/local/bin/opentofu
chmod +x /usr/local/bin/opentofu
rm /tmp/opentofu.zip
echo "OpenTofu installation completed successfully"
/usr/local/bin/opentofu version`,
		},
		{
			name: "Custom install directory",
			params: OpenTofuInstallParams{
				Version:    "1.7.0",
				InstallDir: "/custom/bin",
			},
			want: `set -ex
echo "Downloading OpenTofu..."
curl -L "https://github.com/opentofu/opentofu/releases/download/v1.7.0/tofu_1.7.0_linux_amd64.zip" -o /tmp/opentofu.zip
unzip /tmp/opentofu.zip -d /tmp
mv /tmp/tofu /custom/bin/opentofu
chmod +x /custom/bin/opentofu
rm /tmp/opentofu.zip
echo "OpenTofu installation completed successfully"
/custom/bin/opentofu version`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetOpenTofuInstallCommand(tt.params)
			if got != tt.want {
				t.Errorf("GetOpenTofuInstallCommand() = %v, want %v", got, tt.want)
			}
		})
	}
}
