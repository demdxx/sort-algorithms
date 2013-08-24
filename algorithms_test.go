/**
 * Sorting algorithms
 * @autor Dmitry Ponomarev <demdxx@gmail.com>
 * @year 2013
 */

package algorithms

import (
	"math/rand"
	"sort"
	"testing"
	"time"
)

const randLength = 10000

var array sort.IntSlice = nil

func randIntArray(l int) sort.IntSlice {
	a := make(sort.IntSlice, l)
	for i := 0; i < l; i++ {
		a[i] = rand.Intn(l)
	}
	return a
}

func getRandIntArray() sort.IntSlice {
	a := make(sort.IntSlice, len(array))
	copy(a, array)
	return a
}

func getRandIntArray2() IntSlice {
	a := make(IntSlice, len(array))
	copy(a, array)
	return a
}

func TestInit(t *testing.T) {
	if nil == array {
		rand.Seed(time.Now().UTC().UnixNano())
		array = randIntArray(randLength)
	}
}

// Google quickSort O(n*log(n))
func TestStandartGoogleSort(t *testing.T) {
	arrayInts := getRandIntArray()
	sort.Sort(arrayInts)
	if randLength <= 100 {
		t.Log(arrayInts)
	}
}

// Stable algorithms

// Selection sort O(n^2)
func TestSelectionSort(t *testing.T) {
	arrayInts := getRandIntArray()
	SelectionSort(arrayInts)
	if randLength <= 100 {
		t.Log(arrayInts)
	}
}

// Bubble sort O(n^2)
func TestBubbleSort(t *testing.T) {
	arrayInts := getRandIntArray()
	BubbleSort(arrayInts)
	if randLength <= 100 {
		t.Log(arrayInts)
	}
}

// Combined Selection & Bubble sort
func TestSelectionBubbleSort(t *testing.T) {
	arrayInts := getRandIntArray()
	SelectionBubbleSort(arrayInts)
	if randLength <= 100 {
		t.Log(arrayInts)
	}
}

// Cocktail sort, bidirectional bubble sort O(n^2)
func TestShakerSort(t *testing.T) {
	arrayInts := getRandIntArray()
	ShakerSort(arrayInts)
	if randLength <= 100 {
		t.Log(arrayInts)
	}
}

// Dwarf sorting (Gnome Sort) O(n^2)
func TestGnomeSort(t *testing.T) {
	arrayInts := getRandIntArray()
	GnomeSort(arrayInts)
	if randLength <= 100 {
		t.Log(arrayInts)
	}
}

// Insertion sort O(n^2)
func TestInsertSort(t *testing.T) {
	arrayInts := getRandIntArray2()
	InsertSort(arrayInts)
	if randLength <= 100 {
		t.Log(arrayInts)
	}
}

// Merge sort O(n log n)
func TestMergeSort(t *testing.T) {
	arrayInts := getRandIntArray2()
	MergeSort(arrayInts)
	if randLength <= 100 {
		t.Log(arrayInts)
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
