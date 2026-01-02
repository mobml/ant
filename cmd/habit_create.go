package cmd

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/mobml/ant/internal/cli/wizard"
	"github.com/mobml/ant/internal/models"
	"github.com/spf13/cobra"
)

var habitCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new habit",
	RunE:  runHabitCreateCmd,
}

//create a function to create a habit, with wizard support, remember that habits have a schedule table

func runHabitCreateCmd(cmd *cobra.Command, args []string) error {
	habit := &models.Habit{}
	var daysRaw string

	if err := bindHabitFlags(cmd, habit, &daysRaw); err != nil {
		return err
	}

	if wizard.NeedsHabitWizard(habit, &daysRaw) {
		cmd.Println("Some required fields are missing. Launching interactive wizard...")

		if err := wizard.HabitWizard(habit, &daysRaw); err != nil {
			return err
		}
	}

	days, err := parseDays(daysRaw)
	if err != nil {
		return err
	}

	return HabitService.CreateHabitWithSchedule(habit, days)
}

func bindHabitFlags(cmd *cobra.Command, habit *models.Habit, days *string) error {
	var err error

	if habit.Name, err = cmd.Flags().GetString("name"); err != nil {
		return err
	}

	if habit.Description, err = cmd.Flags().GetString("description"); err != nil {
		return err
	}

	if habit.GoalID, err = cmd.Flags().GetString("goal-id"); err != nil {
		return err
	}

	measureTypeStr, err := cmd.Flags().GetString("measure-type")
	if err != nil {
		return err
	}
	habit.MeasureType = models.MeasureType(measureTypeStr)

	if *days, err = cmd.Flags().GetString("days"); err != nil {
		return err
	}

	return nil
}

func parseDays(input string) ([]int, error) {
	if input == "" {
		return nil, fmt.Errorf("days is required")
	}

	seen := make(map[int]bool)
	var result []int

	parts := strings.Split(input, ",")

	for _, part := range parts {
		if strings.Contains(part, "-") {
			rangeParts := strings.Split(part, "-")
			if len(rangeParts) != 2 {
				return nil, fmt.Errorf("invalid range: %s", part)
			}

			start, err := strconv.Atoi(rangeParts[0])
			if err != nil {
				return nil, err
			}
			end, err := strconv.Atoi(rangeParts[1])
			if err != nil {
				return nil, err
			}

			if start > end {
				return nil, fmt.Errorf("invalid range: %s", part)
			}

			for d := start; d <= end; d++ {
				if d < 1 || d > 7 {
					return nil, fmt.Errorf("day out of range: %d", d)
				}
				if !seen[d] {
					seen[d] = true
					result = append(result, d)
				}
			}
		} else {
			d, err := strconv.Atoi(part)
			if err != nil {
				return nil, err
			}
			if d < 1 || d > 7 {
				return nil, fmt.Errorf("day out of range: %d", d)
			}
			if !seen[d] {
				seen[d] = true
				result = append(result, d)
			}
		}
	}

	sort.Ints(result)
	return result, nil
}

func init() {
	habitCmd.AddCommand(habitCreateCmd)
	habitCreateCmd.Flags().StringP("name", "n", "", "Name of the habit")
	habitCreateCmd.Flags().StringP("description", "d", "", "Description of the habit (optional)")
	habitCreateCmd.Flags().StringP("goal-id", "g", "", "ID of the associated goal")
	habitCreateCmd.Flags().StringP("measure-type", "m", "integer", "Measure type (integer, float, boolean)")
	habitCreateCmd.Flags().StringP("days", "", "", "Days of the week for the habit schedule (e.g., 1,3,5 or 1-5)")
}
