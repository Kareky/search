package binarySearch

import (
	"cmp"
	"github.com/Kareky/search/internal/errors"
)

// SearchInt returns the index and true if target is in s, or -1 and false.
// s must be sorted in ascending order.
func SearchInt(s []int, target int) (int, bool) {
	left, right := 0, len(s)-1
	for left <= right {
		mid := int(uint(left+right) >> 1) // avoid left+right overflow max int
		if s[mid] == target {
			return mid, true
		}
		if s[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1, false
}

// SearchString returns the index and true if target is in s, or -1 and false.
// s must be sorted in alphabetical order.
func SearchString(s []string, target string) (int, bool) {
	left, right := 0, len(s)-1
	for left <= right {
		mid := int(uint(left+right) >> 1) // avoid left+right overflow max int
		if s[mid] == target {
			return mid, true
		}
		if s[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1, false
}

// SearchBasicType returns the index and true if target is in s, or -1 and false.
// s must be sorted in ascending order.
func SearchBasicType[basicType cmp.Ordered](s []basicType, target basicType) (int, bool) {
	left, right := 0, len(s)-1
	for left <= right {
		mid := int(uint(left+right) >> 1) // avoid left+right overflow max int
		if cmp.Compare(s[mid], target) < 0 {
			left = mid + 1
		} else if cmp.Compare(s[mid], target) > 0 {
			right = mid - 1
		} else if cmp.Compare(s[mid], target) == 0 {
			return mid, true
		}
	}
	return -1, false
}

// SearchSlice returns the index and nil error if target is found, or -1 and nil error if not.
// It returns -1 and ErrComparisonFunctionRequired if compareFunc is nil.
// s must be sorted in ascending order according to compareFunc.
func SearchSlice[s ~[]anyType, anyType any](sliceToSearch s, target anyType, compareFunc func(anyType, anyType) int) (int, error) {
	if sliceToSearch == nil {
		return -1, errors.ErrSliceCannotBeNil
	}

	if compareFunc == nil {
		return -1, errors.ErrComparisonFunctionRequired
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
