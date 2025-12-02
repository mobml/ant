package models

import (
	"github.com/google/uuid"
	"time"
)

type WeeklyReport struct {
	ID          uuid.UUID
	PlanID      string
	WeekStart   time.Time
	WeekEnd     time.Time
	ReportMD    string
	GeneratedAt time.Time
}
