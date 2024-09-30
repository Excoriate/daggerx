// Package cmdx provides utilities for generating and managing command-line
// commands in a structured and efficient manner. This package is designed
// to simplify the creation and handling of commands, ensuring that arguments
// with spaces are properly managed and that the commands are constructed
// correctly.
//
// The primary function in this package is GenerateCommand, which takes a main
// command and a variadic slice of arguments, and returns a pointer to a
// DaggerCMD slice. This slice includes the main command followed by the
// provided arguments. The function also performs validation to ensure that
// the main command is not empty.
//
// Example usage:
//
//	// Generate a Terraform plan command
//	cmd, err := GenerateCommand("terraform", "plan", "-var", "foo=bar", "apply --auto-approve")
//	if err != nil {
//	    // handle error
//	}
//	// Use cmd, e.g., fmt.Println(*cmd) // Output: [terraform plan -var foo=bar apply --auto-approve]
//
//	// Generate a Go run command
//	cmd, err = GenerateCommand("go", "run", "main.go", "--verbose")
//	if err != nil {
//	    // handle error
//	}
//	// Use cmd, e.g., fmt.Println(*cmd) // Output: [go run main.go --verbose]
//
// This package is intended for developers who need to programmatically
// generate and manage command-line commands in their Go applications.
package cmdx

import (
	"errors"
	"fmt"
	"strings"

	"github.com/Excoriate/daggerx/pkg/types"
)

// GenerateCommand generates a command with the provided arguments.
// It ensures that arguments with spaces are handled correctly.
//
// Parameters:
//   - command: A string representing the main command to be executed.
//   - args: A variadic slice of strings representing the arguments for the command.
//
// Returns:
//   - A pointer to a DaggerCMD slice, which includes the main command followed by the provided arguments.
//   - An error if the main command is empty.
//
// Example:
//
//	// Generate a Terraform plan command
//	cmd, err := GenerateCommand("terraform", "plan", "-var", "foo=bar", "apply --auto-approve")
//	if err != nil {
//	    // handle error
//	}
//	// Use cmd, e.g., fmt.Println(*cmd) // Output: [terraform plan -var foo=bar apply --auto-approve]
//
//	// Generate a Go run command
//	cmd, err = GenerateCommand("go", "run", "main.go", "--verbose")
//	if err != nil {
//	    // handle error
//	}
//	// Use cmd, e.g., fmt.Println(*cmd) // Output: [go run main.go --verbose]
func GenerateCommand(command string, args ...string) (*types.DaggerCMD, error) {
	// Validate the command
	if command == "" {
		return nil, fmt.Errorf("command cannot be empty")
	}

	// Initialize the command slice with the main command
	cmdWithArgs := types.DaggerCMD{command}

	for _, arg := range args {
		if arg == "" {
			continue
		}
		// Handle arguments with spaces
		if strings.Contains(arg, " ") && !strings.HasPrefix(arg, "'") && !strings.HasPrefix(arg, "\"") {
			parts := strings.SplitN(arg, " ", 2)
			cmdWithArgs = append(cmdWithArgs, parts[0], parts[1])
		} else {
			cmdWithArgs = append(cmdWithArgs, arg)
		}
	}

	return &cmdWithArgs, nil
}

// ConvertCMDToString converts a DaggerCMD slice to a string.
// It preserves the consistency of the original command, handling arguments with spaces appropriately.
//
// Parameters:
//   - cmd: A pointer to a DaggerCMD slice containing the command and its arguments.
//
// Returns:
//   - A string representation of the command, suitable for direct execution.
//
// Example:
//
//	cmd := types.DaggerCMD{"go", "run", "main.go", "--verbose"}
//	cmdString := ConvertCMDToString(&cmd)
//	fmt.Println(cmdString) // Output: go run main.go --verbose
//
//	cmd = types.DaggerCMD{"terraform", "plan", "-var", "foo=bar", "apply --auto-approve"}
//	cmdString = ConvertCMDToString(&cmd)
//	fmt.Println(cmdString) // Output: terraform plan -var foo=bar apply --auto-approve
func ConvertCMDToString(cmd *types.DaggerCMD) string {
	return strings.Join(*cmd, " ")
}

// GenerateShCommand generates a command wrapped for execution using `sh -c`.
// It ensures that arguments with spaces are correctly handled.
//
// Parameters:
//   - command: A string representing the main command to be executed.
//   - args: A variadic slice of strings representing the arguments for the command.
//
// Returns:
//   - A string containing the complete command wrapped for `sh -c` execution.
//   - An error if the main command is empty.
//
// Example:
//
//	cmd, err := GenerateShCommand("echo", "Hello, World!")
//	if err != nil {
//	    // handle error
//	}
//	fmt.Println(cmd) // Output: sh -c "echo Hello, World!"
func GenerateShCommand(command string, args ...string) (string, error) {
	// Validate the command
	if command == "" {
		return "", fmt.Errorf("command cannot be empty")
	}

	// Initialize the command slice with the main command
	cmdWithArgs := types.DaggerCMD{command}

	for _, arg := range args {
		if arg == "" {
			continue
		}
		cmdWithArgs = append(cmdWithArgs, arg)
	}

	cmdString := ConvertCMDToString(&cmdWithArgs)
	//nolint: gocritic // It's okay to use fmt.Sprintf here
	return fmt.Sprintf("sh -c \"%s\"", cmdString), nil
}

// GenerateSHCommandAsDaggerCMD generates a command wrapped for execution using `sh -c` and returns a DaggerCMD.
//
// Parameters:
//   - command: A string representing the main command to be executed.
//   - args: A variadic slice of strings representing the arguments for the command.
//
// Returns:
//   - A pointer to a DaggerCMD slice containing the complete command wrapped for `sh -c` execution.
//   - An error if the main command is empty.
//
// Example:
//
//	cmd, err := GenerateSHCommandAsDaggerCMD("echo", "Hello, World!")
//	if err != nil {
//	    // handle error
//	}
//	fmt.Println(cmd) // Output: [sh -c "echo Hello, World!"]
func GenerateSHCommandAsDaggerCMD(command string, args ...string) (*types.DaggerCMD, error) {
	cmd, err := GenerateShCommand(command, args...)
	if err != nil {
		return nil, err
	}

	return &types.DaggerCMD{cmd}, nil
}

// GenerateDaggerCMDFromStr generates a DaggerCMD from a plain command string.
// It splits the command string into a slice of strings, ensuring it handles
// special characters and consecutive spaces correctly.
//
// Parameters:
//   - commands: A string representing the command.
//
// Returns:
//   - A slice of strings representing the command and its arguments.
//   - An error if the command string is empty or invalid.
func GenerateDaggerCMDFromStr(commands string) ([]string, error) {
	if strings.TrimSpace(commands) == "" {
		return nil, errors.New("command string cannot be empty")
	}

	// Split command string by spaces
	cmdSlice := strings.Fields(commands)
	if len(cmdSlice) == 0 {
		return nil, errors.New("invalid command string")
	}

	return cmdSlice, nil
}

// PtToSlice converts a *types.DaggerCMD to a []string.
//
// Parameters:
//   - cmd: A pointer to a DaggerCMD slice containing the command and its arguments.
//
// Returns:
//   - A slice of strings representing the command and its arguments.
//   - An error if the input DaggerCMD is nil.
func PtToSlice(cmd *types.DaggerCMD) ([]string, error) {
	if cmd == nil {
		return nil, fmt.Errorf("input DaggerCMD cannot be nil")
	}
	return *cmd, nil
}
