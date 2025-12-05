package models

import (
	"time"
)

type Goal struct {
	ID          string
	AreaID      string
	Name        string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
