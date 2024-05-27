package workdir

import (
	"testing"

	"github.com/Excoriate/daggerx/pkg/fixtures"
	"github.com/stretchr/testify/assert"
)

func TestGetDefault(t *testing.T) {
	defaultWorkdir := fixtures.MntPrefix

	result := GetDefault()
	assert.Equal(t, defaultWorkdir, result, "Expected %s but got %s", defaultWorkdir, result)
}

func TestSetOrDefault_ValidInput(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{{
		input:    "/path/to/workdir",
		expected: "/path/to/workdir",
	}, {
		input:    "path/to/workdir",
		expected: "/path/to/workdir",
	}}

	for _, test := range tests {
		result := SetOrDefault(test.input)
		assert.Equal(t, test.expected, result, "Expected %s but got %s", test.expected, result)
	}
}

func TestSetOrDefault_EmptyInput(t *testing.T) {
	defaultWorkdir := fixtures.MntPrefix

	result := SetOrDefault("")
	assert.Equal(t, defaultWorkdir, result, "Expected %s but got %s", defaultWorkdir, result)
}

func TestIsValid_ValidInput(t *testing.T) {
	tests := []struct {
		input string
	}{{
		input: fixtures.MntPrefix + "/path/to/workdir",
	}, {
		input: fixtures.MntPrefix,
	}}

	for _, test := range tests {
		err := IsValid(test.input)
		assert.NoError(t, err, "Expected no error for input %s", test.input)
	}
}

func TestIsValid_InvalidInput_NotAbsolutePath(t *testing.T) {
	tests := []struct {
		input string
	}{{
		input: "relative/path",
	}, {
		input: "./relative/path",
	}}

	for _, test := range tests {
		err := IsValid(test.input)
		assert.Error(t, err, "Expected error for input %s", test.input)
		assert.EqualError(t, err, "workdir must be an absolute path: "+test.input)
	}
}
