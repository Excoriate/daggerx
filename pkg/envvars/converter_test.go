package envvars

import (
	"github.com/Excoriate/daggerx/pkg/types"
	"testing"
)

func TestToEnvVarsFromStr(t *testing.T) {
	tests := []struct {
		input    string
		expected map[string]string
		err      bool
	}{
		{"key1=value1,key2=value2,key3=value3", map[string]string{"key1": "value1", "key2": "value2", "key3": "value3"}, false},
		{"key1=value1,key2=value2,key3=", map[string]string{"key1": "value1", "key2": "value2", "key3": ""}, false},
		{"key1=value1,key2=value2,key3", nil, true},
		{"", nil, true},
		{"key1=,key2=value2,key3=value3", map[string]string{"key1": "", "key2": "value2", "key3": "value3"}, false},
		{",key1=value1,key2=value2", map[string]string{"key1": "value1", "key2": "value2"}, false},
	}

	for _, test := range tests {
		result, err := ToDaggerEnvVarsFromStr(test.input)
		if (err != nil) != test.err {
			t.Errorf("ToDaggerEnvVarsFromStr(%q) returned error %v, expected error: %v", test.input, err, test.err)
		}
		if !test.err && !equal(result, test.expected) {
			t.Errorf("ToDaggerEnvVarsFromStr(%q) = %v, expected %v", test.input, result, test.expected)
		}
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
