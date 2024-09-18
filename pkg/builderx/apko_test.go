package builderx

import (
	"reflect"
	"testing"

	"github.com/Excoriate/daggerx/pkg/fixtures"
)

func TestApkoBuilder(t *testing.T) {
	t.Run("NewApkoBuilder", func(t *testing.T) {
		builder := NewApkoBuilder()
		if builder == nil {
			t.Fatal("NewApkoBuilder returned nil")
		}
		if !reflect.DeepEqual(builder.architectures, []string{"x86_64", "aarch64"}) {
			t.Errorf("Default architectures not set correctly, got %v", builder.architectures)
		}
	})

	t.Run("WithConfigFile", func(t *testing.T) {
		builder := NewApkoBuilder().WithConfigFile("config.yaml")
		if builder.configFile != "config.yaml" {
			t.Errorf("Config file not set correctly, got %s", builder.configFile)
		}
	})

	t.Run("WithOutputImage", func(t *testing.T) {
		builder := NewApkoBuilder().WithOutputImage("my-image:latest")
		if builder.outputImage != "my-image:latest" {
			t.Errorf("Output image not set correctly, got %s", builder.outputImage)
		}
	})

	t.Run("WithOutputTarball", func(t *testing.T) {
		builder := NewApkoBuilder().WithOutputTarball("image.tar")
		if builder.outputTarball != "image.tar" {
			t.Errorf("Output tarball not set correctly, got %s", builder.outputTarball)
		}
	})

	t.Run("WithKeyring", func(t *testing.T) {
		builder := NewApkoBuilder().WithKeyring("/path/to/keyring.pub")
		if !reflect.DeepEqual(builder.keyringPaths, []string{"/path/to/keyring.pub"}) {
			t.Errorf("Keyring path not set correctly, got %v", builder.keyringPaths)
		}
	})

	t.Run("WithWolfiKeyring", func(t *testing.T) {
		builder := NewApkoBuilder().WithWolfiKeyring()
		if !builder.wolfiKeyring {
			t.Error("Wolfi keyring not enabled")
		}
	})

	t.Run("WithAlpineKeyring", func(t *testing.T) {
		builder := NewApkoBuilder().WithAlpineKeyring()
		if !builder.alpineKeyring {
			t.Error("Alpine keyring not enabled")
		}
	})

	t.Run("WithArchitecture", func(t *testing.T) {
		builder := NewApkoBuilder().WithArchitecture("arm64")
		if !reflect.DeepEqual(builder.architectures, []string{"x86_64", "aarch64", "arm64"}) {
			t.Errorf("Architecture not added correctly, got %v", builder.architectures)
		}
	})

	t.Run("WithCacheDir", func(t *testing.T) {
		builder := NewApkoBuilder().WithCacheDir("/tmp/cache")
		if builder.cacheDir != "/tmp/cache" {
			t.Errorf("Cache directory not set correctly, got %s", builder.cacheDir)
		}
	})

	t.Run("WithExtraArg", func(t *testing.T) {
		builder := NewApkoBuilder().WithExtraArg("--debug")
		if !reflect.DeepEqual(builder.extraArgs, []string{"--debug"}) {
			t.Errorf("Extra argument not added correctly, got %v", builder.extraArgs)
		}
	})

	t.Run("WithBuildArch", func(t *testing.T) {
		builder := NewApkoBuilder().WithBuildArch("amd64")
		if builder.buildArch != "amd64" {
			t.Errorf("Build arch not set correctly, got %s", builder.buildArch)
		}
	})

	t.Run("WithBuildContext", func(t *testing.T) {
		builder := NewApkoBuilder().WithBuildContext("/path/to/context")
		if builder.buildContext != "/path/to/context" {
			t.Errorf("Build context not set correctly, got %s", builder.buildContext)
		}
	})

	t.Run("WithDebug", func(t *testing.T) {
		builder := NewApkoBuilder().WithDebug()
		if !builder.debug {
			t.Error("Debug not enabled")
		}
	})

	t.Run("WithKeyringAppendPlaintext", func(t *testing.T) {
		builder := NewApkoBuilder().WithKeyringAppendPlaintext("/path/to/plaintext.key")
		if !reflect.DeepEqual(builder.keyringAppendPlaintext, []string{"/path/to/plaintext.key"}) {
			t.Errorf("Plaintext keyring not set correctly, got %v", builder.keyringAppendPlaintext)
		}
	})

	t.Run("WithNoNetwork", func(t *testing.T) {
		builder := NewApkoBuilder().WithNoNetwork()
		if !builder.noNetwork {
			t.Error("No network option not enabled")
		}
	})

	t.Run("WithRepositoryAppend", func(t *testing.T) {
		builder := NewApkoBuilder().WithRepositoryAppend("https://example.com/repo")
		if !reflect.DeepEqual(builder.repositoryAppend, []string{"https://example.com/repo"}) {
			t.Errorf("Repository append not set correctly, got %v", builder.repositoryAppend)
		}
	})

	t.Run("WithTimestamp", func(t *testing.T) {
		builder := NewApkoBuilder().WithTimestamp("2023-01-01T00:00:00Z")
		if builder.timestamp != "2023-01-01T00:00:00Z" {
			t.Errorf("Timestamp not set correctly, got %s", builder.timestamp)
		}
	})

	t.Run("BuildCommand", func(t *testing.T) {
		builder := NewApkoBuilder().
			WithConfigFile("config.yaml").
			WithOutputImage("my-image:latest").
			WithOutputTarball("image.tar").
			WithKeyring("/custom/keyring.pub").
			WithWolfiKeyring().
			WithAlpineKeyring().
			WithArchitecture("arm64").
			WithCacheDir("/tmp/cache").
			WithExtraArg("--custom-arg").
			WithBuildArch("amd64").
			WithBuildContext("/build/context").
			WithDebug().
			WithKeyringAppendPlaintext("/plaintext.key").
			WithNoNetwork().
			WithRepositoryAppend("https://example.com/repo").
			WithTimestamp("2023-01-01T00:00:00Z").
			WithTag("v1.0")

		expected := []string{
			"apko", "build",
			"--keyring-append", "/custom/keyring.pub",
			"--keyring-append", "/etc/apk/keys/wolfi-signing.rsa.pub",
			"--keyring-append", "/etc/apk/keys/alpine-devel@lists.alpinelinux.org-4a6a0840.rsa.pub",
			"--arch", "x86_64",
			"--arch", "aarch64",
			"--arch", "arm64",
			"--cache-dir", "/tmp/cache",
			"--build-arch", "amd64",
			"--build-context", "/build/context",
			"--debug",
			"--keyring-append-plaintext", "/plaintext.key",
			"--no-network",
			"--repository-append", "https://example.com/repo",
			"--timestamp", "2023-01-01T00:00:00Z",
			"--tag", "v1.0",
			"config.yaml",
			"my-image:latest",
			"image.tar",
			"--custom-arg",
		}

		cmd, err := builder.BuildCommand()
		if err != nil {
			t.Fatalf("BuildCommand returned an error: %v", err)
		}

		if !reflect.DeepEqual(cmd, expected) {
			t.Errorf("BuildCommand did not return expected command.\nGot:  %v\nWant: %v", cmd, expected)
		}
	})

	t.Run("BuildCommand_MissingConfigFile", func(t *testing.T) {
		builder := NewApkoBuilder().WithOutputImage("my-image:latest")
		_, err := builder.BuildCommand()
		if err == nil || err.Error() != "config file is required" {
			t.Errorf("Expected error for missing config file, got: %v", err)
		}
	})

	t.Run("BuildCommand_MissingOutputImage", func(t *testing.T) {
		builder := NewApkoBuilder().WithConfigFile("config.yaml")
		_, err := builder.BuildCommand()
		if err == nil || err.Error() != "output image name is required" {
			t.Errorf("Expected error for missing output image, got: %v", err)
		}
	})
}

