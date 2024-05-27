package cmdbuilder

import (
	"reflect"
	"testing"

	"github.com/Excoriate/daggerx/pkg/types"
)

func TestGenerateCommand(t *testing.T) {
	tests := []struct {
		name     string
		command  string
		args     []string
		expected types.DaggerCMD
		err      bool
	}{
		// Edge case: empty command
		{
			name:     "Empty command",
			command:  "",
			args:     []string{"plan"},
			expected: nil,
			err:      true,
		},
		// Edge case: empty argument
		{
			name:     "Empty argument",
			command:  "terraform",
			args:     []string{""},
			expected: types.DaggerCMD{"terraform"},
			err:      false,
		},
		// Edge case: single command
		{
			name:     "Single command",
			command:  "terraform",
			args:     []string{"plan"},
			expected: types.DaggerCMD{"terraform", "plan"},
			err:      false,
		},
		// Edge case: command with spaces in argument
		{
			name:     "Command with spaces in argument",
			command:  "terraform",
			args:     []string{"apply --auto-approve"},
			expected: types.DaggerCMD{"terraform", "apply", "--auto-approve"},
			err:      false,
		},
		// Edge case: arguments with spaces and quotes
		{
			name:     "Arguments with spaces and quotes",
			command:  "terraform",
			args:     []string{"'apply --auto-approve'"},
			expected: types.DaggerCMD{"terraform", "'apply --auto-approve'"},
			err:      false,
		},
		// Edge case: arguments with double quotes
		{
			name:     "Arguments with double quotes",
			command:  "terraform",
			args:     []string{"\"apply --auto-approve\""},
			expected: types.DaggerCMD{"terraform", "\"apply --auto-approve\""},
			err:      false,
		},
		// Realistic: Terraform plan with variables
		{
			name:     "Terraform plan with variables",
			command:  "terraform",
			args:     []string{"plan", "-var", "foo=bar"},
			expected: types.DaggerCMD{"terraform", "plan", "-var", "foo=bar"},
			err:      false,
		},
		// Realistic: Terragrunt apply-all
		{
			name:     "Terragrunt apply-all",
			command:  "terragrunt",
			args:     []string{"apply-all"},
			expected: types.DaggerCMD{"terragrunt", "apply-all"},
			err:      false,
		},
		// Realistic: Node script with arguments
		{
			name:     "Node script with arguments",
			command:  "node",
			args:     []string{"script.js", "--env", "production"},
			expected: types.DaggerCMD{"node", "script.js", "--env", "production"},
			err:      false,
		},
		// Realistic: Go run with flags
		{
			name:     "Go run with flags",
			command:  "go",
			args:     []string{"run", "main.go", "--verbose"},
			expected: types.DaggerCMD{"go", "run", "main.go", "--verbose"},
			err:      false,
		},
		// Edge case: mixed empty and non-empty arguments
		{
			name:     "Mixed empty and non-empty arguments",
			command:  "terraform",
			args:     []string{"", "plan", "", "-var", "foo=bar", ""},
			expected: types.DaggerCMD{"terraform", "plan", "-var", "foo=bar"},
			err:      false,
		},
		// Edge case: argument with multiple spaces
		{
			name:     "Argument with multiple spaces",
			command:  "echo",
			args:     []string{"hello    world"},
			expected: types.DaggerCMD{"echo", "hello", "   world"},
			err:      false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := GenerateCommand(test.command, test.args...)
			if (err != nil) != test.err {
				t.Errorf("GenerateCommand(%q, %v) returned error %v, expected error %v", test.command, test.args, err, test.err)
			}
			if err == nil && !reflect.DeepEqual(*result, test.expected) {
				t.Errorf("GenerateCommand(%q, %v) = %v, expected %v", test.command, test.args, *result, test.expected)
			}
		})
	}
}
