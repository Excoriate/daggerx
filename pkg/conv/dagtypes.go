package conv

import "dagger.io/dagger"

func ToDaggerPlatform(input interface{}) dagger.Platform {
	platform, ok := input.(dagger.Platform)
	if !ok {
		return ""
	}

	return platform
}

func ToDaggerDir(input any) dagger.Directory {
	dir, ok := input.(dagger.Directory)
	if !ok {
		return dagger.Directory{}
	}

	return dir
}
