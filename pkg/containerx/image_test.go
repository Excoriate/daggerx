package containerx

import (
	"fmt"
	"sync"
	"testing"

	"github.com/Excoriate/daggerx/pkg/fixtures"
	"github.com/stretchr/testify/assert"
)

func TestSetDefaultImageNameIfEmpty(t *testing.T) {
	tests := []struct {
		name          string
		image         string
		fallbackImage string
		expected      string
	}{
		{"Primary image set", "ubuntu", "alpine", "ubuntu"},
		{"Fallback image used", "", "alpine", "alpine"},
		{"Default image used", "", "", fixtures.Image},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SetDefaultImageNameIfEmpty(tt.image, tt.fallbackImage)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestSetDefaultImageVersionIfEmpty(t *testing.T) {
	tests := []struct {
		name            string
		version         string
		fallbackVersion string
		expected        string
	}{
		{"Primary version set", "1.0", "2.0", "1.0"},
		{"Fallback version used", "", "2.0", "2.0"},
		{"Default version used", "", "", "latest"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SetDefaultImageVersionIfEmpty(tt.version, tt.fallbackVersion)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestGetImageURL(t *testing.T) {
	tests := []struct {
		name        string
		opts        *NewBaseContainerOpts
		expected    string
		expectedErr error
	}{
		{
			name: "Valid options",
			opts: &NewBaseContainerOpts{
				Image:   "ubuntu",
				Version: "20.04",
			},
			expected:    "ubuntu:20.04",
			expectedErr: nil,
		},
		{
			name: "Fallback image and version",
			opts: &NewBaseContainerOpts{
				FallbackImage:   "alpine",
				FallBackVersion: "3.14",
			},
			expected:    "alpine:3.14",
			expectedErr: nil,
		},
		{
			name:        "Nil options",
			opts:        nil,
			expected:    "",
			expectedErr: fmt.Errorf("failed to create base container: opts is nil"),
		},
		{
			name: "Empty images",
			opts: &NewBaseContainerOpts{
				Image:         "",
				FallbackImage: "",
			},
			expected:    fmt.Sprintf("%s:latest", fixtures.Image),
			expectedErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := GetImageURL(tt.opts)
			if tt.expectedErr != nil {
				assert.Error(t, err)
				assert.Equal(t, tt.expectedErr.Error(), err.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, result)
			}
		})
	}
}

func TestValidateImageURL(t *testing.T) {
	tests := []struct {
		name     string
		imageURL string
		want     bool
		wantErr  string
	}{
		{"Empty URL", "", false, "image URL cannot be empty"},
		{"Simple image", "ubuntu", true, ""},
		{"Image with tag", "ubuntu:20.04", true, ""},
		{"Image with digest", "ubuntu@sha256:1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef", true, ""},
		{"Image with tag and digest", "ubuntu:20.04@sha256:abcdef1234567890abcdef1234567890abcdef1234567890abcdef1234567890", true, ""},
		{"Docker Hub image", "docker.io/library/ubuntu:20.04", true, ""},
		{"ECR image", "123456789.dkr.ecr.us-west-2.amazonaws.com/my-app:v1.0.0", true, ""},
		{"GHCR image", "ghcr.io/username/repo:tag", true, ""},
		{"GCR image", "gcr.io/project-id/image:tag", true, ""},
		{"Quay.io image", "quay.io/username/repo:tag", true, ""},
		{"Invalid registry", "invalid..registry/image:tag", false, "invalid registry: invalid..registry"},
		{"Invalid repository name", "invalid@repo:tag", false, "invalid repository name: invalid@repo:tag"},
		{"Invalid tag", "registry/repo:invalid_tag!", false, "invalid tag: invalid_tag!"},
		{"Invalid digest", "registry/repo@sha256:invalid", false, "invalid digest: sha256:invalid"},
		{"Too many components", "a/b/c/d/e:tag", false, "too many components in image URL"},
		{"Valid long path", "registry.com:8080/path/to/repo:tag", true, ""},
		{"Valid ECR-like URL", "123456789.dkr.ecr.us-west-2.amazonaws.com/my-repo:latest", true, ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ValidateImageURL(tt.imageURL)
			if (err != nil) != (tt.wantErr != "") {
				t.Errorf("ValidateImageURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil && err.Error() != tt.wantErr {
				t.Errorf("ValidateImageURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ValidateImageURL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidateImageURLConcurrent(t *testing.T) {
	validURLs := []string{
		"ubuntu:latest",
		"golang:1.16",
		"registry.gitlab.com/group/project/image:tag",
		"docker.io/library/redis:6.2",
		"123456789012.dkr.ecr.us-west-2.amazonaws.com/my-app:latest",
		"public.ecr.aws/registry/my-app:latest",
		"gcr.io/project-id/my-app:latest",
		"my-registry.com:5000/my-app:latest",
	}

	invalidURLs := []string{
		"invalid@repo:tag",      // Contains '@' in repository name
		"registry.com/my_repo:", // Empty tag
		"::",                    // Completely invalid format
		"registry.com/my-app@sha256:not-a-valid-hash", // Invalid digest
	}

	var validCount, errorCount int
	var mu sync.Mutex
	var wg sync.WaitGroup

	for _, url := range validURLs {
		wg.Add(1)
		go func(u string) {
			defer wg.Done()
			valid, err := ValidateImageURL(u)
			if !valid || err != nil {
				t.Errorf("Expected valid URL, got invalid: %s, error: %v", u, err)
			}
			mu.Lock()
			validCount++
			mu.Unlock()
		}(url)
	}

	for _, url := range invalidURLs {
		wg.Add(1)
		go func(u string) {
			defer wg.Done()
			valid, err := ValidateImageURL(u)
			if valid || err == nil {
				t.Errorf("Expected invalid URL, got valid: %s", u)
			}
			mu.Lock()
			errorCount++
			mu.Unlock()
		}(url)
	}

	wg.Wait()

	if validCount != len(validURLs) {
		t.Errorf("Expected %d valid URLs, got %d", len(validURLs), validCount)
	}

	if errorCount != len(invalidURLs) {
		t.Errorf("Expected %d errors, got %d", len(invalidURLs), errorCount)
	}
}

func BenchmarkValidateImageURL(b *testing.B) {
	urls := []string{
		"ubuntu:20.04",
		"docker.io/library/alpine",
		"123456789.dkr.ecr.us-west-2.amazonaws.com/my-app:v1.0.0",
		"ghcr.io/username/repo:tag",
		"gcr.io/project-id/image:tag",
		"quay.io/username/repo:tag",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, url := range urls {
			ValidateImageURL(url)
		}
	}
}
