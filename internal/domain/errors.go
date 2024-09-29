package domain

import "errors"

var OrderSizeMustBeGreaterThanZeroError = errors.New("order size must be greater than 0")
var PackSizesNotFoundError = errors.New("no pack sizes available")
