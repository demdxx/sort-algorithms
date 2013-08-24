/**
 * Sorting algorithms
 * @autor Dmitry Ponomarev <demdxx@gmail.com>
 * @year 2013
 */

package algorithms

import (
	"sort"
)

type Any interface{}
type AnyInt struct {
	value int
}

func GetAnyInt(val int) AnyInt {
	var v AnyInt
	v.value = val
	return v
}

// Extend sort interface
type SortInterface interface {
	// Len is the number of elements in the collection.
	Len() int
	// Less returns whether the element with index i should sort
	// before the element with index j.
	Less(i, j int) bool
	// Swap swaps the elements with indexes i and j.
	Swap(i, j int)
	// Get object by index
	Get(i int) Any
	// Set value by index
	Set(i int, v Any)
	// Set value by index
	SetByIndex(i, j int)
	// If object less object by index
	LessValueIndex(v Any, i int) bool
	// If object less or equal object by index
	LessOrEqValueIndex(v Any, i int) bool
	// Clone with size
	Clone(count int) SortInterface
}

type IntSlice []int

// Len is the number of elements in the collection.
func (p IntSlice) Len() int { return len(p) }

// Less returns whether the element with index i should sort
// before the element with index j.
func (p IntSlice) Less(i, j int) bool { return p[i] < p[j] }

// Swap swaps the elements with indexes i and j.
func (p IntSlice) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

// Get object by index
func (p IntSlice) Get(i int) Any { return GetAnyInt(p[i]) }

// Set value by Any
func (p IntSlice) Set(i int, v Any) { p[i] = v.(AnyInt).value }

// Set value by index
func (p IntSlice) SetByIndex(i, j int) { p[i] = p[j] }

// If object less object by index
func (p IntSlice) LessValueIndex(v Any, i int) bool { return p[i] > v.(AnyInt).value }

// If object less or equal object by index
func (p IntSlice) LessOrEqValueIndex(v Any, i int) bool { return p[i] >= v.(AnyInt).value }

// Clone with size
func (p IntSlice) Clone(count int) SortInterface { return make(IntSlice, count) }

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
		for j := i; j < n-i; j++ {
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
		if data.Less(i-1, i) {
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

// Insertion sort O(n^2)
func InsertSort(data SortInterface) {
	var j int
	size := data.Len()
	for i := 1; i < size; i++ {
		key := data.Get(i)
		for j = i - 1; j >= 0 && data.LessValueIndex(key, j); j-- {
			data.SetByIndex(j+1, j)
		}
		data.Set(j+1, key)
	}
}

// Merge sort O(n log n)
func MergeSort(data SortInterface) {
	// if list size is 0 (empty) or 1, consider it sorted and return it
	// (using less than or equal prevents infinite recursion for a zero length m)
	if data.Len() <= 1 {
		return
	}
	// if list size is 0 (empty) or 1, consider it sorted and return it
	// (using less than or equal prevents infinite recursion for a zero length m)
	middle := data.Len() / 2
	left := data.Clone(middle)
	right := data.Clone(data.Len() - middle)
	var i int
	for i = 0; i < middle; i++ {
		left.Set(i, data.Get(i))
	}
	for i = middle; i < data.Len(); i++ {
		right.Set(i-middle, data.Get(i))
	}
	// recursively call merge_sort() to further split each sublist
	// until sublist size is 1
	MergeSort(left)
	MergeSort(right)
	// merge the sublists returned from prior calls to merge_sort()
	// and return the resulting merged sublist
	// 2. CONQUER Part...
	merge(left, right, data)
}

func merge(left, right, result SortInterface) {
	// receive the left and right sublist as arguments.
	// 'result' variable for the merged result of two sublists.
	leftIndex := 0
	rightIndex := 0
	index := 0
	// assign the element of the sublists to 'result' variable until there is no element to merge. 
	for leftIndex < left.Len() || rightIndex < right.Len() {
		if leftIndex < left.Len() && rightIndex < right.Len() {
			// compare the first two element, which is the small one, of each two sublists.
			if right.LessOrEqValueIndex(left.Get(leftIndex), rightIndex) {
				// the small element is copied to 'result' variable.
				// delete the copied one(a first element) in the sublist.
				result.Set(index, left.Get(leftIndex))
				leftIndex++
			} else {
				// same operation as the above(in the right sublist).
				result.Set(index, right.Get(rightIndex))
				rightIndex++
			}
		} else if leftIndex < left.Len() {
			// copy all of remaining elements from the sublist to 'result' variable, when there is no more element to compare with.
			result.Set(index, left.Get(leftIndex))
			leftIndex++
		} else {
			// same operation as the above(in the right sublist).
			result.Set(index, right.Get(rightIndex))
			rightIndex++
		}
		index++
	}
}

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
