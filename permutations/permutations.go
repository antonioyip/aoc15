package permutations

// Generate all permutations for a list of items
// values - list of items to permutate
// return - list of all permutations of values
func GeneratePermutations(values []string) [][]string {
	return permutations(len(values), values)
}

// Heap's algorithm
func permutations(length int, values []string) [][]string {
	result := [][]string{}
	if length == 1 {
		// force a deep copy
		temp := make([]string, len(values))
		copy(temp, values)
		result = append(result, temp)
	} else {
		for i := 0; i < length-1; i++ {
			result = append(result, permutations(length-1, values)...)
			if length%2 == 0 {
				values[i], values[length-1] = values[length-1], values[i]
			} else {
				values[0], values[length-1] = values[length-1], values[0]
			}
		}
		result = append(result, permutations(length-1, values)...)
	}
	return result
}
