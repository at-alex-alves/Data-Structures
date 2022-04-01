package sorting

/*
	Sorts a slice of numbers in ascending order.

	Args:
		passedSlice ([]int): The slice to be sorted.

	Returns:
		[]int: Sorted slice.
*/
func selectionSort(passedSlice []int) []int {
	i := 0

	for i < len(passedSlice) {
		min := passedSlice[i]
		index := i

		for j := i + 1; j < len(passedSlice); j++ {
			if passedSlice[j] < min {
				index = j
				min = passedSlice[j]
			}
		}

		passedSlice[i], passedSlice[index] = passedSlice[index], passedSlice[i]
		i += 1
	}

	return passedSlice
}
