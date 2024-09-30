// Package envvars provides utilities for converting environment variables
// between different formats. This package is designed to facilitate the
// handling of environment variables in Go applications, ensuring that they
// can be easily converted from strings, maps, and other common data structures
// into a standardized format.
//
// The primary focus of this package is to provide functions that convert
// environment variables into slices of DaggerEnvVars, a custom type defined
// in the types package. These functions ensure that the environment variables
// are valid and handle edge cases such as empty strings or maps gracefully.
//
// Key Functions:
//
//   - ToDaggerEnvVarsFromStr: Converts a comma-separated string of key=value pairs
//     into a slice of DaggerEnvVars.
//
//   - ToDaggerEnvVarsFromMap: Converts a map of environment variables into a slice
//     of DaggerEnvVars.
//
// Example Usage:
//
// Converting a comma-separated string of environment variables:
//
//	envVarsStr := "FOO=bar,BAZ=qux"
//	envVars, err := ToDaggerEnvVarsFromStr(envVarsStr)
//	if err != nil {
//	    // handle error
//	}
//	// Use envVars, e.g., fmt.Println(envVars)
//
// Converting a map of environment variables:
//
//	envVarsMap := map[string]string{"FOO": "bar", "BAZ": "qux"}
//	envVarsSlice, err := ToDaggerEnvVarsFromMap(envVarsMap)
//	if err != nil {
//	    // handle error
//	}
//	// Use envVarsSlice, e.g., fmt.Println(envVarsSlice)
package envvars

import (
	"errors"
	"fmt"
	"github.com/Excoriate/daggerx/pkg/types"
	"strings"
)

// ToDaggerEnvVarsFromStr converts a comma-separated string of key=value pairs into a slice of DaggerEnvVars.
// It ensures all entries are valid and handles empty strings gracefully.
//
// Parameters:
//   - envVars: A comma-separated string of key=value pairs. For example: "key1=value1,key2=value2,key3=value3".
//
// Returns:
//   - A slice of DaggerEnvVars, each containing the name and value of an environment variable.
//   - An error if the input string is empty or if any of the key=value pairs are invalid.
//
// Example:
//
//	envVarsStr := "FOO=bar,BAZ=qux"
//	envVars, err := ToDaggerEnvVarsFromStr(envVarsStr)
//	if err != nil {
//	    // handle error
//	}
//	// Use envVars, e.g., fmt.Println(envVars)
func ToDaggerEnvVarsFromStr(envVars string) ([]types.DaggerEnvVars, error) {
	if envVars == "" {
		return nil, errors.New("input string is empty")
	}

	envVarsSlice := strings.Split(envVars, ",")
	return ToDaggerEnvVarsFromSlice(envVarsSlice)
}

// ToDaggerEnvVarsFromMap converts a map of environment variables into a slice of DaggerEnvVars.
// It ensures all entries are valid and handles empty maps gracefully.
//
// Parameters:
//   - envVarsMap: A map of environment variables where each key is a variable name and each value is the corresponding value.
//
// Returns:
//   - A slice of DaggerEnvVars, each containing the name and value of an environment variable.
//   - An error if the input map is empty or contains an empty key.
//
// Example:
//
//	envVarsMap := map[string]string{"FOO": "bar", "BAZ": "qux"}
//	envVarsSlice, err := ToDaggerEnvVarsFromMap(envVarsMap)
//	if err != nil {
//	    // handle error
//	}
//	// Use envVarsSlice, e.g., fmt.Println(envVarsSlice)
func ToDaggerEnvVarsFromMap(envVarsMap map[string]string) ([]types.DaggerEnvVars, error) {
	if len(envVarsMap) == 0 {
		return nil, errors.New("input map is empty")
	}

	var envVars []types.DaggerEnvVars
	for key, value := range envVarsMap {
		if key == "" {
			return nil, errors.New("found empty key in map")
		}
		envVars = append(envVars, types.DaggerEnvVars{
			Name:  key,
			Value: value,
		})
	}
	return envVars, nil
}

// ToDaggerEnvVarsFromSlice converts a slice of key=value strings into a slice of DaggerEnvVars.
// It validates each entry and skips invalid entries.
//
// Parameters:
//   - envVarsSlice: A slice of strings where each string is a key=value pair representing an environment variable.
//
// Returns:
//   - A slice of DaggerEnvVars, each containing the name and value of an environment variable.
//   - An error if the input slice is empty or if any of the key=value pairs are invalid.
//
// Example:
//
//	envVarsSlice := []string{"FOO=bar", "BAZ=qux"}
//	envVars, err := ToDaggerEnvVarsFromSlice(envVarsSlice)
//	if err != nil {
//	    // handle error
//	}
//	// Use envVars, e.g., fmt.Println(envVars)
func ToDaggerEnvVarsFromSlice(envVarsSlice []string) ([]types.DaggerEnvVars, error) {
	if len(envVarsSlice) == 0 {
		return nil, errors.New("input slice is empty")
	}

	var envVars []types.DaggerEnvVars
	for _, envVar := range envVarsSlice {
		if envVar == "" {
			continue
		}

		pair := strings.SplitN(envVar, "=", 2)
		if len(pair) != 2 {
			return nil, fmt.Errorf("invalid environment variable format: %s", envVar)
		}

		key, value := strings.TrimSpace(pair[0]), strings.TrimSpace(pair[1])
		if key == "" {
			return nil, fmt.Errorf("empty key in environment variable: %s", envVar)
		}

		envVars = append(envVars, types.DaggerEnvVars{
			Name:  key,
			Value: value,
		})
	}
	return envVars, nil
}
