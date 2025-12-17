package dailynote

import "errors"

var (
	ErrContentRequired  = errors.New("content is required")
	ErrNoteDateRequired = errors.New("note date is required")
)
