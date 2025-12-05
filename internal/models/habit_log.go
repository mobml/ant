package models

import (
	"time"
)

type HabitLog struct {
	ID        string
	HabitID   string
	LogDate   time.Time
	Value     float64
	Note      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
