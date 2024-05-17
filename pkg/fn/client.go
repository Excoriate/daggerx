package fn

import "dagger.io/dagger"

type Client interface {
	dagger.DaggerObject
	New(image, version string, ctr *dagger.Container) (interface{}, error)
}
