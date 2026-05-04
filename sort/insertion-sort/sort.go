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
	