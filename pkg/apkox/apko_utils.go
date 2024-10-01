package apkox

import (
	"fmt"
	"strings"
)

// KeyringPars holds the parameters for a keyring.
// It includes the path to the keyring file and the URL from which the keyring can be downloaded.
type KeyringPars struct {
	// Path is the file system path to the keyring file.
	Path string
	// URL is the web address from which the keyring can be downloaded.
	URL string
}

// GetPathFromKeyring extracts the path from the keyring string.
// The keyring string is expected to be in the format "path=url".
// It returns the path as a string and an error if the format is invalid.
func (k *KeyringPars) GetPathFromKeyring(keyring string) (string, error) {
	if keyring == "" {
		return "", fmt.Errorf("keyring string is empty")
	}
	parts := strings.SplitN(keyring, "=", 2)
	if len(parts) != 2 {
		return "", fmt.Errorf("invalid keyring format: %s", keyring)
	}
	return parts[0], nil
}

// GetURLFromKeyring extracts the URL from the keyring string.
// The keyring string is expected to be in the format "path=url".
// It returns the URL as a string and an error if the format is invalid.
func (k *KeyringPars) GetURLFromKeyring(keyring string, enforceHTTPS bool) (string, error) {
	if keyring == "" {
		return "", fmt.Errorf("keyring string is empty")
	}
	parts := strings.SplitN(keyring, "=", 2)
	if len(parts) != 2 {
		return "", fmt.Errorf("invalid keyring format: %s", keyring)
	}
	url := parts[1]
	if enforceHTTPS && !strings.HasPrefix(url, "https://") {
		return "", fmt.Errorf("URL does not use HTTPS: %s", url)
	}
	return url, nil
}

// ParseKeyring splits a keyring string into its component parts.
// It returns a KeyringSkeleton struct containing the path and URL.
// If the keyring string only contains a URL, the Path field will be empty.
func ParseKeyring(keyring string) (KeyringSkeleton, error) {
	parts := strings.SplitN(keyring, "=", 2)
	switch len(parts) {
	case 2:
		path := parts[0]
		if !strings.HasPrefix(path, "/etc/apk/keys/") {
			return KeyringSkeleton{}, fmt.Errorf("invalid keyring path: %s", path)
		}
		return KeyringSkeleton{
			Path: path,
			URL:  parts[1],
		}, nil
	case 1:
		return KeyringSkeleton{
			URL: parts[0],
		}, nil
	default:
		return KeyringSkeleton{}, fmt.Errorf("invalid keyring format: %s", keyring)
	}
}

// ValidateKeyring validates a single keyring string.
// It reuses the logic from IsKeyringFormatValid in apko_validations.go.
func ValidateKeyring(keyring string, enforceHTTPS bool) error {
	skeleton, err := ParseKeyring(keyring)
	if err != nil {
		return err
	}
	return validateURL(skeleton.URL, enforceHTTPS)
}
