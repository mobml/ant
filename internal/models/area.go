package models

import (
	"github.com/google/uuid"
	"time"
)

type Area struct {
	ID          uuid.UUID
	PlanID      uuid.UUID
	Name        string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
