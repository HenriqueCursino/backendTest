package exceptions

import "errors"

var (
	INVALID_BALANCE error = errors.New("invalid balance")
)
