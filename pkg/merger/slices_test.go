package merger

import (
	"testing"
)

func TestMergeSlices(t *testing.T) {
	tests := []struct {
		name     string
		input    [][]string
		expected []string
	}{
		{
			name:     "Single slice",
			input:    [][]string{{"a", "b", "c"}},
			expected: []string{"a", "b", "c"},
		},
		{
			name:     "Multiple slices",
			input:    [][]string{{"a", "b"}, {"c", "d"}, {"e", "f"}},
			expected: []string{"a", "b", "c", "d", "e", "f"},
		},
		{
			name:     "Empty slices",
			input:    [][]string{{}, {}, {}},
			expected: []string{},
		},
		{
			name:     "Mix of empty and non-empty slices",
			input:    [][]string{{"a", "b"}, {}, {"c"}, {}},
			expected: []string{"a", "b", "c"},
		},
		{
			name:     "All slices empty",
			input:    [][]string{{}, {}},
			expected: []string{},
		},
		{
			name:     "No slices",
			input:    [][]string{},
			expected: []string{},
		},
		{
			name:     "Nested empty slice",
			input:    [][]string{nil},
			expected: []string{},
		},
		{
			name:     "Single element slices",
			input:    [][]string{{"a"}, {"b"}, {"c"}},
			expected: []string{"a", "b", "c"},
		},
		{
			name:     "Mixed length slices",
			input:    [][]string{{"a"}, {"b", "c"}, {"d", "e", "f"}},
			expected: []string{"a", "b", "c", "d", "e", "f"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := MergeSlices(test.input...)
			if len(result) != len(test.expected) || !equalSlices(result, test.expected) {
				t.Errorf("MergeSlices(%v) = %v, expected %v", test.input, result, test.expected)
			}
		})
	}
}

func equalSlices(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
