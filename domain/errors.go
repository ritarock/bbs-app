package domain

import "errors"

var (
	ErrNotFound = errors.New("your requested item is not found")
)
