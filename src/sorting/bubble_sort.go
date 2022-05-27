package sorting

// bubbleSort sorts a given a slice it in swapping adjacent elements if they are in the wrong order.
func bubbleSort(passedSlice []int) []int {
	movement := 0

	// Loops the same number of time as the number of elements.
	for movement < len(passedSlice) {
		// Loops throw all elements.
		for i := 0; i < len(passedSlice)-1; i++ {
			// Checks if the next element is smaller.
			if passedSlice[i] > passedSlice[i+1] {
				passedSlice[i], passedSlice[i+1] = passedSlice[i+1], passedSlice[i]
			}
		}
		movement += 1
	}
	return passedSlice
}
