package models

import (
	"github.com/google/uuid"
	"time"
)

type Plan struct {
	ID          uuid.UUID
	Name        string
	Description string
	StartDate   time.Time
	Duration    int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
