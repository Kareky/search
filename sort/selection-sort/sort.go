package selectionSort

import (
	"cmp"
)

// sorts an array of integers in increasing order using selection sort algorithm.
func SortIntArray(arrayToSort []int) {
	for i := 0; i < len(arrayToSort)-1; i++ {
		minIndex := findMinIndex(arrayToSort, i)
		swap(arrayToSort, i, minIndex)
	}
}

// sorts an array of strings in increasing order using selection sort algorithm.
func SortStringArray(arrayToSort []string) {
	for i := 0; i < len(arrayToSort)-1; i++ {
		minIndex := findMinIndex(arrayToSort, i)
		swap(arrayToSort, i, minIndex)
	}
}

// sorts an array of any basic type that can be ordered in increasing order using selection sort algorithm.
func SortBasicTypeArray[basicType cmp.Ordered](arrayToSort []basicType) {
	for i := 0; i < len(arrayToSort)-1; i++ {
		minIndex := findMinIndex(arrayToSort, i)
		swap(arrayToSort, i, minIndex)
	}
}

// swaps the elements at firstIndex and secondIndex in the array
func swap[basicType cmp.Ordered](arrayToSwap []basicType, firstIndex, secondIndex int) []basicType {
	if arrayToSwap == nil || firstIndex < 0 || secondIndex < 0 || firstIndex >= len(arrayToSwap) || secondIndex >= len(arrayToSwap) {
		return nil
	}
	var temp = arrayToSwap[firstIndex]
	arrayToSwap[firstIndex] = arrayToSwap[secondIndex]
	arrayToSwap[secondIndex] = temp
	return arrayToSwap
}

// finds the index of the smallest element in the array starting from startIndex
func findMinIndex[basicType cmp.Ordered](arrayToSearch []basicType, startIndex int) int {
    var minIndex = startIndex;

	//start from second element, no reason to compare first element with itself
    for i := minIndex + 1; i < len(arrayToSearch); i++ {
        if(arrayToSearch[i] < arrayToSearch[minIndex]) {
            minIndex = i
        }
    } 
    return minIndex;
};

// sorts a slice of any type using selection sort algorithm and a custom comparison function.
// The comparison function should return a negative integer if the first argument is less than the second,
// zero if they are equal, and a positive integer if the first argument is greater than the second.
// The function returns an error only if the slice is nil or if the comparison function is nil.
func SortSlice[s ~[]anyType, anyType any](sliceToSort *s, compareFunc func(anyType, anyType) int) (error) {
	if sliceToSort == nil {
		return ErrArrayCannotBeNil
	}

	if compareFunc == nil {
		return ErrComparisonFunctionRequired
	}

	for i := 0; i < len(*sliceToSort)-1; i++ {
		minIndex := findMinIndexSlice(*sliceToSort, i, compareFunc)
		// swapping the lower element at index minIndex with the current element at index i
		(*sliceToSort)[i], (*sliceToSort)[minIndex] = (*sliceToSort)[minIndex], (*sliceToSort)[i]
	}

	return nil
}

// finds the index of the smallest element in the slice starting from startIndex using the provided comparison function.
func findMinIndexSlice[s ~[]anyType, anyType any](sliceToSearch s, startIndex int, compareFunc func(anyType, anyType) int) int {
	var minIndex = startIndex;
	//start from second element, no reason to compare first element with itself
	for i := minIndex + 1; i < len(sliceToSearch); i++ {
		if compareFunc((sliceToSearch)[i], (sliceToSearch)[minIndex]) < 0 {
			minIndex = i
		}
	}
	return minIndex
}
