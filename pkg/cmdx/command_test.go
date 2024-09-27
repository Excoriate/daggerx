package cmdx

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"

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

func TestConvertCMDToString(t *testing.T) {
	tests := []struct {
		cmd      *types.DaggerCMD
		expected string
	}{
		{
			cmd:      &types.DaggerCMD{"go", "run", "main.go", "--verbose"},
			expected: "go run main.go --verbose",
		},
		{
			cmd:      &types.DaggerCMD{"echo", "Hello, World!"},
			expected: "echo Hello, World!",
		},
		{
			cmd:      &types.DaggerCMD{"terraform", "plan", "-var", "foo=bar", "apply", "--auto-approve"},
			expected: "terraform plan -var foo=bar apply --auto-approve",
		},
	}

	for _, test := range tests {
		result := ConvertCMDToString(test.cmd)
		assert.Equal(t, test.expected, result)
	}
}

func TestGenerateShCommand(t *testing.T) {
	tests := []struct {
		command  string
		args     []string
		expected string
		hasError bool
	}{
		{
			command:  "echo",
			args:     []string{"Hello, World!"},
			expected: "sh -c \"echo Hello, World!\"",
			hasError: false,
		},
		{
			command:  "go",
			args:     []string{"run", "main.go", "--verbose"},
			expected: "sh -c \"go run main.go --verbose\"",
			hasError: false,
		},
		{
			command:  "",
			args:     []string{"run", "main.go"},
			expected: "",
			hasError: true,
		},
	}

	for _, test := range tests {
		result, err := GenerateShCommand(test.command, test.args...)
		if test.hasError {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
			assert.Equal(t, test.expected, result)
		}
	}
}

func TestGenerateSHCommandAsDaggerCMD(t *testing.T) {
	tests := []struct {
		name        string
		command     string
		args        []string
		expected    *types.DaggerCMD
		expectError bool
	}{
		{
			name:        "Empty command",
			command:     "",
			args:        []string{},
			expected:    nil,
			expectError: true,
		},
		{
			name:        "Simple command without args",
			command:     "echo",
			args:        []string{},
			expected:    &types.DaggerCMD{"sh -c \"echo\""},
			expectError: false,
		},
		{
			name:        "Simple command with args",
			command:     "echo",
			args:        []string{"Hello, World!"},
			expected:    &types.DaggerCMD{"sh -c \"echo Hello, World!\""},
			expectError: false,
		},
		{
			name:        "Command with multiple args",
			command:     "terragrunt",
			args:        []string{"-auto-approve", "--var-file=asdadas"},
			expected:    &types.DaggerCMD{"sh -c \"terragrunt -auto-approve --var-file=asdadas\""},
			expectError: false,
		},
		{
			name:        "Command with special characters",
			command:     "echo",
			args:        []string{"Hello", "&&", "echo", "World!"},
			expected:    &types.DaggerCMD{"sh -c \"echo Hello && echo World!\""},
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd, err := GenerateSHCommandAsDaggerCMD(tt.command, tt.args...)
			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, cmd)
			}
		})
	}
}

func TestGenerateDaggerCMDFromStr(t *testing.T) {
	tests := []struct {
		name        string
		commands    string
		expected    []string
		expectError bool
	}{
		{
			name:        "Empty command string",
			commands:    "",
			expected:    nil,
			expectError: true,
		},
		{
			name:        "Blank spaces command string",
			commands:    "     ",
			expected:    nil,
			expectError: true,
		},
		{
			name:        "Simple command without args",
			commands:    "terraform plan",
			expected:    []string{"terraform", "plan"},
			expectError: false,
		},
		{
			name:        "Command with args",
			commands:    "terraform plan --var-file=asdasda -auto-approve",
			expected:    []string{"terraform", "plan", "--var-file=asdasda", "-auto-approve"},
			expectError: false,
		},
		{
			name:        "Complex command with pipes",
			commands:    "terraform show -json | jq",
			expected:    []string{"terraform", "show", "-json", "|", "jq"},
			expectError: false,
		},
		{
			name:        "Command with consecutive spaces",
			commands:    "  terraform   plan   --var-file=asdasda    -auto-approve  ",
			expected:    []string{"terraform", "plan", "--var-file=asdasda", "-auto-approve"},
			expectError: false,
		},
		{
			name:        "Command with special characters",
			commands:    "echo 'Hello, World!' && echo 'Done!'",
			expected:    []string{"echo", "'Hello,", "World!'", "&&", "echo", "'Done!'"},
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := GenerateDaggerCMDFromStr(tt.commands)
			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, result)
			}
		})
	}
}
