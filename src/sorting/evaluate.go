package sorting

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	bubbleSortTime    int
	insertionSortTime int
	mergeSortTime     int
	quickSortTime     int
	selectionSortTime int
)

// run calculates the processing time for each algorithm based on the generated slice.
func run(sliceSize int) {
	rand.Seed(time.Now().Unix())
	slice := rand.Perm(sliceSize)

	start := time.Now()
	bubbleSort(slice)
	end := time.Since(start).Nanoseconds()

	bubbleSortTime += int(end)

	start = time.Now()
	insertionSort(slice)
	end = time.Since(start).Nanoseconds()

	insertionSortTime += int(end)

	start = time.Now()
	mergeSort(slice)
	end = time.Since(start).Nanoseconds()

	mergeSortTime += int(end)

	start = time.Now()
	quickSort(slice, 0, len(slice)-1)
	end = time.Since(start).Nanoseconds()

	quickSortTime += int(end)

	start = time.Now()
	selectionSort(slice)
	end = time.Since(start).Nanoseconds()

	selectionSortTime += int(end)
}

// EvaluatePerformance evaluates the performance by getting the mean value of the execution of all algorithms with the passed slice size
func EvaluatePerformance(arraySizeFrom, arraySizeTo int) {
	bubbleSortTime = 0
	insertionSortTime = 0
	mergeSortTime = 0
	quickSortTime = 0
	selectionSortTime = 0

	for i := arraySizeFrom; i <= arraySizeTo; i++ {
		run(i)
	}

	fmt.Printf(
		"\nSorting Algorithms Average in Nanoseconds:\n\tBubble Sort: %v\n\tInsertion Sort: %v\n\tMerge Sort: %v\n\tQuick Sort: %v\n\tSelection Sort: %v\n",
		bubbleSortTime/(arraySizeTo-arraySizeFrom),
		insertionSortTime/(arraySizeTo-arraySizeFrom),
		mergeSortTime/(arraySizeTo-arraySizeFrom),
		quickSortTime/(arraySizeTo-arraySizeFrom),
		selectionSortTime/(arraySizeTo-arraySizeFrom),
	)
}
