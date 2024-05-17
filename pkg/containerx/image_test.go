package containerx

import (
	"fmt"
	"testing"

	"github.com/Excoriate/daggerx/pkg/fixtures"
	"github.com/stretchr/testify/assert"
)

func TestSetDefaultIfEmpty(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			input:    "",
			expected: fixtures.ImageVersion,
		},
		{
			input:    "1.0.0",
			expected: "1.0.0",
		},
	}

	for _, test := range tests {
		result := setDefaultIfEmpty(test.input)
		assert.Equal(t, test.expected, result, "Expected %s but got %s", test.expected, result)
	}
}

func TestGetImageURL(t *testing.T) {
	tests := []struct {
		name        string
		opts        *NewBaseContainerOpts
		expectedURL string
		expectedErr error
	}{
		{
			name: "Valid input with version",
			opts: &NewBaseContainerOpts{
				Image:   "golang",
				Version: "1.16",
			},
			expectedURL: "golang:1.16",
			expectedErr: nil,
		},
		{
			name: "Valid input without version",
			opts: &NewBaseContainerOpts{
				Image: "golang",
			},
			expectedURL: "golang:" + fixtures.ImageVersion,
			expectedErr: nil,
		},
		{
			name:        "Nil opts",
			opts:        nil,
			expectedURL: "",
			expectedErr: fmt.Errorf("failed to create base container: opts is nil"),
		},
		{
			name: "Empty image",
			opts: &NewBaseContainerOpts{
				Image: "",
			},
			expectedURL: "",
			expectedErr: fmt.Errorf("failed to create base container: image is empty"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := GetImageURL(test.opts)
			if test.expectedErr != nil {
				assert.EqualError(t, err, test.expectedErr.Error())
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, test.expectedURL, result)
		})
	}
}
