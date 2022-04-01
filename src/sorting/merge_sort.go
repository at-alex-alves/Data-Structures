package sorting

/*
	Divides the slice into several sub slices until each one have a single element.

    Args:
        passedSlice ([]int): The slice to be sorted.

    Returns:
        []int: The sorted slice.
*/
func mergeSort(passedSlice []int) []int {
	if len(passedSlice) == 1 {
		return passedSlice
	}
	length := len(passedSlice)
	midle := length / 2

	left := passedSlice[:int(midle)]
	right := passedSlice[int(midle):]

	return merge(mergeSort(left), mergeSort(right))
}

/*
	Merges left and right sub slices.

    Args:
		left ([]int): The left sub slice.
		right ([]int): The right sub slice.

	Returns:
		[]int: The merged slice.
*/
func merge(left, right []int) []int {
	result := []int{}
	leftIndex := 0
	rightIndex := 0

	for leftIndex < len(left) && rightIndex < len(right) {
		if left[leftIndex] < right[rightIndex] {
			result = append(result, left[leftIndex])
			leftIndex += 1
		} else {
			result = append(result, right[rightIndex])
			rightIndex += 1
		}
	}

	return append(append(result, left[leftIndex:]...), right[rightIndex:]...)
}
