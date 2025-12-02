package models

import (
	"github.com/google/uuid"
	"time"
)

type Habit struct {
	ID          uuid.UUID
	GoalID      uuid.UUID
	Name        string
	Description string
	MeasureType MeasureType
	MeasureUnit string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
