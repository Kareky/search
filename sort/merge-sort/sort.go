package mergeSort

import "cmp"

// SortInt sorts s in ascending order using merge sort.
func SortInt(s []int) []int {
	if len(s) <= 1 {
		return s
	}
	divideAndOrder(s, 0, len(s)-1)
	return s
}

// SortString sorts s in alphabetical order using merge sort.
func SortString(s []string) []string {
	if len(s) <= 1 {
		return s
	}
	divideAndOrder(s, 0, len(s)-1)
	return s
}

// SortBasicType sorts s in ascending order using merge sort.
// E must be a type that can be ordered using the <= operator, such as int, float64, string, etc.
func SortBasicType[E cmp.Ordered](s []E) []E {
	if len(s) <= 1 {
		return s
	}
	divideAndOrder(s, 0, len(s)-1)
	return s
}

// divideAndOrder recursively sorts s[left:right+1] using merge sort.
func divideAndOrder[E cmp.Ordered](s []E, left int, right int) {
	if left >= right {
		return
	}

	mid := int(uint(left+right) >> 1)
	// recursively divide the slice into two halves until
	// we reach slices of length 1 or 0 which are already sorted
	divideAndOrder(s, left, mid)
	divideAndOrder(s, mid+1, right)
	merge(s, left, mid, right)
}

/// merge merges the sorted halves s[left:middle] and s[middle+1:right] into a single sorted subslice.
func merge[E cmp.Ordered](s []E, left int, middle int, right int) {
	// split the slice into two halves since the two halves
	// were sorted when previously calling divideAndOrder() and merge()
	leftSlice := make([]E, middle-left+1)
	rightSlice := make([]E, right-middle)
	copy(leftSlice, s[left:middle+1])
	copy(rightSlice, s[middle+1:right+1])

	i, j, k := 0, 0, left
	// Merge the two halves checking which element is smaller
	// and copying it to the original slice
	for i < len(leftSlice) && j < len(rightSlice) {
		if leftSlice[i] <= rightSlice[j] {
			s[k] = leftSlice[i]
			i++
		} else {
			s[k] = rightSlice[j]
			j++
		}
		k++
	}

	// Since the merge stop when one of the two halves is fully copied
	// we need to copy the remaining elements of the other half
	for i < len(leftSlice) {
		s[k] = leftSlice[i]
		i++
		k++
	}

	for j < len(rightSlice) {
		s[k] = rightSlice[j]
		j++
		k++
	}
}

// SortSlice sorts s using merge sort and compareFunc.
// compareFunc should return a negative value if a < b, zero if they are equal,
// and a positive value if a > b.
// It returns ErrSliceCannotBeNil if s is nil, or ErrComparisonFunctionRequired if compareFunc is nil.
func SortSlice[s ~[]E, E any](sliceToSort s, compareFunc func(E, E) int) error {
	if sliceToSort == nil {
		return ErrSliceCannotBeNil
	}

	if compareFunc == nil {
		return ErrComparisonFunctionRequired
	}

	divideAndOrderSlice(sliceToSort, 0, len(sliceToSort)-1, compareFunc)
	return nil
}

// divideAndOrderSlice recursively sorts s[left:right+1] using merge sort and compareFunc.
func divideAndOrderSlice[s ~[]E, E any](sliceToSort s, left int, right int, compareFunc func(E, E) int) {
	if left >= right {
		return
	}

	mid := int(uint(left+right) >> 1)

	// recursively divide the slice into two halves until
	// we reach slices of length 1 or 0 which are already sorted
	divideAndOrderSlice(sliceToSort, left, mid, compareFunc)
	divideAndOrderSlice(sliceToSort, mid+1, right, compareFunc)
	mergeSlice(sliceToSort, left, mid, right, compareFunc)
}

// mergeSlice merges the sorted halves s[left:middle] and s[middle+1:right] into a single sorted subslice
// using compareFunc to determine the order.
func mergeSlice[s ~[]E, E any](sliceToSort s, left int, middle int, right int, compareFunc func(E, E) int) {
	// split the slice into two halves since the two halves
	// were sorted when previously calling divideAndOrder() and merge()
	leftSlice := make([]E, middle-left+1)
	rightSlice := make([]E, right-middle)
	copy(leftSlice, sliceToSort[left:middle+1])
	copy(rightSlice, sliceToSort[middle+1:right+1])

	i, j, k := 0, 0, left
	// Merge the two halves checking which element is smaller
	// and copying it to the original slice
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
