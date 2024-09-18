package builderx

import (
	"fmt"
	"path/filepath"

	"github.com/Excoriate/daggerx/pkg/fixtures"
)

// Architecture represents supported CPU architectures for APKO builds
type Architecture string

const (
	// ArchX8664 represents the x86_64 architecture
	ArchX8664 Architecture = "x86_64"
	// ArchAarch64 represents the aarch64 architecture
	ArchAarch64 Architecture = "aarch64"
	// ArchArmv7 represents the armv7 architecture
	ArchArmv7 Architecture = "armv7"
	// ArchPpc64le represents the ppc64le architecture
	ArchPpc64le Architecture = "ppc64le"
	// ArchS390x represents the s390x architecture
	ArchS390x Architecture = "s390x"
	// ApkoDefaultRepositoryURL is the default repository URL for APKO builds
	ApkoDefaultRepositoryURL = "cgr.dev/chainguard/apko"
)

// ApkoBuilder represents a builder for APKO (Alpine Package Keeper for OCI) images.
// It encapsulates all the configuration options and settings needed to build an APKO image.
type ApkoBuilder struct {
	// configFile is the path to the APKO configuration file.
	configFile string

	// outputImage is the name of the output OCI image.
	outputImage string

	// outputTarball is the path where the output tarball will be saved.
	outputTarball string

	// keyringPaths is a slice of paths to keyring files used for package verification.
	keyringPaths []string

	// architectures is a slice of target architectures for the build.
	architectures []string

	// cacheDir is the directory used for caching build artifacts.
	cacheDir string

	// extraArgs is a slice of additional arguments to pass to the APKO build command.
	extraArgs []string

	// wolfiKeyring indicates whether to use the Wolfi keyring.
	wolfiKeyring bool

	// alpineKeyring indicates whether to use the Alpine keyring.
	alpineKeyring bool

	// buildArch specifies the architecture to build for.
	buildArch string

	// buildContext is the build context directory.
	buildContext string

	// debug enables debug mode for verbose output.
	debug bool

	// keyringAppendPlaintext is a slice of plaintext keys to append to the keyring.
	keyringAppendPlaintext []string

	// noNetwork disables network access during the build.
	noNetwork bool

	// repositoryAppend is a slice of additional repositories to append.
	repositoryAppend []string

	// timestamp sets a specific timestamp for reproducible builds.
	timestamp string

	// tags is a slice of additional tags for the output image.
	tags []string
}

// WithBuildArch sets the build architecture for the APKO build.
// It takes an Architecture parameter 'arch' which is the desired build architecture.
// It returns the updated ApkoBuilder instance.
func (b *ApkoBuilder) WithBuildArch(arch Architecture) *ApkoBuilder {
	b.buildArch = string(arch)
	return b
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

// WithTag adds a tag to the APKO build.
// If no tag is provided, it defaults to "latest".
// It returns the updated ApkoBuilder instance.
func (b *ApkoBuilder) WithTag(tag ...string) *ApkoBuilder {
	if len(tag) > 0 {
		b.tags = append(b.tags, tag[0])
	} else {
		b.tags = append(b.tags, "latest")
	}
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

	if len(b.tags) > 0 {
		cmd = append(cmd, "--tag", b.tags[0])
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
	if mntPrefix == "" {
		mntPrefix = fixtures.MntPrefix
	}

	return filepath.Join(mntPrefix, "var", "cache", "apko")
}

// GetApkoConfigOrPreset returns the configuration file path if it is valid.
// It takes two string parameters: 'mntPrefix' which is the mount prefix, and 'cfgFile' which is the configuration file path.
// If 'mntPrefix' is empty, it defaults to fixtures.MntPrefix.
// If 'cfgFile' is empty, it returns an error indicating that the config file is required.
// If 'cfgFile' does not have an extension, it returns an error indicating that the config file must have an extension.
// If 'cfgFile' does not have a .yaml or .yml extension, it returns an error indicating that the config file must have a .yaml or .yml extension.
// It returns the configuration file path if all checks pass, otherwise it returns an error.
func GetApkoConfigOrPreset(mntPrefix, cfgFile string) (string, error) {
	if mntPrefix == "" {
		mntPrefix = fixtures.MntPrefix
	}

	if cfgFile == "" {
		return "", fmt.Errorf("config file is required")
	}

	ext := filepath.Ext(cfgFile)
	if ext == "" {
		return "", fmt.Errorf("config file must have an extension")
	}

	// Check if the file extension is .yaml or .yml
	if ext != ".yaml" && ext != ".yml" {
		return "", fmt.Errorf("config file must have a .yaml or .yml extension")
	}

	return cfgFile, nil
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
