package quickSort

import (
	"fmt"
)

var ErrSliceCannotBeNil = fmt.Errorf("the slice to sort cannot be nil")
var ErrComparisonFunctionRequired = fmt.Errorf("a comparison function is required for sorting non-basic types")