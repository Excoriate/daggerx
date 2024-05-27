package cmdbuilder

import (
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
