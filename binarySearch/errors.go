package binarySearch

import "fmt"

var ErrUnmatchedTargetAndArrayTypes = fmt.Errorf("the type of the target does not match the type of the elements in the array")
var ErrComparisonFunctionRequired = fmt.Errorf("a comparison function is required for searching non-basic types")