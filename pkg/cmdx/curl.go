package cmdx

import (
	"fmt"
	"time"
)

// BuildCurlCommand generates a curl command string based on the provided parameters.
//
// Parameters:
//   - baseURL: The target URL for the curl command.
//   - headers: A map of HTTP headers to be included in the request.
//   - timeout: The maximum time allowed for the curl command to complete.
//   - authType: The type of authentication to use ("basic" or "bearer").
//   - authCredentials: The credentials for authentication (username:password for basic, token for bearer).
//
// Returns:
//   - A string containing the complete curl command.
func BuildCurlCommand(
	baseURL string,
	headers map[string]string,
	timeout time.Duration,
	authType string,
	authCredentials string,
) string {
	curlCmd := fmt.Sprintf("curl -m %d", int(timeout.Seconds()))

	for key, value := range headers {
		curlCmd += fmt.Sprintf(" -H '%s: %s'", key, value)
	}

	switch authType {
	case "basic":
		curlCmd += fmt.Sprintf(" -u '%s'", authCredentials)
	case "bearer":
		curlCmd += fmt.Sprintf(" -H 'Authorization: Bearer %s'", authCredentials)
	}

	curlCmd += fmt.Sprintf(" '%s'", baseURL)

	return curlCmd
}
