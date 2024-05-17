package envvars

import (
	"errors"
	"fmt"
	"github.com/Excoriate/daggerx/pkg/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestToDaggerEnvVarsFromStr(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []types.DaggerEnvVars
		err      error
	}{
		{
			name:  "Valid key=value pairs",
			input: "FOO=bar,BAZ=qux",
			expected: []types.DaggerEnvVars{
				{Name: "FOO", Value: "bar"},
				{Name: "BAZ", Value: "qux"},
			},
			err: nil,
		},
		{
			name:     "Empty input string",
			input:    "",
			expected: nil,
			err:      errors.New("input string is empty"),
		},
		{
			name:     "Invalid key=value format",
			input:    "FOO=bar,INVALID",
			expected: nil,
			err:      fmt.Errorf("invalid environment variable format: INVALID"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := ToDaggerEnvVarsFromStr(test.input)
			if test.err != nil {
				assert.EqualError(t, err, test.err.Error())
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, test.expected, result)
		})
	}
}

func TestToEnvVarsDaggerFromMap(t *testing.T) {
	tests := []struct {
		input    map[string]string
		expected []types.DaggerEnvVars
		err      bool
	}{
		{map[string]string{"key1": "value1", "key2": "value2"}, []types.DaggerEnvVars{
			{Name: "key1", Value: "value1"},
			{Name: "key2", Value: "value2"},
		}, false},
		{map[string]string{}, nil, true},
		{map[string]string{"key1": "value1", "": "value2"}, nil, true},
	}

	for _, test := range tests {
		result, err := ToDaggerEnvVarsFromMap(test.input)
		if (err != nil) != test.err {
			t.Errorf("toEnvVarsDaggerFromMap(%v) returned error %v, expected error: %v", test.input, err, test.err)
		}
		if !test.err && !equalDaggerEnvVars(result, test.expected) {
			t.Errorf("toEnvVarsDaggerFromMap(%v) = %v, expected %v", test.input, result, test.expected)
		}
	}
}

func TestToEnvVarsDaggerFromSlice(t *testing.T) {
	tests := []struct {
		input    []string
		expected []types.DaggerEnvVars
		err      bool
	}{
		{[]string{"key1=value1", "key2=value2"}, []types.DaggerEnvVars{
			{Name: "key1", Value: "value1"},
			{Name: "key2", Value: "value2"},
		}, false},
		{[]string{}, nil, true},
		{[]string{"key1=value1", "key2value2"}, nil, true},
		{[]string{"key1=value1", "=value2"}, nil, true},
		{[]string{"key1=value1", ""}, []types.DaggerEnvVars{
			{Name: "key1", Value: "value1"},
		}, false},
	}

	for _, test := range tests {
		result, err := ToDaggerEnvVarsFromSlice(test.input)
		if (err != nil) != test.err {
			t.Errorf("ToDaggerEnvVarsFromSlice(%v) returned error %v, expected error: %v", test.input, err, test.err)
		}
		if !test.err && !equalDaggerEnvVars(result, test.expected) {
			t.Errorf("ToDaggerEnvVarsFromSlice(%v) = %v, expected %v", test.input, result, test.expected)
		}
	}
}

func equal(a, b map[string]string) bool {
	if len(a) != len(b) {
		return false
	}
	for k, v := range a {
		if b[k] != v {
			return false
		}
	}
	return true
}

func equalDaggerEnvVars(a, b []types.DaggerEnvVars) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if b[i] != v {
			return false
		}
	}
	return true
}
