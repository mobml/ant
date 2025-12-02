package models

import (
	"github.com/google/uuid"
	"time"
)

type DailyNote struct {
	ID        uuid.UUID
	NoteDate  time.Time
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
}
