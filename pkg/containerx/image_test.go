package containerx

import (
	"fmt"
	"testing"

	"github.com/Excoriate/daggerx/pkg/fixtures"
	"github.com/stretchr/testify/assert"
)

func TestSetDefaultIfEmpty(t *testing.T) {
	tests := []struct {
		inputVersion    string
		fallbackVersion string
		expected        string
	}{
		{
			inputVersion:    "",
			fallbackVersion: fixtures.ImageVersion,
			expected:        fixtures.ImageVersion,
		},
		{
			inputVersion:    "1.0.0",
			fallbackVersion: fixtures.ImageVersion,
			expected:        "1.0.0",
		},
		{
			inputVersion:    "",
			fallbackVersion: "",
			expected:        fixtures.ImageVersion,
		},
	}

	for _, test := range tests {
		result := SetDefaultImageVersionIfEmpty(test.inputVersion, test.fallbackVersion)
		assert.Equal(t, test.expected, result, "Expected %s but got %s", test.expected, result)
	}
}

func TestSetDefaultImageIfEmpty(t *testing.T) {
	tests := []struct {
		inputImage    string
		fallbackImage string
		expectedImage string
	}{
		{
			inputImage:    "",
			fallbackImage: "fallback-image",
			expectedImage: "fallback-image",
		},
		{
			inputImage:    "primary-image",
			fallbackImage: "fallback-image",
			expectedImage: "primary-image",
		},
		{
			inputImage:    "",
			fallbackImage: "",
			expectedImage: fixtures.Image,
		},
	}

	for _, test := range tests {
		result := SetDefaultImageNameIfEmpty(test.inputImage, test.fallbackImage)
		assert.Equal(t, test.expectedImage, result, "Expected %s but got %s", test.expectedImage, result)
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
				Image:           "golang",
				FallBackVersion: fixtures.ImageVersion,
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
			expectedErr: fmt.Errorf("failed to create base container: both image and fallback image are empty"),
		},
		{
			name: "Fallback image used",
			opts: &NewBaseContainerOpts{
				FallbackImage: "fallback-image",
				Version:       "1.0.0",
			},
			expectedURL: "fallback-image:1.0.0",
			expectedErr: nil,
		},
		{
			name: "Both image and fallback image used",
			opts: &NewBaseContainerOpts{
				Image:         "image-name",
				FallbackImage: "fallback-image",
				Version:       "2.0.0",
			},
			expectedURL: "image-name:2.0.0",
			expectedErr: nil,
		},
		{
			name: "Primary image empty, fallback image and version used",
			opts: &NewBaseContainerOpts{
				Image:           "",
				FallbackImage:   "fallback-image",
				FallBackVersion: "2.0.0",
			},
			expectedURL: "fallback-image:2.0.0",
			expectedErr: nil,
		},
		{
			name: "Both image and version empty, fallback image and version used",
			opts: &NewBaseContainerOpts{
				Image:           "",
				FallbackImage:   "fallback-image",
				Version:         "",
				FallBackVersion: "1.0.0",
			},
			expectedURL: "fallback-image:1.0.0",
			expectedErr: nil,
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
