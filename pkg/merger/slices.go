package merger

// MergeSlices merges multiple slices of strings into a single slice.
// It skips any nil slices and handles empty slices gracefully.
//
// Parameters:
//   - slices: A variadic parameter that takes multiple slices of strings.
//
// Returns:
//   - A single slice of strings containing all the elements from the input slices, in the order they were provided.
//
// Example:
//
//	// Merge multiple slices of strings
//	slice1 := []string{"a", "b"}
//	slice2 := []string{"c", "d"}
//	slice3 := []string{"e", "f"}
//
//	merged := MergeSlices(slice1, slice2, slice3)
//	// merged now contains: ["a", "b", "c", "d", "e", "f"]
//
//	// Merge slices with nil and empty slices
//	slice4 := []string{"g", "h"}
//	emptySlice := []string{}
//
//	merged = MergeSlices(nil, slice1, emptySlice, slice4)
//	// merged now contains: ["a", "b", "g", "h"]
func MergeSlices(slices ...[]string) []string {
	var merged []string
	for _, slice := range slices {
		if slice != nil {
			merged = append(merged, slice...)
		}
	}
	return merged
}
