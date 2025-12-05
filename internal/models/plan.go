package models

import (
	"time"
)

type Plan struct {
	ID          string
	Name        string
	Description string
	StartDate   time.Time
	Duration    int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
