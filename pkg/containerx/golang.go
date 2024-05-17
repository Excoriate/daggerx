package containerx

const (
	GolangAlpineImage = "golang"
)

func (c *DaggerFnContainerClient) NewGolangAlpineContainer(version string) (any, error) {
	ctr, err := c.NewBaseContainer(&NewBaseContainerOpts{
		Image:   GolangAlpineImage,
		Version: version,
	})

	if err != nil {
		return nil, err
	}

	return ctr, nil
}
