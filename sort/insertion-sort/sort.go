package insertionSort

import (
	"cmp"
)

// sorts an array of integers in increasing order using insertion sort algorithm.
func SortIntArray(arrayToSort []int) {
	for i := 1; i < len(arrayToSort); i++ {
		insert(arrayToSort, i-1, arrayToSort[i])
	}
}

// sorts an array of strings in increasing order using insertion sort algorithm.
func SortStringArray(arrayToSort []string) {
	for i := 1; i < len(arrayToSort); i++ {
		insert(arrayToSort, i-1, arrayToSort[i])
	}
}

// sorts an array of any basic type that can be ordered in increasing order using insertion sort algorithm.
func SortBasicTypeArray[basicType cmp.Ordered](arrayToSort []basicType) {
	for i := 1; i < len(arrayToSort); i++ {
		insert(arrayToSort, i-1, arrayToSort[i])
	}
}

// insert the value passed as parameter in the right position in the array
// assuming that the array is sorted until orderedToIndex
func insert[basicType cmp.Ordered](arrayForInsertion []basicType, orderedToIndex int, value basicType) {
    var i = orderedToIndex
    for(i>=0 && arrayForInsertion[i]>value){
        arrayForInsertion[i+1] = arrayForInsertion[i]
		i--
    }
    arrayForInsertion[i+1] = value
}

// sorts a slice of any type using insertion sort algorithm and a custom comparison function.
// The comparison function should return a negative integer if the first argument is less than the second,
// zero if they are equal, and a positive integer if the first argument is greater than the second.
// The function returns an error only if the slice is nil or if the comparison function is nil.
func SortSlice[s ~[]anyType, anyType any](sliceToSort *s, compareFunc func(anyType, anyType) int) (error) {
	if sliceToSort == nil {
		return ErrSliceCannotBeNil
	}

	if compareFunc == nil {
		return ErrComparisonFunctionRequired
	}

	// Insertion sort algorithm
	for i := 1; i < len(*sliceToSort); i++ {
		insertToSlice(sliceToSort, i-1, (*sliceToSort)[i], compareFunc)
	}
	return nil
}


// insert the value passed as parameter in the right position in the slice
// assuming that the slice is sorted until orderedToIndex
// it uses a custom comparison function to compare the values
func insertToSlice[s ~[]anyType, anyType any](sliceForInsertion *s, orderedToIndex int, value anyType, compareFunc func(anyType, anyType) int) {
    var i = orderedToIndex
    for(i>=0 && compareFunc((*sliceForInsertion)[i], value) > 0){
        (*sliceForInsertion)[i+1] = (*sliceForInsertion)[i]
		i--
    }
    (*sliceForInsertion)[i+1] = value
}