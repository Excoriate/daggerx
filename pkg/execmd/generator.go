package execmd

import (
	"fmt"
	"strings"
)

// DaggerCMD is an alias for a slice of strings representing a command.
type DaggerCMD []string

// GenerateCommand generates a command with the provided arguments.
// It ensures that arguments with spaces are handled correctly.
func GenerateCommand(command string, args ...string) (*DaggerCMD, error) {
	// Validate the command
	if command == "" {
		return nil, fmt.Errorf("command cannot be empty")
	}

	// Initialize the command slice with the main command
	cmdWithArgs := DaggerCMD{command}

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
