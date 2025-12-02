package models

import (
	"github.com/google/uuid"
	"time"
)

type HabitLog struct {
	ID        uuid.UUID
	HabitID   uuid.UUID
	LogDate   time.Time
	Value     float64
	Note      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
