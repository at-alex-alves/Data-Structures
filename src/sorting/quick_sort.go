package sorting

/*
	Given a slice, selects a pivot element from the slice and partitioning the other elements into two sub slice,
	according to whether they are less than or greater than the pivot.

	Args:
       passedSlice ([]int): The slice to be sorted.
       left (int): The first index of the slice.
       right (int): The last index of the slice.

	Returns:
       []int: The sorted slice.
*/
func quickSort(passedSlice []int, left, right int) []int {
	if left < right {
		var partitionIndex int

		passedSlice, partitionIndex = partition(passedSlice, left, right)
		passedSlice = quickSort(passedSlice, left, partitionIndex-1)
		passedSlice = quickSort(passedSlice, partitionIndex+1, right)
	}

	return passedSlice
}

/*
	Given a slice, a left and right index, swaps all elements less than the pivot value to the left
	of the pivot index, and all elements greater than the pivot value to the right of the pivot index.

	Args:
		passedSlice ([]int): The slice to be sorted.
		left (int): The first index of the slice.
		right (int): The last index of the slice.

	Returns:
		[]int: Sorted slice.
		int: The index of the next partition.
*/
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