func TestGetKeyringInfoForPreset(t *testing.T) {
	testCases := []struct {
		preset    string
		expectErr bool
		expected  KeyringInfo
	}{
		{
			preset:    "alpine",
			expectErr: false,
			expected: KeyringInfo{
				KeyURL:  "https://alpinelinux.org/keys/alpine-devel@lists.alpinelinux.org-4a6a0840.rsa.pub",
				KeyPath: "/etc/apk/keys/alpine-devel@lists.alpinelinux.org-4a6a0840.rsa.pub",
			},
		},
		{
			preset:    "wolfi",
			expectErr: false,
			expected: KeyringInfo{
				KeyURL:  "https://packages.wolfi.dev/os/wolfi-signing.rsa.pub",
				KeyPath: "/etc/apk/keys/wolfi-signing.rsa.pub",
			},
		},
		{
			preset:    "unsupported",
			expectErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.preset, func(t *testing.T) {
			info, err := GetKeyringInfoForPreset(tc.preset)

			if tc.expectErr {
				if err == nil {
					t.Errorf("Expected error for preset %s, but got none", tc.preset)
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error for preset %s: %v", tc.preset, err)
				}
				if info != tc.expected {
					t.Errorf("Expected keyring info %+v, but got %+v", tc.expected, info)
				}
			}
		})
	}
}

