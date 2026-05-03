package selectionsort

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