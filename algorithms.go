package algorithms

import (
	"sort"
)

// Stable algorithms

// Selection sort O(n^2)
func SelectionSort(data sort.Interface) {
	n := data.Len()
	for i := 0; i < n; i++ {
		min := i /* position of minimal element */

		/* Finde minimal element */
		for j := i + 1; j < n; j++ {
			if data.Less(j, min) {
				min = j
			}
		}
		if min != i { /* Exchange elements */
			data.Swap(i, min)
		}
	}
}

// Bubble sort O(n^2)
func BubbleSort(data sort.Interface) {
	n := data.Len() - 1
	b := false
	for i := 0; i < n; i++ {
		for j := 0; j < n-i; j++ {
			if data.Less(j+1, j) {
				data.Swap(j+1, j)
				b = true
			}
		}
		if !b {
			break
		}
		b = false
	}
}

// Cocktail sort, bidirectional bubble sort O(n^2)
func ShakerSort(data sort.Interface) {
	left := 1
	right := data.Len() - 1

	for {
		// Move the array to the beginning of the "light elements"
		for i := right; i >= left; i-- {
			if data.Less(i, i-1) {
				data.Swap(i, i-1)
			}
		}
		left++

		// We shift to the end of the array "heavy elements"
		for i := left; i <= right; i++ {
			if data.Less(i, i-1) {
				data.Swap(i, i-1)
			}
		}
		right--
		if left > right {
			break
		}
	}
}

// Dwarf sorting
// Insertion sort
// Merge sort
// Tree sort
// Timsort
// Counting sort
// Bucket sort

// Unstable algorithms

// Shell sort
// Comb sort
// Smoothsort
// Quicksort
// Introsort
// Patience sorting
// Stooge sort
// Radix sort

// Impractical sorting algorithms

// Bogosort
// Sort permutation
// Stupid sort
// Bead Sort
// Pancake sorting

// The algorithms that are not based on comparisons

// Bucket sort
// Radix sort
// Counting sort

// Other

// The topological sorting
// External sorting
