package execmd

import (
	"strings"

	"github.com/Excoriate/daggerx/pkg/types"
)

// BuildArgs processes and builds a command argument list from the provided arguments.
// It trims spaces from each argument and splits arguments with spaces into separate arguments.
//
// Parameters:
//   - args: A variadic parameter that takes multiple strings representing command arguments.
//
// Returns:
//   - A DaggerCMD slice containing the processed arguments.
//
// Example:
//
//	// Build arguments for a Terraform command
//	args := BuildArgs("plan", "-var 'foo=bar'", "apply --auto-approve")
//	// args now contains: ["plan", "-var", "'foo=bar'", "apply", "--auto-approve"]
//
//	// Build arguments for a Go command
//	args = BuildArgs("run", "main.go", "--verbose")
//	// args now contains: ["run", "main.go", "--verbose"]
func BuildArgs(args ...string) types.DaggerCMD {
	var merged []string
	for _, arg := range args {
		if arg = strings.TrimSpace(arg); arg != "" {
			parts := strings.Fields(arg) // Splits the string into substrings, removing any space characters, including newlines.
			merged = append(merged, parts...)
		}
	}
	return merged
}
