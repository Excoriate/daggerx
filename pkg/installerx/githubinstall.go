package installerx

import (
	"fmt"
	"path/filepath"
	"strings"
)

// GitHubAssetParams represents parameters for downloading a GitHub release asset
type GitHubAssetParams struct {
	// Organization or user who owns the repository
	Owner string
	// Repository name
	Repo string
	// Version of the release (with or without 'v' prefix)
	Version string
	// Asset name pattern (e.g., "terraform-docs-v{version}-{os}-{arch}.tar.gz")
	// If empty, AssetName will be used directly
	AssetPattern string
	// AssetName is the direct name of the asset if pattern is not used
	AssetName string
	// InstallDir is the directory where the binary will be installed
	InstallDir string
	// BinaryName is the name of the binary to be installed after extraction
	BinaryName string
	// OS target operating system (defaults to "linux")
	OS string
	// Arch target architecture (defaults to "amd64")
	Arch string
	// ExtractPath specifies the path to the binary within the archive
	// If empty, BinaryName will be used
	ExtractPath string
}

// GetGitHubAssetInstallCommand generates a command to download and install a GitHub release asset
//
// Parameters:
// - params: GitHubAssetParams containing the required information
//
// Returns:
// - string: The installation command
// - error: Error if parameters are invalid
func GetGitHubAssetInstallCommand(params GitHubAssetParams) (string, error) {
	if params.Owner == "" {
		return "", fmt.Errorf("owner is required")
	}

	if params.Repo == "" {
		return "", fmt.Errorf("repo is required")
	}

	if params.Version == "" {
		return "", fmt.Errorf("version is required")
	}

	if params.AssetPattern == "" && params.AssetName == "" {
		return "", fmt.Errorf("either asset pattern or asset name is required")
	}

	if params.BinaryName == "" {
		return "", fmt.Errorf("binary name is required")
	}

	// Set defaults
	if params.InstallDir == "" {
		params.InstallDir = DefaultInstallDir
	}

	if params.OS == "" {
		params.OS = "linux"
	}

	if params.Arch == "" {
		params.Arch = "amd64"
	}

	if params.ExtractPath == "" {
		params.ExtractPath = params.BinaryName
	}

	// Ensure version has 'v' prefix
	version := params.Version
	if !strings.HasPrefix(version, "v") {
		version = "v" + version
	}

	// Determine the asset name
	var asset string
	if params.AssetPattern != "" {
		asset = strings.NewReplacer(
			"{version}", strings.TrimPrefix(version, "v"),
			"{os}", params.OS,
			"{arch}", params.Arch,
		).Replace(params.AssetPattern)
	} else {
		asset = params.AssetName
	}

	installPath := filepath.Join(params.InstallDir, params.BinaryName)
	isTarGz := strings.HasSuffix(strings.ToLower(asset), ".tar.gz") || strings.HasSuffix(strings.ToLower(asset), ".tgz")
	isZip := strings.HasSuffix(strings.ToLower(asset), ".zip")

	// Build the command based on the asset type
	var command string

	switch {
	case isTarGz:
		command = fmt.Sprintf(`set -ex
curl -fL https://github.com/%[1]s/%[2]s/releases/download/%[3]s/%[4]s -o /tmp/%[4]s
cd /tmp && tar -xzf %[4]s
mv /tmp/%[5]s %[6]s
chmod +x %[6]s
rm -f /tmp/%[4]s`, params.Owner, params.Repo, version, asset, params.ExtractPath, installPath)
	case isZip:
		command = fmt.Sprintf(`set -ex
curl -fL https://github.com/%[1]s/%[2]s/releases/download/%[3]s/%[4]s -o /tmp/%[4]s
cd /tmp && unzip -o %[4]s
mv /tmp/%[5]s %[6]s
chmod +x %[6]s
rm -f /tmp/%[4]s`, params.Owner, params.Repo, version, asset, params.ExtractPath, installPath)
	default:
		command = fmt.Sprintf(`set -ex
curl -fL https://github.com/%[1]s/%[2]s/releases/download/%[3]s/%[4]s -o %[5]s
chmod +x %[5]s`, params.Owner, params.Repo, version, asset, installPath)
	}

	return strings.TrimSpace(command), nil
}
