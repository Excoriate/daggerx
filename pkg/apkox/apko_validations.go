package apkox

import (
	"fmt"
	"net/url"
	"strings"
)

// IsKeyringFormatValid validates the format of the provided keyrings.
// Each keyring should be in one of two valid formats:
// 1. "path=url" where:
//   - path is a relative path under "/etc/apk/keys/"
//   - url is a valid URL
//
// 2. "url" where:
//   - url is a valid URL
//
// The function also takes an optional parameter enforceHTTPS which defaults to true.
// If enforceHTTPS is true, the URL must start with "https://".
//
// Returns an error if any keyring is invalid, nil otherwise.
func IsKeyringFormatValid(keyrings []string, enforceHTTPS ...bool) error {
	httpsRequired := true
	if len(enforceHTTPS) > 0 {
		httpsRequired = enforceHTTPS[0]
	}

	for _, keyring := range keyrings {
		parts := strings.Split(keyring, "=")
		var urlStr string

		switch len(parts) {
		case 2:
			path := parts[0]
			if !strings.HasPrefix(path, "/etc/apk/keys/") {
				return fmt.Errorf("invalid keyring path: %s", path)
			}
			urlStr = parts[1]
		case 1:
			urlStr = parts[0]
		default:
			return fmt.Errorf("invalid keyring format: %s", keyring)
		}

		if err := validateURL(urlStr, httpsRequired); err != nil {
			return fmt.Errorf("invalid keyring URL: %s, error: %w", urlStr, err)
		}
	}
	return nil
}

// validateURL checks if the given string is a valid URL and optionally enforces HTTPS.
func validateURL(urlStr string, enforceHTTPS bool) error {
	parsedURL, err := url.Parse(urlStr)
	if err != nil {
		return err
	}
	if parsedURL.Scheme == "" {
		return fmt.Errorf("missing URL scheme")
	}
	if enforceHTTPS && parsedURL.Scheme != "https" {
		return fmt.Errorf("HTTPS is required")
	}
	return nil
}
