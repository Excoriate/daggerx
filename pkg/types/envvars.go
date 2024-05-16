package types

// DaggerEnvVars represents the environment variables for a Dagger task
type DaggerEnvVars struct {
	// Name is the name of the environment variable
	Name string
	// Value is the value of the environment variable
	Value string
	// Expand is a boolean that determines whether to expand the value of the environment variable
	Expand bool
}
