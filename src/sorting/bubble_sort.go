package sorting

/*
	Given a slice, sort it in swapping adjacent elements if they are in the wrong order.

    Args:
		passedSlice ([]int): The slice to be sorted.

    Returns:
		[]int: Sorted slice.
*/
func bubbleSort(passedSlice []int) []int {
	moviment := 0

	// Loops the same number of time as the number of elements.
	for moviment < len(passedSlice) {
		// Loops throw all elements.
		for i := 0; i < len(passedSlice)-1; i++ {
			// Checks if the next element is smaller.
			if passedSlice[i] > passedSlice[i+1] {
				passedSlice[i], passedSlice[i+1] = passedSlice[i+1], passedSlice[i]
			}
		}
		moviment += 1
	}
	return passedSlice
}
