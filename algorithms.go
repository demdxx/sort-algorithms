/**
 * Sorting algorithms
 * @autor Dmitry Ponomarev <demdxx@gmail.com>
 * @year 2013
 */

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

// Combined Selection & Bubble sort
func SelectionBubbleSort(data sort.Interface) {
	n := data.Len() - 1
	b := false
	for i := 0; i < n; i++ {
		min := i
		for j := 0; j < n-i; j++ {
			if data.Less(j+1, j) {
				data.Swap(j+1, j)
				b = true
			}
			if data.Less(j, min) {
				min = j
			}
		}
		if !b {
			break
		}
		if min != i { /* Exchange elements */
			data.Swap(i, min)
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

// Dwarf sorting (Gnome Sort) O(n^2)
func GnomeSort(data sort.Interface) {
	size := data.Len()
	i := 1
	j := 2
	for {
		if i >= size {
			break
		}
		if data.Less(i, i-1) {
			i = j
			j++
		} else {
			data.Swap(i, i-1)
			i--
			if 0 == i {
				i = j
				j++
			}
		}
	}
}

// Gnome Sort Optimization from me @autor demdxx
func OptimizedGnomeSort(data sort.Interface) {
	size := data.Len()
	i := 1
	j := 2
	for {
		if i >= size {
			break
		}

		// Optimization block {{{
		if data.Less(i, 0) {
			data.Swap(i, 0)
			i = j
			j++
		} else if i > 1 {
			x := i / 2
			y := i
			for {
				if data.Less(i, x) {
					y = x
					if x > 2 {
						x /= 2
						continue
					}
				}
				break
			}
			if y != i {
				data.Swap(i, y)
				i = y
				continue
			}
		}
		// }}} Optimization block

		if data.Less(i, i-1) {
			i = j
			j++
		} else {
			data.Swap(i, i-1)
			i--
			if 0 == i {
				i = j
				j++
			}
		}
	}
}

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
