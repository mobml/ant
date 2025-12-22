package wizard

import (
	"strconv"
	"time"

	"github.com/mobml/ant/internal/models"
)

func parseDate(value string) (time.Time, error) {
	return time.Parse("2006-01-02", value)
}

func PlanWizard(plan *models.Plan) error {
	fields := []Field{
		{
			Label: "Plan name",
			Value: func() string {
				return plan.Name
			},
			SetValue: func(v string) error {
				plan.Name = v
				return nil
			},
			Optional: false,
		},
		{
			Label: "Start date (YYYY-MM-DD)",
			Value: func() string {
				if plan.StartDate.IsZero() {
					return ""
				}
				return plan.StartDate.Format("2006-01-02")
			},
			SetValue: func(v string) error {
				d, err := parseDate(v)
				if err != nil {
					return err
				}
				plan.StartDate = d
				return nil
			},
			Optional: false,
		},
		{
			Label: "Duration (weeks)",
			Value: func() string {
				if plan.Duration == 0 {
					return ""
				}
				return strconv.Itoa(plan.Duration)
			},
			SetValue: func(v string) error {
				d, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				plan.Duration = d
				return nil
			},
			Optional: false,
		},
		{
			Label: "Description",
			Value: func() string {
				return plan.Description
			},
			SetValue: func(v string) error {
				plan.Description = v
				return nil
			},
			Optional: true,
		},
	}

	return Run(fields)
}

func NeedsPlanWizard(p *models.Plan) bool {
	return p.Name == "" ||
		p.Duration == 0 ||
		p.StartDate.IsZero()
}
