package habitlog

import "errors"

var (
	ErrLogDateRequired = errors.New("log date is required")
	ErrValueRequired   = errors.New("value is required")
)
