package builderx

import (
	"fmt"
	"path/filepath"
)

// ApkoBuilder represents a builder for APKO commands
type ApkoBuilder struct {
	configFile             string
	outputImage            string
	outputTarball          string
	keyringPaths           []string
	architectures          []string
	cacheDir               string
	extraArgs              []string
	wolfiKeyring           bool
	alpineKeyring          bool
	buildArch              string
	buildContext           string
	debug                  bool
	keyringAppendPlaintext []string
	noNetwork              bool
	repositoryAppend       []string
	timestamp              string
}

// NewApkoBuilder creates a new ApkoBuilder with default settings.
// It initializes the ApkoBuilder with default architectures "x86_64" and "aarch64".
func NewApkoBuilder() *ApkoBuilder {
	return &ApkoBuilder{
		architectures: []string{"x86_64", "aarch64"}, // Default architectures
	}
}

// WithConfigFile sets the configuration file for the APKO build.
// It takes a string parameter 'configFile' which is the path to the configuration file.
// It returns the updated ApkoBuilder instance.
func (b *ApkoBuilder) WithConfigFile(configFile string) *ApkoBuilder {
	b.configFile = configFile
	return b
}

// WithOutputImage sets the output image name for the APKO build.
// It takes a string parameter 'outputImage' which is the name of the output image.
// It returns the updated ApkoBuilder instance.
func (b *ApkoBuilder) WithOutputImage(outputImage string) *ApkoBuilder {
	b.outputImage = outputImage
	return b
}

// WithOutputTarball sets the output tarball path for the APKO build.
// It takes a string parameter 'outputTarball' which is the path to the output tarball.
// It returns the updated ApkoBuilder instance.
func (b *ApkoBuilder) WithOutputTarball(outputTarball string) *ApkoBuilder {
	b.outputTarball = outputTarball
	return b
}

// WithKeyring adds a keyring path to the APKO build.
// It takes a string parameter 'keyringPath' which is the path to the keyring file.
// It returns the updated ApkoBuilder instance.
func (b *ApkoBuilder) WithKeyring(keyringPath string) *ApkoBuilder {
	b.keyringPaths = append(b.keyringPaths, keyringPath)
	return b
}

// WithWolfiKeyring adds the Wolfi keyring to the APKO build.
// It sets the wolfiKeyring field to true.
// It returns the updated ApkoBuilder instance.
func (b *ApkoBuilder) WithWolfiKeyring() *ApkoBuilder {
	b.wolfiKeyring = true
	return b
}

// WithAlpineKeyring adds the Alpine keyring to the APKO build.
// It sets the alpineKeyring field to true.
// It returns the updated ApkoBuilder instance.
func (b *ApkoBuilder) WithAlpineKeyring() *ApkoBuilder {
	b.alpineKeyring = true
	return b
}

// WithArchitecture adds an architecture to the APKO build.
// It takes a string parameter 'arch' which is the architecture to be added.
// It returns the updated ApkoBuilder instance.
func (b *ApkoBuilder) WithArchitecture(arch string) *ApkoBuilder {
	b.architectures = append(b.architectures, arch)
	return b
}

// WithCacheDir sets the cache directory for the APKO build.
// It takes a string parameter 'cacheDir' which is the path to the cache directory.
// It returns the updated ApkoBuilder instance.
func (b *ApkoBuilder) WithCacheDir(cacheDir string) *ApkoBuilder {
	b.cacheDir = cacheDir
	return b
}

// WithExtraArg adds an extra argument to the APKO build command.
// It takes a string parameter 'arg' which is the extra argument to be added.
// It returns the updated ApkoBuilder instance.
func (b *ApkoBuilder) WithExtraArg(arg string) *ApkoBuilder {
	b.extraArgs = append(b.extraArgs, arg)
	return b
}

// WithBuildArch sets the build architecture
func (b *ApkoBuilder) WithBuildArch(arch string) *ApkoBuilder {
	b.buildArch = arch
	return b
}

// WithBuildContext sets the build context directory
func (b *ApkoBuilder) WithBuildContext(dir string) *ApkoBuilder {
	b.buildContext = dir
	return b
}

// WithDebug enables debug output
func (b *ApkoBuilder) WithDebug() *ApkoBuilder {
	b.debug = true
	return b
}

// WithKeyringAppendPlaintext appends a plaintext keyring
func (b *ApkoBuilder) WithKeyringAppendPlaintext(keyring string) *ApkoBuilder {
	b.keyringAppendPlaintext = append(b.keyringAppendPlaintext, keyring)
	return b
}

// WithNoNetwork disables network access during the build
func (b *ApkoBuilder) WithNoNetwork() *ApkoBuilder {
	b.noNetwork = true
	return b
}

// WithRepositoryAppend appends a repository to use for the build
func (b *ApkoBuilder) WithRepositoryAppend(repo string) *ApkoBuilder {
	b.repositoryAppend = append(b.repositoryAppend, repo)
	return b
}

