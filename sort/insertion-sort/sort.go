package insertionSort

import (
	"cmp"
	"github.com/Kareky/search/internal/errors"
)

// SortInt sorts s in ascending order using insertion sort.
func SortInt(s []int) {
	for i := 1; i < len(s); i++ {
		insert(s, i-1, s[i])
	}
}

// SortString sorts s in alphabetical order using insertion sort.
func SortString(s []string) {
	for i := 1; i < len(s); i++ {
		insert(s, i-1, s[i])
	}
}

// SortBasicType sorts s in ascending order using insertion sort.
// E must be a type that can be ordered using the <= operator, such as int, float64, string, etc.
func SortBasicType[E cmp.Ordered](s []E) {
	for i := 1; i < len(s); i++ {
		insert(s, i-1, s[i])
	}
}

// insert inserts value into s in sorted order, assuming s[:orderedToIndex+1] is already sorted.
func insert[E cmp.Ordered](s []E, orderedToIndex int, value E) {
    var i = orderedToIndex
    for(i>=0 && s[i]>value){
        s[i+1] = s[i]
		i--
    }
    s[i+1] = value
}

// SortSlice sorts s using insertion sort and the provided compareFunc.
// compareFunc should return a negative integer if a < b, zero if a == b, and a positive integer if a > b.
// It returns ErrSliceCannotBeNil if s is nil, or ErrComparisonFunctionRequired if compareFunc is nil.
func SortSlice[S ~[]E, E any](s S, compareFunc func(E, E) int) error {
	if s == nil {
		return errors.ErrSliceCannotBeNil
	}

	if compareFunc == nil {
		return errors.ErrComparisonFunctionRequired
	}

	// Insertion sort algorithm
	for i := 1; i < len(s); i++ {
		insertToSlice(s, i-1, s[i], compareFunc)
	}
	return nil
}

// insertToSlice inserts value into sin sorted order, assuming s[:orderedToIndex+1] is already sorted, using compareFunc for comparisons.
func insertToSlice[S ~[]E, E any](s S, orderedToIndex int, value E, compareFunc func(E, E) int) {
    var i = orderedToIndex
    for(i>=0 && compareFunc(s[i], value) > 0){
        s[i+1] = s[i]
		i--
    }
    s[i+1] = value
}