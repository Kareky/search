package binarySearch

import "fmt"

var ErrSliceCannotBeNil = fmt.Errorf("the slice to search cannot be nil")
var ErrComparisonFunctionRequired = fmt.Errorf("a comparison function is required for searching non-basic types")