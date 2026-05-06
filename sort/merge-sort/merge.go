package mergeSort

import "cmp"

// Sorts the given array of integers using the merge sort algorithm and returns the sorted array.
func SortInt(arrayToSort []int) []int {
	if len(arrayToSort) <= 1 {
		return arrayToSort
	}
	divideAndOrder(arrayToSort, 0, len(arrayToSort)-1)
	return arrayToSort
}

// Sorts the given array of strings using the merge sort algorithm and returns the sorted array.
func SortString(arrayToSort []string) []string {
	if len(arrayToSort) <= 1 {
		return arrayToSort
	}
	divideAndOrder(arrayToSort, 0, len(arrayToSort)-1)
	return arrayToSort
}

// Sorts the given array of any basic type that can be ordered using the merge sort algorithm and returns the sorted array.
func SortBasicType[basicType cmp.Ordered](arrayToSort []basicType) []basicType {
	if len(arrayToSort) <= 1 {
		return arrayToSort
	}
	divideAndOrder(arrayToSort, 0, len(arrayToSort)-1)
	return arrayToSort
}

// divide the given array into two halves and sort them recursively, then merge the sorted halves
func divideAndOrder[anyType cmp.Ordered](arrayToSort []anyType, left int, right int) {
	if len(arrayToSort) <= 1 {
		return
	}

	mid := int(uint(left+right) >> 1)
	// recursively divide the array into two halves until
	// we reach arrays of length 1 or 0 which are already sorted
	divideAndOrder(arrayToSort, left, mid)
	divideAndOrder(arrayToSort, mid+1, right)
	merge(arrayToSort, left, mid, right)
}

// merge two sorted halves of the array into a single sorted array
func merge[anyType cmp.Ordered](arrayToSort []anyType, left int, middle int, right int) {
	// split the array into two halves since the two halves
	// were sorted when previously calling divideAndOrder() and merge()
	leftArray := make([]anyType, middle-left+1)
	rightArray := make([]anyType, right-middle)
	copy(leftArray, arrayToSort[left:middle+1])
	copy(rightArray, arrayToSort[middle+1:right+1])

	i, j, k := 0, 0, left
	// Merge the two halves checking which element is smaller
	// and copying it to the original array
	for i < len(leftArray) && j < len(rightArray) {
		if leftArray[i] <= rightArray[j] {
			arrayToSort[k] = leftArray[i]
			i++
		} else {
			arrayToSort[k] = rightArray[j]
			j++
		}
		k++
	}

	// Since the merge stop when one of the two halves is fully copied
	// we need to copy the remaining elements of the other half
	for i < len(leftArray) {
		arrayToSort[k] = leftArray[i]
		i++
		k++
	}

	for j < len(rightArray) {
		arrayToSort[k] = rightArray[j]
		j++
		k++
	}
}

// Sorts the given slice of any type using the merge sort algorithm and a custom comparison function, and returns the sorted slice.
// The comparison function should return a negative integer if the first argument is less than the second,
// zero if they are equal, and a positive integer if the first argument is greater than the second.
// The function returns an error only if the slice is nil or if the comparison function is nil.
func SortSlice[s ~[]anyType, anyType any](sliceToSort s, compareFunc func(anyType, anyType) int) error {
	if sliceToSort == nil {
		return ErrSliceCannotBeNil
	}

	if compareFunc == nil {
		return ErrComparisonFunctionRequired
	}

	divideAndOrderSlice(sliceToSort, 0, len(sliceToSort)-1, compareFunc)
	return nil
}

// divide the given slice into two halves and sort them recursively, then merge the sorted halves
func divideAndOrderSlice[s ~[]anyType, anyType any](sliceToSort s, left int, right int, compareFunc func(anyType, anyType) int) {
	if len(sliceToSort) <= 1 {
		return
	}

	mid := int(uint(left+right) >> 1)

	// recursively divide the array into two halves until
	// we reach arrays of length 1 or 0 which are already sorted
	divideAndOrderSlice(sliceToSort, left, mid, compareFunc)
	divideAndOrderSlice(sliceToSort, mid+1, right, compareFunc)
	mergeSlice(sliceToSort, left, mid, right, compareFunc)
}

// merge two sorted halves of the slice into a single sorted slice
func mergeSlice[s ~[]anyType, anyType any](sliceToSort s, left int, middle int, right int, compareFunc func(anyType, anyType) int) {
	// split the array into two halves since the two halves
	// were sorted when previously calling divideAndOrder() and merge()
	leftSlice := make([]anyType, middle-left+1)
	rightSlice := make([]anyType, right-middle)
	copy(leftSlice, sliceToSort[left:middle+1])
	copy(rightSlice, sliceToSort[middle+1:right+1])

	i, j, k := 0, 0, left
	// Merge the two halves checking which element is smaller
	// and copying it to the original array
	for i < len(leftSlice) && j < len(rightSlice) {
		if compareFunc(leftSlice[i], rightSlice[j]) <= 0 {
			sliceToSort[k] = leftSlice[i]
			i++
		} else {
			sliceToSort[k] = rightSlice[j]
			j++
		}
		k++
	}

	// Since the merge stop when one of the two halves is fully copied
	// we need to copy the remaining elements of the other half
	for i < len(leftSlice) {
		sliceToSort[k] = leftSlice[i]
		i++
		k++
	}

	for j < len(rightSlice) {
		sliceToSort[k] = rightSlice[j]
		j++
		k++
	}
}
