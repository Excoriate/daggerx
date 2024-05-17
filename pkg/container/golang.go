package container

import "dagger.io/dagger"

const (
	// GolangAlpineImage is the name of the Golang Alpine base image.
	GolangAlpineImage = "golang"
)

// NewGolangAlpineContainer creates a new container based on the Golang Alpine image with the specified version.
// It returns an error if the version is not provided or if the base container cannot be created.
//
// Parameters:
//   - version: The version of the Golang Alpine image to use. If empty, a default version is used.
//
// Returns:
//   - A pointer to the Dagger container created from the specified Golang Alpine image and version.
//   - An error if the version is not provided or if the base container cannot be created.
//
// Example:
//
//	// Create a new Golang Alpine container with a specific version
//	container, err := NewGolangAlpineContainer("1.16-alpine")
//	if err != nil {
//	    // handle error
//	}
//	// Use the container, e.g., fmt.Println(container)
//
//	// Create a new Golang Alpine container with the default version
//	container, err = NewGolangAlpineContainer("")
//	if err != nil {
//	    // handle error
//	}
//	// Use the container, e.g., fmt.Println(container)
func NewGolangAlpineContainer(version string) (*dagger.Container, error) {
	ctr, err := NewBase(&NewBaseContainerOpts{
		Image:   GolangAlpineImage,
		Version: version,
	})

	if err != nil {
		return nil, err
	}

	return ctr, nil
}
