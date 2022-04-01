package sorting

/*
	Given a slice, sort it in ascending order using the insertion sort algorithm.

	Args:
		passedSlice ([]int): The slice to be sorted.

	Returns:
		[]int: Sorted slice.
*/
func insertionSort(passedSlice []int) []int {
	for i := 0; i < len(passedSlice); i++ {
		for j := i; j > 0 && passedSlice[j-1] > passedSlice[j]; j-- {
			passedSlice[j], passedSlice[j-1] = passedSlice[j-1], passedSlice[j]
		}
	}
	return passedSlice
}
