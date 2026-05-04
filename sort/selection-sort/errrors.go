package selectionSort

import (
	"fmt"
)

var ErrArrayCannotBeNil = fmt.Errorf("the array to sort cannot be nil")
var ErrComparisonFunctionRequired = fmt.Errorf("a comparison function is required for sorting non-basic types")