func TestGetCacheDir(t *testing.T) {
	t.Run("WithCustomMntPrefix", func(t *testing.T) {
		mntPrefix := "/mnt"
		expected := "/mnt/var/cache/apko"
		result := GetCacheDir(mntPrefix)
		if result != expected {
			t.Errorf("Expected cache dir %s, got %s", expected, result)
		}
	})

	t.Run("WithEmptyMntPrefix", func(t *testing.T) {
		mntPrefix := ""
		expected := fixtures.MntPrefix + "/var/cache/apko"
		result := GetCacheDir(mntPrefix)
		if result != expected {
			t.Errorf("Expected cache dir %s, got %s", expected, result)
		}
	})
}

func TestGetOutputTarPath(t *testing.T) {
	mntPrefix := "/mnt"
	expected := "/mnt/image.tar"
	result := GetOutputTarPath(mntPrefix)
	if result != expected {
		t.Errorf("Expected output tar path %s, got %s", expected, result)
	}
}

func TestApkoBuilder_WithKeyRingWolfi(t *testing.T) {
	builder := NewApkoBuilder()
	builder.WithKeyRingWolfi()

	expectedKeyPath := "/etc/apk/keys/wolfi-signing.rsa.pub"
	if len(builder.keyringPaths) != 1 || builder.keyringPaths[0] != expectedKeyPath {
		t.Errorf("Expected Wolfi keyring path %s, but got %v", expectedKeyPath, builder.keyringPaths)
	}
}

func TestApkoBuilder_WithKeyRingAlpine(t *testing.T) {
	builder := NewApkoBuilder()
	builder.WithKeyRingAlpine()

	expectedKeyPath := "/etc/apk/keys/alpine-devel@lists.alpinelinux.org-4a6a0840.rsa.pub"
	if len(builder.keyringPaths) != 1 || builder.keyringPaths[0] != expectedKeyPath {
		t.Errorf("Expected Alpine keyring path %s, but got %v", expectedKeyPath, builder.keyringPaths)
	}
}

func TestApkoBuilder_WithTag(t *testing.T) {
	t.Run("WithSpecificTag", func(t *testing.T) {
		builder := NewApkoBuilder().WithTag("v1.0")
		if !reflect.DeepEqual(builder.tags, []string{"v1.0"}) {
			t.Errorf("Tag not set correctly, got %v", builder.tags)
		}
	})

	t.Run("WithDefaultTag", func(t *testing.T) {
		builder := NewApkoBuilder().WithTag()
		if !reflect.DeepEqual(builder.tags, []string{"latest"}) {
			t.Errorf("Default tag not set correctly, got %v", builder.tags)
		}
	})
}
