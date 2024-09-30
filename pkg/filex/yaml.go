// Package filex provides utility functions for working with files and directories,
// including validation of file extensions, existence, content, and structure.
// It offers functions to check if a file has a valid YAML extension, if it exists,
// if it has content, and if it can be properly unmarshaled into a provided struct.
//
// The package includes the following main functions:
//
//  1. ValidateYAMLExtension(filename string) bool:
//     Checks if the given file has a .yaml or .yml extension.
//  2. ValidateYAMLExists(filename string) bool:
//     Checks if the given YAML file exists.
//  3. ValidateYAMLHasContent(filename string) (bool, error):
//     Checks if the given YAML file is not empty and returns a boolean indicating
//     whether the file has content and an error if any occurred during reading the file.
//  4. ValidateYAMLStructure(filename string, out interface{}) error:
//     Checks if the given YAML file can be properly unmarshaled into the provided struct.
//  5. ValidateYAML(filename string, out interface{}) error:
package filex

import (
	"errors"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

// ValidateYAMLExtension checks if the given file has a .yaml or .yml extension.
// It returns true if the file has a valid YAML extension, otherwise false.
//
// Example usage:
//
//	valid := ValidateYAMLExtension("config.yaml")
//	fmt.Println(valid) // Output: true
func ValidateYAMLExtension(filename string) bool {
	ext := filepath.Ext(filename)
	return ext == ".yaml" || ext == ".yml"
}

// ValidateYAMLExists checks if the given YAML file exists.
// It returns true if the file exists, otherwise false.
//
// Example usage:
//
//	exists := ValidateYAMLExists("config.yaml")
//	fmt.Println(exists) // Output: true
func ValidateYAMLExists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

// ValidateYAMLHasContent checks if the given YAML file is not empty.
// It returns a boolean indicating whether the file has content and an error if any occurred during reading the file.
//
// Example usage:
//
//	hasContent, err := ValidateYAMLHasContent("config.yaml")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Println(hasContent) // Output: true
func ValidateYAMLHasContent(filename string) (bool, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return false, err
	}
	return len(content) > 0, nil
}

// ValidateYAMLStructure checks if the given YAML file can be properly unmarshaled into the provided struct.
// It returns an error if the file cannot be read or if the content cannot be unmarshaled into the provided struct.
//
// Example usage:
//
//	var config Config
//	err := ValidateYAMLStructure("config.yaml", &config)
//	if err != nil {
//	    log.Fatal(err)
//	}
func ValidateYAMLStructure(filename string, out interface{}) error {
	content, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(content, out)
	if err != nil {
		return err
	}

	return nil
}

// ValidateYAML performs all YAML validations: extension, existence, content, and structure.
// It returns an error if any of the validations fail.
//
// Example usage:
//
//	var config Config
//	err := ValidateYAML("config.yaml", &config)
//	if err != nil {
//	    log.Fatal(err)
//	}
func ValidateYAML(filename string, out interface{}) error {
	if !ValidateYAMLExtension(filename) {
		return errors.New("invalid YAML file extension")
	}

	if !ValidateYAMLExists(filename) {
		return errors.New("YAML file does not exist")
	}

	hasContent, err := ValidateYAMLHasContent(filename)
	if err != nil {
		return err
	}
	if !hasContent {
		return errors.New("YAML file is empty")
	}

	return ValidateYAMLStructure(filename, out)
}
