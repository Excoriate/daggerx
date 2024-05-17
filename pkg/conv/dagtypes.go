package conv

// ToAnyType attempts to convert an input of any type to the specified type T.
// It returns a pointer to the converted value if successful, or nil if the conversion fails.
//
// Parameters:
//   - input: An interface{} representing the value to be converted.
//
// Returns:
//   - A pointer to the converted value of type T, or nil if the conversion fails.
//
// Example:
//
//	var input interface{} = 123
//	result := ToAnyType[int](input)
//	if result != nil {
//	    fmt.Println(*result) // Output: 123
//	} else {
//	    fmt.Println("Conversion failed")
//	}
func ToAnyType[T any](input interface{}) *T {
	// Check if the input is already of type T.
	if v, ok := input.(T); ok {
		return &v
	}

	// Check if the input is a pointer to type T.
	if v, ok := input.(*T); ok {
		return v
	}

	// Handle other potential cases.
	return nil
}
