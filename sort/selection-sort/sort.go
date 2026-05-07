package selectionSort

import (
	"cmp"
	"github.com/Kareky/search/internal/errors"
)

// SortInt sorts s in ascending order using selection sort.
func SortInt(s []int) {
	for i := 0; i < len(s)-1; i++ {
		minIndex := findMinIndex(s, i)
		swap(s, i, minIndex)
	}
}

// SortString sorts s in alphabetical order using selection sort.
func SortString(s []string) {
	for i := 0; i < len(s)-1; i++ {
		minIndex := findMinIndex(s, i)
		swap(s, i, minIndex)
	}
}

// SortBasicType sorts s in ascending order using selection sort.
// E must be a type that can be ordered using the <= operator, such as int, float64, string, etc.
func SortBasicType[E cmp.Ordered](s []E) {
	for i := 0; i < len(s)-1; i++ {
		minIndex := findMinIndex(s, i)
		swap(s, i, minIndex)
	}
}

// swap swaps s[firstIndex] and s[secondIndex] and returns s.
// It returns nil if s is nil or the indices are out of bounds.
func swap[E cmp.Ordered](s []E, firstIndex, secondIndex int) []E {
	if s == nil || firstIndex < 0 || secondIndex < 0 || firstIndex >= len(s) || secondIndex >= len(s) {
		return nil
	}
	var temp = s[firstIndex]
	s[firstIndex] = s[secondIndex]
	s[secondIndex] = temp
	return s
}

// findMinIndex returns the index of the smallest element in s[startIndex:].
func findMinIndex[E cmp.Ordered](s []E, startIndex int) int {
    var minIndex = startIndex;

	//start from second element, no reason to compare first element with itself
    for i := minIndex + 1; i < len(s); i++ {
        if(s[i] < s[minIndex]) {
            minIndex = i
        }
    } 
    return minIndex;
};

// SortSlice sorts s using selection sort and compareFunc.
// compareFunc should return a negative value if a < b, zero if they are equal,
// and a positive value if a > b.
// It returns an error if s or compareFunc is nil.
func SortSlice[s ~[]E, E any](sliceToSort s, compareFunc func(E, E) int) (error) {
	if sliceToSort == nil {
		return errors.ErrSliceCannotBeNil
	}

	if compareFunc == nil {
		return errors.ErrComparisonFunctionRequired
	}

	for i := 0; i < len(sliceToSort)-1; i++ {
		minIndex := findMinIndexSlice(sliceToSort, i, compareFunc)
		// swapping the lower element at index minIndex with the current element at index i
		sliceToSort[i], sliceToSort[minIndex] = sliceToSort[minIndex], sliceToSort[i]
	}

	return nil
}

// findMinIndexSlice returns the index of the smallest element in s[startIndex:],
// using compareFunc for comparisons.
func findMinIndexSlice[s ~[]E, E any](sliceToSearch s, startIndex int, compareFunc func(E, E) int) int {
	var minIndex = startIndex;
	//start from second element, no reason to compare first element with itself
	for i := minIndex + 1; i < len(sliceToSearch); i++ {
		if compareFunc((sliceToSearch)[i], (sliceToSearch)[minIndex]) < 0 {
			minIndex = i
		}
	}
	return minIndex
}
