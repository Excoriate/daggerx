package containerx

import (
	"dagger.io/dagger"
	"dagger.io/dagger/dag"
)

type Client interface {
	dagger.DaggerObject
	New(image, version string, ctr *dagger.Container) (interface{}, error)
}

type DaggerFnContainer interface {
	NewBaseContainer(opts *NewBaseContainerOpts) (any, error)
	NewGolangAlpineContainer(version string) (interface{}, error)
}

type DaggerFnContainerClient struct{}

// NewBaseContainerOpts contains options for creating a new base container.
type NewBaseContainerOpts struct {
	Image   string // The name of the base image.
	Version string // The version of the base image. If empty, a default version is used.
}

// NewBaseContainer NewBase creates a new base container from the specified image and version.
// It returns an error if the options are nil or the image name is empty.
//
// Parameters:
//   - opts: A pointer to NewBaseContainerOpts containing the image name and version.
//
// Returns:
//   - A pointer to the Dagger container created from the specified image and version.
//   - An error if the options are nil or the image name is empty.
//
// Example:
//
//	opts := &NewBaseContainerOpts{
//	    Image:   "golang",
//	    Version: "1.16",
//	}
//	container, err := NewBase(opts)
//	if err != nil {
//	    // handle error
//	}
//	// Use the container, e.g., fmt.Println(container)
func (c *DaggerFnContainerClient) NewBaseContainer(opts *NewBaseContainerOpts) (any, error) {
	imageURL, err := GetImageURL(opts)
	if err != nil {
		return nil, err
	}

	return dag.Container().From(imageURL), nil
}
