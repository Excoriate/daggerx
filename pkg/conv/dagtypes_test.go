package conv

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestStruct struct {
	Field1 string
	Field2 int
}

func TestToAnyType(t *testing.T) {
	tests := []struct {
		name     string
		input    interface{}
		expected interface{}
	}{
		{
			name:     "Convert int to int",
			input:    123,
			expected: 123,
		},
		{
			name:     "Convert string to string",
			input:    "test",
			expected: "test",
		},
		{
			name:     "Convert float64 to float64",
			input:    123.45,
			expected: 123.45,
		},
		{
			name:     "Convert nil to int (fail)",
			input:    nil,
			expected: nil,
		},
		{
			name: "Convert struct to struct",
			input: TestStruct{
				Field1: "value",
				Field2: 10,
			},
			expected: TestStruct{
				Field1: "value",
				Field2: 10,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			switch test.expected.(type) {
			case int:
				result := ToAnyType[int](test.input)
				if test.expected == nil {
					assert.Nil(t, result)
				} else {
					assert.NotNil(t, result)
					assert.Equal(t, test.expected.(int), *result)
				}
			case string:
				result := ToAnyType[string](test.input)
				if test.expected == nil {
					assert.Nil(t, result)
				} else {
					assert.NotNil(t, result)
					assert.Equal(t, test.expected.(string), *result)
				}
			case float64:
				result := ToAnyType[float64](test.input)
				if test.expected == nil {
					assert.Nil(t, result)
				} else {
					assert.NotNil(t, result)
					assert.Equal(t, test.expected.(float64), *result)
				}
			case TestStruct:
				result := ToAnyType[TestStruct](test.input)
				if test.expected == nil {
					assert.Nil(t, result)
				} else {
					assert.NotNil(t, result)
					assert.Equal(t, test.expected.(TestStruct), *result)
				}
			default:
				result := ToAnyType[interface{}](test.input)
				assert.Nil(t, result)
			}
		})
	}
}
