package sorting

// mergeSort divides the slice into several sub slices until each one have a single element.
func mergeSort(passedSlice []int) []int {
	if len(passedSlice) == 1 {
		return passedSlice
	}
	length := len(passedSlice)
	middle := length / 2

	left := passedSlice[:int(middle)]
	right := passedSlice[int(middle):]

	return merge(mergeSort(left), mergeSort(right))
}

// merge merges left and right sub slices.
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