// WithTimestamp sets the timestamp for the build
func (b *ApkoBuilder) WithTimestamp(timestamp string) *ApkoBuilder {
	b.timestamp = timestamp
	return b
}

// BuildCommand generates the APKO build command based on the current configuration of the ApkoBuilder.
// It returns a slice of strings representing the command and an error if any required fields are missing.
func (b *ApkoBuilder) BuildCommand() ([]string, error) {
	if b.configFile == "" {
		return nil, fmt.Errorf("config file is required")
	}

	if b.outputImage == "" {
		return nil, fmt.Errorf("output image name is required")
	}

	cmd := []string{"apko", "build"}

	for _, keyring := range b.keyringPaths {
		cmd = append(cmd, "--keyring-append", keyring)
	}

	if b.wolfiKeyring {
		cmd = append(cmd, "--keyring-append", "/etc/apk/keys/wolfi-signing.rsa.pub")
	}

	if b.alpineKeyring {
		cmd = append(cmd, "--keyring-append", "/etc/apk/keys/alpine-devel@lists.alpinelinux.org-4a6a0840.rsa.pub")
	}

	for _, arch := range b.architectures {
		cmd = append(cmd, "--arch", arch)
	}

	if b.cacheDir != "" {
		cmd = append(cmd, "--cache-dir", b.cacheDir)
	}

	if b.buildArch != "" {
		cmd = append(cmd, "--build-arch", b.buildArch)
	}

	if b.buildContext != "" {
		cmd = append(cmd, "--build-context", b.buildContext)
	}

	if b.debug {
		cmd = append(cmd, "--debug")
	}

	for _, keyring := range b.keyringAppendPlaintext {
		cmd = append(cmd, "--keyring-append-plaintext", keyring)
	}

	if b.noNetwork {
		cmd = append(cmd, "--no-network")
	}

	for _, repo := range b.repositoryAppend {
		cmd = append(cmd, "--repository-append", repo)
	}

	if b.timestamp != "" {
		cmd = append(cmd, "--timestamp", b.timestamp)
	}

	cmd = append(cmd, b.configFile, b.outputImage)

	if b.outputTarball != "" {
		cmd = append(cmd, b.outputTarball)
	}

	cmd = append(cmd, b.extraArgs...)

	return cmd, nil
}

// GetKeyringInfoForPreset returns the keyring information based on the preset.
// It takes a string parameter 'preset' which specifies the keyring preset ("alpine" or "wolfi").
// It returns a KeyringInfo struct and an error if the preset is unsupported.
func GetKeyringInfoForPreset(preset string) (KeyringInfo, error) {
	switch preset {
	case "alpine":
		return KeyringInfo{
			KeyURL:  "https://alpinelinux.org/keys/alpine-devel@lists.alpinelinux.org-4a6a0840.rsa.pub",
			KeyPath: "/etc/apk/keys/alpine-devel@lists.alpinelinux.org-4a6a0840.rsa.pub",
		}, nil
	case "wolfi":
		return KeyringInfo{
			KeyURL:  "https://packages.wolfi.dev/os/wolfi-signing.rsa.pub",
			KeyPath: "/etc/apk/keys/wolfi-signing.rsa.pub",
		}, nil
	default:
		return KeyringInfo{}, fmt.Errorf("unsupported preset: %s", preset)
	}
}

// GetCacheDir returns the APKO cache directory path.
// It takes a string parameter 'mntPrefix' which is the mount prefix.
// It returns the full path to the cache directory.
func GetCacheDir(mntPrefix string) string {
	return filepath.Join(mntPrefix, "var", "cache", "apko")
}

// GetOutputTarPath returns the APKO output tar file path.
// It takes a string parameter 'mntPrefix' which is the mount prefix.
// It returns the full path to the output tar file.
func GetOutputTarPath(mntPrefix string) string {
	return filepath.Join(mntPrefix, "image.tar")
}

// WithKeyRingWolfi adds the Wolfi keyring to the APKO build.
// It appends the Wolfi signing key to the keyringPaths.
// Returns the updated ApkoBuilder instance.
func (b *ApkoBuilder) WithKeyRingWolfi() *ApkoBuilder {
	wolfiKeyInfo, err := GetKeyringInfoForPreset("wolfi")
	if err == nil {
		b.keyringPaths = append(b.keyringPaths, wolfiKeyInfo.KeyPath)
	}
	return b
}

// WithKeyRingAlpine adds the Alpine keyring to the APKO build.
// It appends the Alpine signing key to the keyringPaths.
// Returns the updated ApkoBuilder instance.
func (b *ApkoBuilder) WithKeyRingAlpine() *ApkoBuilder {
	alpineKeyInfo, err := GetKeyringInfoForPreset("alpine")
	if err == nil {
		b.keyringPaths = append(b.keyringPaths, alpineKeyInfo.KeyPath)
	}
	return b
}

// KeyringInfo holds information about a keyring
type KeyringInfo struct {
	KeyURL  string
	KeyPath string
}
