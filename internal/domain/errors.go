package domain

import "errors"

var (
	OrderSizeMustBeGreaterThanZeroError = errors.New("order size must be greater than 0")
	PackSizesNotFoundError              = errors.New("no pack sizes available")
)

var (
	EmptyPackSizesError  = errors.New("pack sizes cannot be empty")
	InvalidPackSizeError = errors.New("invalid pack size")
)
