package merger

// MergeSlices merges slices of strings.
func MergeSlices(slices ...[]string) []string {
	var merged []string
	for _, slice := range slices {
		if slice != nil {
			merged = append(merged, slice...)
		}
	}
	return merged
}
