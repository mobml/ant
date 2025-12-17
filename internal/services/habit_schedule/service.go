package habitschedule

import (
	dc "github.com/mobml/ant/internal/domain/common"
	dhs "github.com/mobml/ant/internal/domain/habit_schedule"
	"github.com/mobml/ant/internal/models"
	repo "github.com/mobml/ant/internal/repositories"
)

type HabitScheduleService interface {
	CreateHabitSchedule(habitSchedule *models.HabitSchedule) error
	ListHabitSchedules() ([]*models.HabitSchedule, error)
	UpdateHabitSchedule(habitSchedule *models.HabitSchedule) error
	DeleteHabitSchedule(id string) error
}

type habitScheduleService struct {
	habitScheduleRepo repo.HabitScheduleRepository
}

func NewHabitScheduleService(habitScheduleRepo repo.HabitScheduleRepository) HabitScheduleService {
	return &habitScheduleService{
		habitScheduleRepo: habitScheduleRepo,
	}
}

func (hss *habitScheduleService) CreateHabitSchedule(habitSchedule *models.HabitSchedule) error {

	if err := dhs.ValidateHabitSchedule(habitSchedule); err != nil {
		return err
	}

	return hss.habitScheduleRepo.Create(habitSchedule)
}

func (hss *habitScheduleService) ListHabitSchedules() ([]*models.HabitSchedule, error) {
	return hss.habitScheduleRepo.List()
}

func (hss *habitScheduleService) UpdateHabitSchedule(habitSchedule *models.HabitSchedule) error {

	if err := dhs.ValidateHabitSchedule(habitSchedule); err != nil {
		return err
	}

	if err := dc.ValidateID(habitSchedule.ID); err != nil {
		return err
	}

	return hss.habitScheduleRepo.Update(habitSchedule)
}

func (hss *habitScheduleService) DeleteHabitSchedule(id string) error {
	if err := dc.ValidateID(id); err != nil {
		return err
	}

	return hss.habitScheduleRepo.Delete(id)
}
