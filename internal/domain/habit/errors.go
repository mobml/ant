package habit

import "errors"

var (
	ErrInvalidMeasureType = errors.New("invalid measure type")
	ErrGoalIDRequired     = errors.New("goal ID is required")
)
