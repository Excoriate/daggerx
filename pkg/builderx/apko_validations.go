package builderx

import (
	"fmt"
	"strings"
)

// IsKeyringFormatValid validates the format of the provided keyrings.
// Each keyring should be in the format "filesystempath=URL", where:
// - filesystempath is a relative path under "/etc/apk/keys/"
// - URL is a valid URL
// The function also takes an optional parameter enforceHTTPS which defaults to true.
// If enforceHTTPS is true, the URL must start with "https://".
func IsKeyringFormatValid(keyrings []string, enforceHTTPS ...bool) error {
	httpsRequired := true
	if len(enforceHTTPS) > 0 {
		httpsRequired = enforceHTTPS[0]
	}

	for _, keyring := range keyrings {
		parts := strings.Split(keyring, "=")
		if len(parts) != 2 {
			return fmt.Errorf("invalid keyring format: %s", keyring)
		}
		path := parts[0]
		url := parts[1]
		if !strings.HasPrefix(path, "/etc/apk/keys/") {
			return fmt.Errorf("invalid keyring path: %s", path)
		}
		if httpsRequired && !strings.HasPrefix(url, "https://") {
			return fmt.Errorf("invalid keyring URL: %s", url)
		}
	}
	return nil
}
