package installerx

import (
	"strings"
	"testing"
)

func TestGetGitHubAssetInstallCommand(t *testing.T) {
	tests := []struct {
		name    string
		params  GitHubAssetParams
		want    string
		wantErr bool
	}{
		{
			name: "terraform-docs tarball with pattern",
			params: GitHubAssetParams{
				Owner:        "terraform-docs",
				Repo:         "terraform-docs",
				Version:      "0.19.0",
				AssetPattern: "terraform-docs-v{version}-{os}-{arch}.tar.gz",
				BinaryName:   "terraform-docs",
			},
			want: `set -ex
curl -fL https://github.com/terraform-docs/terraform-docs/releases/download/v0.19.0/terraform-docs-v0.19.0-linux-amd64.tar.gz -o /tmp/terraform-docs-v0.19.0-linux-amd64.tar.gz
cd /tmp && tar -xzf terraform-docs-v0.19.0-linux-amd64.tar.gz
mv /tmp/terraform-docs /usr/local/bin/terraform-docs
chmod +x /usr/local/bin/terraform-docs
rm -f /tmp/terraform-docs-v0.19.0-linux-amd64.tar.gz`,
			wantErr: false,
		},
		{
			name: "direct binary download with pattern",
			params: GitHubAssetParams{
				Owner:        "gruntwork-io",
				Repo:         "terragrunt",
				Version:      "v0.38.0",
				AssetPattern: "terragrunt_{os}_{arch}",
				BinaryName:   "terragrunt",
			},
			want: `set -ex
curl -fL https://github.com/gruntwork-io/terragrunt/releases/download/v0.38.0/terragrunt_linux_amd64 -o /usr/local/bin/terragrunt
chmod +x /usr/local/bin/terragrunt`,
			wantErr: false,
		},
		{
			name: "direct asset name without pattern",
			params: GitHubAssetParams{
				Owner:      "cli",
				Repo:       "cli",
				Version:    "v2.0.0",
				AssetName:  "gh_2.0.0_linux_amd64.tar.gz",
				BinaryName: "gh",
			},
			want: `set -ex
curl -fL https://github.com/cli/cli/releases/download/v2.0.0/gh_2.0.0_linux_amd64.tar.gz -o /tmp/gh_2.0.0_linux_amd64.tar.gz
cd /tmp && tar -xzf gh_2.0.0_linux_amd64.tar.gz
mv /tmp/gh /usr/local/bin/gh
chmod +x /usr/local/bin/gh
rm -f /tmp/gh_2.0.0_linux_amd64.tar.gz`,
			wantErr: false,
		},
		{
			name: "custom extract path",
			params: GitHubAssetParams{
				Owner:       "cli",
				Repo:        "cli",
				Version:     "v2.0.0",
				AssetName:   "gh_2.0.0_linux_amd64.tar.gz",
				BinaryName:  "gh",
				ExtractPath: "gh_2.0.0_linux_amd64/bin/gh",
			},
			want: `set -ex
curl -fL https://github.com/cli/cli/releases/download/v2.0.0/gh_2.0.0_linux_amd64.tar.gz -o /tmp/gh_2.0.0_linux_amd64.tar.gz
cd /tmp && tar -xzf gh_2.0.0_linux_amd64.tar.gz
mv /tmp/gh_2.0.0_linux_amd64/bin/gh /usr/local/bin/gh
chmod +x /usr/local/bin/gh
rm -f /tmp/gh_2.0.0_linux_amd64.tar.gz`,
			wantErr: false,
		},
		{
			name: "tgz extension",
			params: GitHubAssetParams{
				Owner:      "example",
				Repo:       "tool",
				Version:    "1.0.0",
				AssetName:  "tool-1.0.0.tgz",
				BinaryName: "tool",
			},
			want: `set -ex
curl -fL https://github.com/example/tool/releases/download/v1.0.0/tool-1.0.0.tgz -o /tmp/tool-1.0.0.tgz
cd /tmp && tar -xzf tool-1.0.0.tgz
mv /tmp/tool /usr/local/bin/tool
chmod +x /usr/local/bin/tool
rm -f /tmp/tool-1.0.0.tgz`,
			wantErr: false,
		},
		{
			name: "missing both pattern and name",
			params: GitHubAssetParams{
				Owner:      "terraform-docs",
				Repo:       "terraform-docs",
				Version:    "0.19.0",
				BinaryName: "terraform-docs",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetGitHubAssetInstallCommand(tt.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetGitHubAssetInstallCommand() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil && !strings.EqualFold(got, tt.want) {
				t.Errorf("GetGitHubAssetInstallCommand() = %v, want %v", got, tt.want)
			}
		})
	}
}
