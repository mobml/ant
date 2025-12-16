package habit

import (
	dc "github.com/mobml/ant/internal/domain/common"
	dh "github.com/mobml/ant/internal/domain/habit"
	"github.com/mobml/ant/internal/models"
	repo "github.com/mobml/ant/internal/repositories"
)

type HabitService interface {
	CreateHabit(habit *models.Habit) error
	ListHabits() ([]*models.Habit, error)
	UpdateHabit(habit *models.Habit) error
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
