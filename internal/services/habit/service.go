package habit

import (
	"time"

	dc "github.com/mobml/ant/internal/domain/common"
	dh "github.com/mobml/ant/internal/domain/habit"
	"github.com/mobml/ant/internal/models"
	repo "github.com/mobml/ant/internal/repositories"
)

type HabitService interface {
	CreateHabit(habit *models.Habit) error
	CreateHabitWithSchedule(habit *models.Habit, days []int) error
	GetHabitByID(id string) (*models.Habit, error)
	HabitsForToday() ([]models.HabitWithStatus, error)
	ListHabits() ([]*models.Habit, error)
	UpdateHabit(habit *models.Habit) error
	UpdateHabitWithSchedule(habit *models.Habit, days []int) error
	DeleteHabit(id string) error
}

type habitService struct {
	habitRepo repo.HabitRepository
}

func NewHabitService(habitRepo repo.HabitRepository) HabitService {
	return &habitService{
		habitRepo: habitRepo,
	}
}

func (hs *habitService) CreateHabit(habit *models.Habit) error {
	if err := dh.ValidateHabit(habit); err != nil {
		return err
	}
	return hs.habitRepo.Create(habit)
}

func (hs *habitService) CreateHabitWithSchedule(habit *models.Habit, days []int) error {
	if err := dh.ValidateHabit(habit); err != nil {
		return err
	}
	return hs.habitRepo.CreateHabit(habit, days)
}

func (hs *habitService) GetHabitByID(id string) (*models.Habit, error) {
	if err := dc.ValidateID(id); err != nil {
		return nil, err
	}
	return hs.habitRepo.FindByID(id)
}

func (hs *habitService) HabitsForToday() ([]models.HabitWithStatus, error) {
	now := time.Now()

	day := int(now.Weekday())
	if day == 0 {
		day = 7
	}

	habits, err := hs.habitRepo.HabitsForToday(day)
	if err != nil {
		return nil, err
	}
	return habits, nil
}

func (hs *habitService) ListHabits() ([]*models.Habit, error) {
	return hs.habitRepo.List()
}

func (hs *habitService) UpdateHabit(habit *models.Habit) error {
	if err := dh.ValidateHabit(habit); err != nil {
		return err
	}

	if err := dc.ValidateID(habit.ID); err != nil {
		return err
	}

	return hs.habitRepo.Update(habit)
}

func (hs *habitService) DeleteHabit(id string) error {
	if err := dc.ValidateID(id); err != nil {
		return err
	}
	return hs.habitRepo.Delete(id)
}

func (hs *habitService) UpdateHabitWithSchedule(habit *models.Habit, days []int) error {
	if err := dh.ValidateHabit(habit); err != nil {
		return err
	}

	if err := dc.ValidateID(habit.ID); err != nil {
		return err
	}

	if err := hs.habitRepo.DeleteHabitSchedules(habit.ID); err != nil {
		return err
	}

	if err := hs.habitRepo.Update(habit); err != nil {
		return err
	}

	if err := hs.habitRepo.CreateHabitSchedule(habit.ID, days); err != nil {
		return err
	}

	return nil
}
