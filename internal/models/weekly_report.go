package models

import (
	"time"
)

type WeeklyReport struct {
	ID          string
	PlanID      string
	WeekStart   time.Time
	WeekEnd     time.Time
	ReportMD    string
	GeneratedAt time.Time
}
