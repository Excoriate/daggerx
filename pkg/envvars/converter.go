package envvars

import (
	"errors"
	"fmt"
	"github.com/Excoriate/daggerx/pkg/types"
	"strings"
)

// ToDaggerEnvVarsFromStr converts a comma-separated string of key=value pairs into a map.
// It ensures all entries are valid and handles empty strings gracefully.
// This function converts a string of key=value pairs into a map of strings.
// Example input: "key1=value1,key2=value2,key3=value3"
func ToDaggerEnvVarsFromStr(envVars string) (map[string]string, error) {
	if envVars == "" {
		return nil, errors.New("input string is empty")
	}

	envVarsMap := make(map[string]string)
	parts := strings.Split(envVars, ",")
	for _, envVar := range parts {
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

		envVarsMap[key] = value
	}
	return envVarsMap, nil
}

// ToDaggerEnvVarsFromMap converts a map of environment variables into a slice of DaggerEnvVars.
// It ensures all entries are valid and handles empty maps gracefully.
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
