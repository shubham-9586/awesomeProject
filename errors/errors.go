package errors

import "errors"

var (
	INVALID_AGE_ERROR = errors.New("User must be an adult")
)
