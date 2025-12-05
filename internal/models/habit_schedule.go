package models

import (
	"time"
)

type HabitSchedule struct {
	ID        string
	HabitID   string
	DayOfWeek int
	CreatedAt time.Time
	UpdatedAt time.Time
}
