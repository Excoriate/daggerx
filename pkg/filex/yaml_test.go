package filex

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestValidateYAMLExtension(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		want     bool
	}{
		{"Valid YAML", "test.yaml", true},
		{"Valid YML", "test.yml", true},
		{"Invalid extension", "test.txt", false},
		{"No extension", "test", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidateYAMLExtension(tt.filename); got != tt.want {
				t.Errorf("ValidateYAMLExtension() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidateYAMLExists(t *testing.T) {
	// Create a temporary YAML file
	tmpfile, err := os.CreateTemp("", "test*.yaml")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())

	tests := []struct {
		name     string
		filename string
		want     bool
	}{
		{"Existing file", tmpfile.Name(), true},
		{"Non-existing file", "non_existing.yaml", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidateYAMLExists(tt.filename); got != tt.want {
				t.Errorf("ValidateYAMLExists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidateYAMLHasContent(t *testing.T) {
	// Create temporary files
	emptyFile, err := ioutil.TempFile("", "empty*.yaml")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(emptyFile.Name())

	contentFile, err := os.CreateTemp("", "content*.yaml")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(contentFile.Name())

	// Write content to the non-empty file
	if _, err := contentFile.Write([]byte("key: value")); err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name     string
		filename string
		want     bool
		wantErr  bool
	}{
		{"Empty file", emptyFile.Name(), false, false},
		{"File with content", contentFile.Name(), true, false},
		{"Non-existing file", "non_existing.yaml", false, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ValidateYAMLHasContent(tt.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateYAMLHasContent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ValidateYAMLHasContent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidateYAMLStructure(t *testing.T) {
	type TestStruct struct {
		Key string `yaml:"key"`
	}

	// Create temporary files
	validFile, err := os.CreateTemp("", "valid*.yaml")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(validFile.Name())

	invalidFile, err := os.CreateTemp("", "invalid*.yaml")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(invalidFile.Name())

	// Write content to the files
	if _, err := validFile.Write([]byte("key: value")); err != nil {
		t.Fatal(err)
	}
	if _, err := invalidFile.Write([]byte("invalid: - yaml: content")); err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name     string
		filename string
		out      interface{}
		wantErr  bool
	}{
		{"Valid YAML", validFile.Name(), &TestStruct{}, false},
		{"Invalid YAML", invalidFile.Name(), &TestStruct{}, true},
		{"Non-existing file", "non_existing.yaml", &TestStruct{}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateYAMLStructure(tt.filename, tt.out); (err != nil) != tt.wantErr {
				t.Errorf("ValidateYAMLStructure() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateYAML(t *testing.T) {
	type TestStruct struct {
		Key string `yaml:"key"`
	}

	// Create temporary files
	validFile, err := os.CreateTemp("", "valid*.yaml")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(validFile.Name())

	invalidExtFile, err := os.CreateTemp("", "invalid*.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(invalidExtFile.Name())

	emptyFile, err := os.CreateTemp("", "empty*.yaml")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(emptyFile.Name())

	// Write content to the valid file
	if _, err := validFile.Write([]byte("key: value")); err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name     string
		filename string
		out      interface{}
		wantErr  bool
	}{
		{"Valid YAML", validFile.Name(), &TestStruct{}, false},
		{"Invalid extension", invalidExtFile.Name(), &TestStruct{}, true},
		{"Empty file", emptyFile.Name(), &TestStruct{}, true},
		{"Non-existing file", "non_existing.yaml", &TestStruct{}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateYAML(tt.filename, tt.out); (err != nil) != tt.wantErr {
				t.Errorf("ValidateYAML() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
