package binarySearch

import (
	"cmp"
)

// performs a binary search for the target integer in the sorted array of integers.
// The array must be sorted in increasing order.
// It returns the index of the target if found, or -1 and false if not found.
func SearchInt(arrayToSearch []int, target int) (int, bool) {
	left, right := 0, len(arrayToSearch)-1
	for left <= right {
		mid := int(uint(left+right) >> 1) // avoid left+right overflow max int
		if arrayToSearch[mid] == target {
			return mid, true
		}
		if arrayToSearch[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1, false
}

// performs a binary search for the target string in the sorted array of strings.
// The array must be sorted in alphabetical order.
// It returns the index of the target if found, or -1 and false if not found.
func SearchString(arrayToSearch []string, target string) (int, bool) {
	left, right := 0, len(arrayToSearch)-1
	for left <= right {
		mid := int(uint(left+right) >> 1) // avoid left+right overflow max int
		if arrayToSearch[mid] == target {
			return mid, true
		}
		if arrayToSearch[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1, false
}

// performs a binary search for the target of any basic type
// that can be ordered in the sorted array of the same type.
// The array must be sorted in increasing order.
// It returns the index of the target if found, or -1 and false if not found.
func SearchBasicType[basicType cmp.Ordered](arrayToSearch []basicType, target basicType) (int, bool) {
	left, right := 0, len(arrayToSearch)-1
	for left <= right {
		mid := int(uint(left+right) >> 1) // avoid left+right overflow max int
		if cmp.Compare(arrayToSearch[mid], target) < 0 {
			left = mid + 1
		} else if cmp.Compare(arrayToSearch[mid], target) > 0 {
			right = mid - 1
		} else if cmp.Compare(arrayToSearch[mid], target) == 0 {
			return mid, true
		}
	}
	return -1, false
}

// performs a binary search for the target in the sorted slice of any type.
// When returning an error, the index is set to -1. When the target is not found, the error is nil and the index is set to -1.
// The slice must be sorted  in increasing order according to the comparison function if provided. 
// compareFunc should return 0 if the slice element matches the target,
// a negative number if the slice element precedes the target,
// or a positive number if the slice element follows the target.
// The function use the same behavior as slices.BinarySearch
func SearchSlice[s ~[]anyType, anyType any](sliceToSearch s, target anyType, compareFunc func(anyType, anyType) int) (int, error) {
	if compareFunc == nil {
		return -1, ErrComparisonFunctionRequired
	}

	left, right := 0, len(sliceToSearch)-1
	for left <= right {
		mid := int(uint(left+right) >> 1) // avoid left+right overflow max int
		if compareFunc(sliceToSearch[mid], target) < 0 {
			left = mid + 1
		} else if compareFunc(sliceToSearch[mid], target) > 0 {
			right = mid - 1
		} else if compareFunc(sliceToSearch[mid], target) == 0 {
			return mid, nil
		}
	}

	return -1, nil
}

