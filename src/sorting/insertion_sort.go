package sorting

// insertionSort sorts a given a slice it in ascending order using the insertion sort algorithm.
func insertionSort(passedSlice []int) []int {
	for i := 0; i < len(passedSlice); i++ {
		for j := i; j > 0 && passedSlice[j-1] > passedSlice[j]; j-- {
			passedSlice[j], passedSlice[j-1] = passedSlice[j-1], passedSlice[j]
		}
	}
	return passedSlice
}
