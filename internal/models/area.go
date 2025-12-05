package models

import (
	"time"
)

type Area struct {
	ID          string
	PlanID      string
	Name        string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
