// Package cleaner provides utilities for cleaning and sanitizing strings.
// This package is particularly useful for preprocessing command arguments
// or any other strings that may contain unwanted characters.
//
// The cleaner package currently includes the following functions:
//
//   - RemoveCommas: Removes all commas from a given string.
//
// Example usage:
//
//	import (
//	    "fmt"
//	    "github.com/yourusername/cleaner"
//	)
//
//	func main() {
//	    cmd := "terragrunt run-all plan --terragrunt-non-interactive -compact-warnings, -no-color, -lock=false"
//	    cleanCmd := cleaner.RemoveCommas(cmd)
//	    fmt.Println(cleanCmd) // Output: "terragrunt run-all plan --terragrunt-non-interactive -compact-warnings -no-color -lock=false"
//	}
package cleaner

import "strings"

// RemoveCommas removes all commas from the provided string.
// This function is useful for cleaning command arguments that may contain invalid commas.
//
// Parameters:
//   - s: A string from which all commas should be removed.
//
// Returns:
//   - A string with all commas removed.
//
// Example:
//
//	cmd := "terragrunt run-all plan --terragrunt-non-interactive -compact-warnings, -no-color, -lock=false"
//	cleanCmd := RemoveCommas(cmd)
//	fmt.Println(cleanCmd) // Output: "terragrunt run-all plan --terragrunt-non-interactive -compact-warnings -no-color -lock=false"
func RemoveCommas(s string) string {
	return strings.ReplaceAll(s, ",", "")
}
