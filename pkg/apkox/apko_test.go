package apkox

import (
	"reflect"
	"testing"
)

func TestApkoBuilder(t *testing.T) {
	t.Run("NewApkoBuilder", func(t *testing.T) {
		builder := NewApkoBuilder()
		if builder == nil {
			t.Fatal("NewApkoBuilder returned nil")
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
		if builder.buildArch != "arm64" {
			t.Errorf("Architecture not set correctly, got %v, want %v", builder.buildArch, "arm64")
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

	t.Run("WithAnnotations", func(t *testing.T) {
		builder := NewApkoBuilder().WithAnnotations(map[string]string{"key": "value"})
		if !reflect.DeepEqual(builder.annotations, map[string]string{"key": "value"}) {
			t.Errorf("Annotations not set correctly, got %v", builder.annotations)
		}
	})

	t.Run("WithBuildDate", func(t *testing.T) {
		builder := NewApkoBuilder().WithBuildDate("2023-01-01T00:00:00Z")
		if builder.buildDate != "2023-01-01T00:00:00Z" {
			t.Errorf("Build date not set correctly, got %s", builder.buildDate)
		}
	})

	t.Run("WithLockfile", func(t *testing.T) {
		builder := NewApkoBuilder().WithLockfile("/path/to/lockfile.json")
		if builder.lockfile != "/path/to/lockfile.json" {
			t.Errorf("Lockfile not set correctly, got %s", builder.lockfile)
		}
	})

	t.Run("WithOffline", func(t *testing.T) {
		builder := NewApkoBuilder().WithOffline()
		if !builder.offline {
			t.Error("Offline mode not enabled")
		}
	})

	t.Run("WithPackageAppend", func(t *testing.T) {
		builder := NewApkoBuilder().WithPackageAppend("pkg1", "pkg2")
		if !reflect.DeepEqual(builder.packageAppend, []string{"pkg1", "pkg2"}) {
			t.Errorf("Package append not set correctly, got %v", builder.packageAppend)
		}
	})

	t.Run("WithSBOM", func(t *testing.T) {
		builder := NewApkoBuilder().WithSBOM(false)
		if builder.sbom {
			t.Error("SBOM generation not disabled")
		}
	})

	t.Run("WithSBOMFormats", func(t *testing.T) {
		builder := NewApkoBuilder().WithSBOMFormats("spdx", "cyclonedx")
		if !reflect.DeepEqual(builder.sbomFormats, []string{"spdx", "cyclonedx"}) {
			t.Errorf("SBOM formats not set correctly, got %v", builder.sbomFormats)
		}
	})

	t.Run("WithSBOMPath", func(t *testing.T) {
		builder := NewApkoBuilder().WithSBOMPath("/path/to/sbom")
		if builder.sbomPath != "/path/to/sbom" {
			t.Errorf("SBOM path not set correctly, got %s", builder.sbomPath)
		}
	})

	t.Run("WithVCS", func(t *testing.T) {
		builder := NewApkoBuilder().WithVCS(false)
		if builder.vcs {
			t.Error("VCS detection not disabled")
		}
	})

	t.Run("WithLogLevel", func(t *testing.T) {
		builder := NewApkoBuilder().WithLogLevel("debug")
		if builder.logLevel != "debug" {
			t.Errorf("Log level not set correctly, got %s", builder.logLevel)
		}
	})

	t.Run("WithLogPolicy", func(t *testing.T) {
		builder := NewApkoBuilder().WithLogPolicy("builtin:stderr", "/tmp/log/foo")
		if !reflect.DeepEqual(builder.logPolicy, []string{"builtin:stderr", "/tmp/log/foo"}) {
			t.Errorf("Log policy not set correctly, got %v", builder.logPolicy)
		}
	})

	t.Run("WithWorkdir", func(t *testing.T) {
		builder := NewApkoBuilder().WithWorkdir("/path/to/workdir")
		if builder.workdir != "/path/to/workdir" {
			t.Errorf("Workdir not set correctly, got %s", builder.workdir)
		}
	})

	t.Run("BuildCommand", func(t *testing.T) {
		builder := NewApkoBuilder().
			WithConfigFile("config.yaml").
			WithOutputImage("my-image").
			WithTag("latest").
			WithOutputTarball("output.tar").
			WithCacheDir("/tmp/cache").
			WithKeyring("/custom/keyring.pub").
			WithArchitecture("amd64").
			WithSBOM(false).
			WithVCS(false)

		cmd, err := builder.BuildCommand()
		if err != nil {
			t.Fatalf("BuildCommand returned unexpected error: %v", err)
		}

		expected := []string{
			"apko", "build",
			"--cache-dir", "/tmp/cache",
			"--keyring-append", "/custom/keyring.pub",
			"--arch", "amd64",
			"--sbom=false",
			"--vcs=false",
			"config.yaml",
			"my-image:latest",
			"output.tar",
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

	t.Run("WithTag", func(t *testing.T) {
		builder := NewApkoBuilder().WithTag("v1.0.0")
		if builder.tag != "v1.0.0" {
			t.Errorf("Tag not set correctly, got %s", builder.tag)
		}
	})

	t.Run("BuildCommand", func(t *testing.T) {
		builder := NewApkoBuilder().
			WithConfigFile("config.yaml").
			WithOutputImage("my-image").
			WithTag("v1.0.0").
			WithOutputTarball("output.tar").
			WithCacheDir("/cache/dir").
			WithKeyring("/path/to/keyring.pub").
			WithArchitecture("aarch64").
			WithSBOM(false).
			WithVCS(false)

		cmd, err := builder.BuildCommand()
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}

		expected := []string{
			"apko", "build",
			"--cache-dir", "/cache/dir",
			"--keyring-append", "/path/to/keyring.pub",
			"--arch", "aarch64",
			"--sbom=false",
			"--vcs=false",
			"config.yaml",
			"my-image:v1.0.0",
			"output.tar",
		}

		if !reflect.DeepEqual(cmd, expected) {
			t.Errorf("Command mismatch.\nExpected: %v\nGot: %v", expected, cmd)
		}
	})

	t.Run("BuildCommand_DefaultTag", func(t *testing.T) {
		builder := NewApkoBuilder().
			WithConfigFile("config.yaml").
			WithOutputImage("my-image").
			WithOutputTarball("output.tar")

		cmd, err := builder.BuildCommand()
		if err != nil {
			t.Fatalf("BuildCommand returned unexpected error: %v", err)
		}

		expectedImageRef := "my-image:latest"
		imageRefFound := false
		for i, arg := range cmd {
			if arg == expectedImageRef {
				imageRefFound = true
				// Verify the position - should be second to last argument
				if i != len(cmd)-2 {
					t.Errorf("Image reference in wrong position. Got position %d, want second to last", i)
				}
				break
			}
		}

		if !imageRefFound {
			t.Errorf("Default tag not applied correctly. Expected to find %s in command %v", expectedImageRef, cmd)
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
	mntPrefix := "/mnt"
	expected := "/mnt/var/cache/apko"
	result := GetCacheDir(mntPrefix)
	if result != expected {
		t.Errorf("Expected cache dir %s, got %s", expected, result)
	}
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

// TestGetApkoConfigOrPreset tests the GetApkoConfigOrPreset function
func TestGetApkoConfigOrPreset(t *testing.T) {
	tests := []struct {
		name      string
		mntPrefix string
		cfgFile   string
		want      string
		wantErr   bool
	}{
		{
			name:      "Valid config file with .yaml extension",
			mntPrefix: "/mnt",
			cfgFile:   "config.yaml",
			want:      "config.yaml",
			wantErr:   false,
		},
		{
			name:      "Valid config file with .yml extension",
			mntPrefix: "/mnt",
			cfgFile:   "config.yml",
			want:      "config.yml",
			wantErr:   false,
		},
		{
			name:      "Empty mntPrefix",
			mntPrefix: "",
			cfgFile:   "config.yaml",
			want:      "config.yaml",
			wantErr:   false,
		},
		{
			name:      "Empty config file",
			mntPrefix: "/mnt",
			cfgFile:   "",
			want:      "",
			wantErr:   true,
		},
		{
			name:      "Config file without extension",
			mntPrefix: "/mnt",
			cfgFile:   "config",
			want:      "",
			wantErr:   true,
		},
		{
			name:      "Config file with invalid extension",
			mntPrefix: "/mnt",
			cfgFile:   "config.txt",
			want:      "",
			wantErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetApkoConfigOrPreset(tt.mntPrefix, tt.cfgFile)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetApkoConfigOrPreset() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetApkoConfigOrPreset() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestApkoBuilderCommand(t *testing.T) {
	builder := NewApkoBuilder().
		WithConfigFile("config.yaml").
		WithOutputImage("my-image").
		WithTag("v1.0.0").
		WithOutputTarball("output.tar").
		WithCacheDir("/cache/dir").
		WithKeyring("/path/to/keyring.pub").
		WithArchitecture("aarch64").
		WithSBOM(false).
		WithVCS(false)

	cmd, err := builder.BuildCommand()
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	expected := []string{
		"apko", "build",
		"--cache-dir", "/cache/dir",
		"--keyring-append", "/path/to/keyring.pub",
		"--arch", "aarch64",
		"--sbom=false",
		"--vcs=false",
		"config.yaml",
		"my-image:v1.0.0",
		"output.tar",
	}

	if !reflect.DeepEqual(cmd, expected) {
		t.Errorf("Command mismatch.\nExpected: %v\nGot: %v", expected, cmd)
	}
}
