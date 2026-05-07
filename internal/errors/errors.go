package errors

import "errors"

var (
    ErrSliceCannotBeNil           = errors.New("slice cannot be nil")
    ErrComparisonFunctionRequired = errors.New("comparison function must not be nil")
)