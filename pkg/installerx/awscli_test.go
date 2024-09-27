package installerx

import (
	"strings"
	"testing"
)

func TestGetAwsCliInstallCommand(t *testing.T) {
	tests := []struct {
		name         string
		architecture string
		wantContains []string
	}{
		{
			name:         "Default architecture",
			architecture: "",
			wantContains: []string{"https://awscli.amazonaws.com/awscli-exe-linux-x86_64.zip"},
		},
		{
			name:         "x86_64 architecture",
			architecture: "x86_64",
			wantContains: []string{"https://awscli.amazonaws.com/awscli-exe-linux-x86_64.zip"},
		},
		{
			name:         "aarch64 architecture",
			architecture: "aarch64",
			wantContains: []string{"https://awscli.amazonaws.com/awscli-exe-linux-arm64.zip"},
		},
		{
			name:         "arm64 architecture",
			architecture: "arm64",
			wantContains: []string{"https://awscli.amazonaws.com/awscli-exe-linux-arm64.zip"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetAwsCliInstallCommand(tt.architecture)
			for _, want := range tt.wantContains {
				if !strings.Contains(got, want) {
					t.Errorf("GetAwsCliInstallCommand() = %v, want to contain %v", got, want)
				}
			}
			if !strings.Contains(got, "curl -L") || !strings.Contains(got, "unzip awscliv2.zip") {
				t.Errorf("GetAwsCliInstallCommand() does not contain expected commands")
			}
		})
	}
}
