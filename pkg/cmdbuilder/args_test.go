package cmdbuilder

import (
	"reflect"
	"testing"

	"github.com/Excoriate/daggerx/pkg/types"
)

func TestBuildArgs(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected types.DaggerCMD
	}{
		// Edge case: single command with spaces
		{
			name:     "Single command with spaces",
			input:    []string{"   plan   "},
			expected: types.DaggerCMD{"plan"},
		},
		// Edge case: multiple spaces between arguments
		{
			name:     "Multiple spaces between arguments",
			input:    []string{"plan", "   -var 'foo=bar'", "  apply --auto-approve  "},
			expected: types.DaggerCMD{"plan", "-var", "'foo=bar'", "apply", "--auto-approve"},
		},
		// Realistic: Terraform plan with variables
		{
			name:     "Terraform plan with variables",
			input:    []string{"plan", "-var 'foo=bar'", "apply --auto-approve"},
			expected: types.DaggerCMD{"plan", "-var", "'foo=bar'", "apply", "--auto-approve"},
		},
		// Realistic: Go run with flags
		{
			name:     "Go run with flags",
			input:    []string{"run", "main.go", "--verbose"},
			expected: types.DaggerCMD{"run", "main.go", "--verbose"},
		},
		// Realistic: Node script with arguments
		{
			name:     "Node script with arguments",
			input:    []string{"script.js", "--env production"},
			expected: types.DaggerCMD{"script.js", "--env", "production"},
		},
		// Edge case: mixed empty and non-empty arguments
		{
			name:     "Mixed empty and non-empty arguments",
			input:    []string{"", "plan", "", "-var", "foo=bar", ""},
			expected: types.DaggerCMD{"plan", "-var", "foo=bar"},
		},
		// Edge case: argument with multiple spaces
		{
			name:     "Argument with multiple spaces",
			input:    []string{"hello    world"},
			expected: types.DaggerCMD{"hello", "world"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := BuildArgs(test.input...)
			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("BuildArgs(%v) = %v, expected %v", test.input, result, test.expected)
			}
		})
	}
}
