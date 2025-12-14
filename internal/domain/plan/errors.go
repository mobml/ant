package plan

import "errors"

var (
	ErrDurationInvalid  = errors.New("duration must be greater than 0")
	ErrStartDateInvalid = errors.New("start date is required")
	ErrStartDatePast    = errors.New("start date cannot be in the past")
)
