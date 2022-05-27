package sorting

// quickSort selects a pivot element from the slice and partitioning the other elements into two sub slice,
// according to whether they are less than or greater than the pivot.
func quickSort(passedSlice []int, left, right int) []int {
	if left < right {
		var partitionIndex int

		passedSlice, partitionIndex = partition(passedSlice, left, right)
		passedSlice = quickSort(passedSlice, left, partitionIndex-1)
		passedSlice = quickSort(passedSlice, partitionIndex+1, right)
	}

	return passedSlice
}

// partition swaps all elements less than the pivot value to the left of the pivot index, and all elements
// greater than the pivot value to the right of the pivot index.
func partition(passedSlice []int, left, right int) ([]int, int) {
	pivot := passedSlice[right]
	partitionIndex := left

	for i := left; i < right; i++ {
		if passedSlice[i] < pivot {
			passedSlice[partitionIndex], passedSlice[i] = passedSlice[i], passedSlice[partitionIndex]
			partitionIndex++
		}
	}

	passedSlice[partitionIndex], passedSlice[right] = passedSlice[right], passedSlice[partitionIndex]
	return passedSlice, partitionIndex
}
