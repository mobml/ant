package wizard

import (
	"github.com/mobml/ant/internal/models"
)

func AreaWizard(area *models.Area) error {
	fields := []Field{
		{
			Label: "Area name",
			Value: func() string {
				return area.Name
			},
			SetValue: func(v string) error {
				area.Name = v
				return nil
			},
			Optional: false,
		},
		{
			Label: "Plan ID",
			Value: func() string {
				return area.PlanID
			},
			SetValue: func(v string) error {
				area.PlanID = v
				return nil
			},
			Optional: false,
		},
		{
			Label: "Description",
			Value: func() string {
				return area.Description
			},
			SetValue: func(v string) error {
				area.Description = v
				return nil
			},
			Optional: true,
		},
	}

	return Run(fields)
}

func NeedsAreaWizard(area *models.Area) bool {
	return area.Name == "" || area.PlanID == ""
}
