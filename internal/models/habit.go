package models

import (
	"time"
)

type Habit struct {
	ID          string
	GoalID      string
	Name        string
	Description string
	MeasureType MeasureType
	MeasureUnit string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type HabitWithStatus struct {
	ID          string
	Name        string
	Description string
	WorkedToday bool
}
