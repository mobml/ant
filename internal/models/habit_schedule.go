package models

import (
	"github.com/google/uuid"
	"time"
)

type HabitSchedule struct {
	ID        uuid.UUID
	HabitID   uuid.UUID
	DayOfWeek int
	CreatedAt time.Time
	UpdatedAt time.Time
}
