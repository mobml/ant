package area

import (
	da "github.com/mobml/ant/internal/domain/area"
	dc "github.com/mobml/ant/internal/domain/common"
	"github.com/mobml/ant/internal/models"
	repo "github.com/mobml/ant/internal/repositories"
)

type AreaService interface {
	CreateArea(area *models.Area) error
	ListAreasByPlan(planID string) ([]*models.Area, error)
	UpdateArea(area *models.Area) error
	DeleteArea(id string) error
}

type areaService struct {
	areaRepo repo.AreaRepository
}

func NewAreaService(areaRepo repo.AreaRepository) AreaService {
	return &areaService{
		areaRepo: areaRepo,
	}
}

func (s *areaService) CreateArea(area *models.Area) error {
	if err := da.ValidateArea(area); err != nil {
		return err
	}
	return s.areaRepo.Create(area)
}

func (s *areaService) ListAreasByPlan(planID string) ([]*models.Area, error) {
	if err := dc.ValidateID(planID); err != nil {
		return nil, err
	}
	return s.areaRepo.ListByPlan(planID)
}

func (s *areaService) UpdateArea(area *models.Area) error {
	if err := da.ValidateArea(area); err != nil {
		return err
	}
	if err := dc.ValidateID(area.ID); err != nil {
		return err
	}
	return s.areaRepo.Update(area)
}

func (s *areaService) DeleteArea(id string) error {
	if err := dc.ValidateID(id); err != nil {
		return err
	}
	return s.areaRepo.Delete(id)
}
