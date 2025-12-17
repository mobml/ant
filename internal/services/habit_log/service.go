package habitlog

import (
	dc "github.com/mobml/ant/internal/domain/common"
	dhl "github.com/mobml/ant/internal/domain/habit_log"
	"github.com/mobml/ant/internal/models"
	repo "github.com/mobml/ant/internal/repositories"
)

type HabitLogService interface {
	CreateHabitLog(habitLog *models.HabitLog) error
	ListHabitLogs() ([]*models.HabitLog, error)
	ListHabitLogsByHabitID(habitID string) ([]*models.HabitLog, error)
	UpdateHabitLog(habitLog *models.HabitLog) error
	DeleteHabitLog(id string) error
}

type habitLogService struct {
	habitLogRepo repo.HabitLogRepository
}

func NewHabitLogService(habitLogRepo repo.HabitLogRepository) HabitLogService {
	return &habitLogService{
		habitLogRepo: habitLogRepo,
	}
}

func (hls *habitLogService) CreateHabitLog(habitLog *models.HabitLog) error {

	if err := dhl.ValidateHabitLog(habitLog); err != nil {
		return err
	}

	return hls.habitLogRepo.Create(habitLog)
}

func (hls *habitLogService) ListHabitLogs() ([]*models.HabitLog, error) {
	return hls.habitLogRepo.List()
}

func (hls *habitLogService) ListHabitLogsByHabitID(habitID string) ([]*models.HabitLog, error) {
	if err := dc.ValidateID(habitID); err != nil {
		return nil, err
	}

	return hls.habitLogRepo.FindByHabitID(habitID)
}

func (hls *habitLogService) UpdateHabitLog(habitLog *models.HabitLog) error {

	if err := dhl.ValidateHabitLog(habitLog); err != nil {
		return err
	}

	if err := dc.ValidateID(habitLog.ID); err != nil {
		return err
	}

	return hls.habitLogRepo.Update(habitLog)
}

func (hls *habitLogService) DeleteHabitLog(id string) error {

	if err := dc.ValidateID(id); err != nil {
		return err
	}
	return hls.habitLogRepo.Delete(id)
}
