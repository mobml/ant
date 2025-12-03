package models

import (
	"github.com/google/uuid"
	"time"
)

type Goal struct {
	ID          uuid.UUID
	AreaID      uuid.UUID
	Name        string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
