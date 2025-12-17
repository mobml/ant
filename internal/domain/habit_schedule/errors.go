package habitschedule

import "errors"

var (
	ErrInvalidDayOfWeek = errors.New("day of week must be between 0 (Sunday) and 6 (Saturday)")
)
