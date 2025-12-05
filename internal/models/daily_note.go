package models

import (
	"time"
)

type DailyNote struct {
	ID        string
	NoteDate  time.Time
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
}
