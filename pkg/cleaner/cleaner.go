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
