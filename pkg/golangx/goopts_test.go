package golangx

import (
	"testing"

	"github.com/Excoriate/daggerx/pkg/types"
	"github.com/stretchr/testify/assert"
)

func TestWithGoCgoEnabled(t *testing.T) {
	expected := types.DaggerEnvVars{
		Name:   "CGO_ENABLED",
		Value:  "1",
		Expand: false,
	}

	result := WithGoCgoEnabled()

	assert.Equal(t, expected, result, "Expected CGO_ENABLED=1, but got a different value")
}

func TestWithGoCgoDisabled(t *testing.T) {
	expected := types.DaggerEnvVars{
		Name:   "CGO_ENABLED",
		Value:  "0",
		Expand: false,
	}

	result := WithGoCgoDisabled()

	assert.Equal(t, expected, result, "Expected CGO_ENABLED=0, but got a different value")
}

func TestWithGoPlatform(t *testing.T) {
	tests := []struct {
		platform string
		expected []types.DaggerEnvVars
	}{
		{
			platform: "linux/amd64",
			expected: []types.DaggerEnvVars{
				{
					Name:   "GOOS",
					Value:  "linux",
					Expand: false,
				},
				{
					Name:   "GOARCH",
					Value:  "amd64",
					Expand: false,
				},
			},
		},
		{
			platform: "linux/arm/v7",
			expected: []types.DaggerEnvVars{
				{
					Name:   "GOOS",
					Value:  "linux",
					Expand: false,
				},
				{
					Name:   "GOARCH",
					Value:  "arm",
					Expand: false,
				},
				{
					Name:   "GOARM",
					Value:  "v7",
					Expand: false,
				},
			},
		},
	}

	for _, test := range tests {
		result := WithGoPlatform(test.platform)
		assert.Equal(t, test.expected, result, "Expected different environment variables for platform %s", test.platform)
	}
}
