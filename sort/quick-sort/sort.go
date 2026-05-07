package quickSort

import (
	"cmp"
	"math/rand/v2"
)

// SortInt sorts s in ascending order.
func SortInt(s []int) {
	if len(s) <= 1 {
		return
	}

	pivotIndex := partition(s)
	// Recursively sort the left and right partitions.
	SortInt(s[:pivotIndex])
	SortInt(s[pivotIndex+1:])
}

// SortString sorts s in alphabetical order.
func SortString(s []string) {
	if len(s) <= 1 {
		return
	}

	pivotIndex := partition(s)
	// Recursively sort the left and right partitions.
	SortString(s[:pivotIndex])
	SortString(s[pivotIndex+1:])
}

// SortBasicType sorts the slice s in ascending order.
// E must satisfy cmp.Ordered.
func SortBasicType[E cmp.Ordered](s []E) {
	if len(s) <= 1 {
		return
	}

	pivotIndex := partition(s)
	// Recursively sort the left and right partitions.
	SortBasicType(s[:pivotIndex])
	SortBasicType(s[pivotIndex+1:])
}

// randomPivot returns the index of a pivot element in s selected using the
// median-of-three method. It picks three random elements and returns the index
// of the median. This helps avoid worst-case behavior on already-sorted inputs.
func randomPivot[E cmp.Ordered](s []E) int {
	r1 := rand.IntN(len(s))
	r2 := rand.IntN(len(s))
	r3 := rand.IntN(len(s))

	if (s[r1] < s[r2]) != (s[r1] < s[r3]) {
		return r1
	}
	if (s[r2] < s[r1]) != (s[r2] < s[r3]) {
		return r2
	}
	return r3
}

// swap swaps the elements at indices i and j in the slice s.
func swap[E any](s []E, i, j int) {
	s[i], s[j] = s[j], s[i]
}

// partition partitions the slice s around a randomly chosen pivot.
// After partitioning, all elements less than the pivot are placed before it,
// and all others are placed after it. It returns the final index of the pivot.
func partition[E cmp.Ordered](s []E) int {
	pivotIndex := randomPivot(s)
	swap(s, pivotIndex, len(s)-1) // Move the pivot element to the end of the slice.
	pivot := s[len(s)-1]
	i := 0
	// Iterate through the slice and move elements less than the pivot to the left side of the slice.
	// After the loop, all elements less than the pivot will be on the left side of the slice,
	// and all elements greater than or equal to the pivot will be on the right side.
	for j := 0; j < len(s)-1; j++ {
		if cmp.Compare(s[j], pivot) < 0 {
			swap(s, i, j)
			i++
		}
	}

	swap(s, i, len(s)-1) // Move the pivot element to its final position.
	return i
}

// SortSlice sorts s using compareFunc.
// compareFunc should return a negative value if a < b, zero if they are equal,
// and a positive value if a > b.
// It returns ErrSliceCannotBeNil if s is nil, or ErrComparisonFunctionRequired if compareFunc is nil.
func SortSlice[S ~[]E, E any](s S, compareFunc func(E,E) int) error {
	if s == nil {
		return ErrSliceCannotBeNil
	}

	if compareFunc == nil {
		return ErrComparisonFunctionRequired
	}

	if len(s) <= 1 {
		return nil
	}

	pivotIndex := partitionSlice(s, compareFunc)
	// Recursively sort the left and right partitions.
	SortSlice(s[:pivotIndex],compareFunc)
	SortSlice(s[pivotIndex+1:], compareFunc)
	return nil
}

// randomPivotForSlice returns the index of a pivot element selected using the
// median-of-three method, comparing elements with compareFunc.
// compareFunc should return a negative number if a < b, zero if a == b,
// and a positive number if a > b.
func randomPivotForSlice[S ~[]E, E any] (s S, compareFunc func(E,E) int) int {
	r1 := rand.IntN(len(s))
	r2 := rand.IntN(len(s))
	r3 := rand.IntN(len(s))

	if (compareFunc(s[r1], s[r2]) < 0) != (compareFunc(s[r1], s[r3]) < 0) {
		return r1
	} else if (compareFunc(s[r2], s[r1]) < 0) != (compareFunc(s[r2], s[r3]) < 0) {
		return r2
	}

	return r3
}

// partitionSlice partitions the slice s around a randomly chosen pivot.
// After calling partitionSlice, all elements for which compareFunc(x, pivot) < 0
// are placed before the pivot, and all others are placed after it.
// It returns the final index of the pivot.
//
// partitionSlice is a generic version of partition; it uses compareFunc to
// determine element order. compareFunc should return a negative number if
// a < b, zero if a == b, and a positive number if a > b.
func partitionSlice[S ~[]E, E any](s S, compareFunc func(E,E) int) int {
	pivotIndex := randomPivotForSlice(s, compareFunc)
	swap(s, pivotIndex, len(s)-1)  // Move the pivot element to the end of the slice.
	pivot := s[len(s)-1]
	i := 0
	// Iterate through the slice and move elements less than the pivot to the left side of the slice.
	// After the loop, all elements less than the pivot will be on the left side of the slice,
	// and all elements greater than or equal to the pivot will be on the right side.
	for j := 0; j < len(s)-1; j++ {
		if compareFunc(s[j], pivot) < 0 {
			swap(s, i, j)
			i++
		}
	}

	swap(s, i, len(s)-1)
	return i
}