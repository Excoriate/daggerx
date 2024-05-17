package containerx

import (
	"dagger.io/dagger"
	"dagger.io/dagger/dag"
)

// WithGoBuildCache sets the build cache for the Go module.
// This function mounts a cache volume at the default path "/go/build-cache"
// and sets the environment variable "GOCACHE" to this path.
//
// Parameters:
//   - ctr: A pointer to the Dagger container to which the build cache will be added.
//
// Returns:
//   - A pointer to the Dagger container with the build cache set.
//
// Example:
//
//	// Create a new Dagger container
//	ctr := dagger.NewContainer()
//
//	// Set the Go build cache
//	updatedCtr := WithGoBuildCache(ctr)
//	// The container now has the build cache set at "/go/build-cache"
//	// with the environment variable "GOCACHE" pointing to it.
func WithGoBuildCache(ctr *dagger.Container) *dagger.Container {
	goBuildCache := dag.CacheVolume("gobuildcache")

	ctr = ctr.WithMountedCache("/go/build-cache", goBuildCache).
		WithEnvVariable("GOCACHE", "/go/build-cache")

	return ctr
}

// WithGoModCache sets the module cache for the Go module.
// This function mounts a cache volume at the default path "/go/pkg/mod"
// and sets the environment variable "GOMODCACHE" to this path.
//
// Parameters:
//   - ctr: A pointer to the Dagger container to which the module cache will be added.
//
// Returns:
//   - A pointer to the Dagger container with the module cache set.
//
// Example:
//
//	// Create a new Dagger container
//	ctr := dagger.NewContainer()
//
//	// Set the Go module cache
//	updatedCtr := WithGoModCache(ctr)
//	// The container now has the module cache set at "/go/pkg/mod"
//	// with the environment variable "GOMODCACHE" pointing to it.
func WithGoModCache(ctr *dagger.Container) *dagger.Container {
	goModCache := dag.CacheVolume("gomodcache")

	ctr = ctr.WithMountedCache("/go/pkg/mod", goModCache).
		WithEnvVariable("GOMODCACHE", "/go/pkg/mod")

	return ctr
}

// WithTerraformCache sets the Terraform cache.
// This function mounts two cache volumes: one at the default path "/root/.terraform.d"
// and another at the default path "/.terraform".
//
// Parameters:
//   - ctr: A pointer to the Dagger container to which the Terraform cache will be added.
//
// Returns:
//   - A pointer to the Dagger container with the Terraform cache set.
//
// Example:
//
//	// Create a new Dagger container
//	ctr := dagger.NewContainer()
//
//	// Set the Terraform cache
//	updatedCtr := WithTerraformCache(ctr)
//	// The container now has the Terraform cache set at "/root/.terraform.d"
//	// and another at "/.terraform".
func WithTerraformCache(ctr *dagger.Container) *dagger.Container {
	terraformCache := dag.CacheVolume("terraformcache")
	dotTerraformCache := dag.CacheVolume(".terraform")

	ctr = ctr.WithMountedCache("/root/.terraform.d", terraformCache).
		WithMountedCache("/.terraform", dotTerraformCache)

	return ctr
}

// WithTerragruntCache sets the Terragrunt cache.
// This function mounts a cache volume at the default path "/.terragrunt-cache".
//
// Parameters:
//   - ctr: A pointer to the Dagger container to which the Terragrunt cache will be added.
//
// Returns:
//   - A pointer to the Dagger container with the Terragrunt cache set.
//
// Example:
//
//	// Create a new Dagger container
//	ctr := dagger.NewContainer()
//
//	// Set the Terragrunt cache
//	updatedCtr := WithTerragruntCache(ctr)
//	// The container now has the Terragrunt cache set at "/.terragrunt-cache".
func WithTerragruntCache(ctr *dagger.Container) *dagger.Container {
	terragruntCache := dag.CacheVolume(".terragrunt-cache")

	ctr = ctr.WithMountedCache("/.terragrunt-cache", terragruntCache)

	return ctr
}
