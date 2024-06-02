package cleaner

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveCommas(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			input:    "terragrunt run-all plan --terragrunt-non-interactive -compact-warnings, -no-color, -lock=false",
			expected: "terragrunt run-all plan --terragrunt-non-interactive -compact-warnings -no-color -lock=false",
		},
		{
			input:    "echo Hello, World!",
			expected: "echo Hello World!",
		},
		{
			input:    "command,arg1,arg2,arg3",
			expected: "commandarg1arg2arg3",
		},
		{
			input:    "no commas here",
			expected: "no commas here",
		},
		{
			input:    "",
			expected: "",
		},
	}

	for _, test := range tests {
		result := RemoveCommas(test.input)
		assert.Equalf(t, test.expected, result, "expected %s but got %s", test.expected, result)
	}
}
