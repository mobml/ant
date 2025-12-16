package area

import (
	"testing"

	"github.com/mobml/ant/internal/models"
)

type mockAreaRepository struct {
	createCalled bool
	listCalled   bool
	updateCalled bool
	deleteCalled bool

	areas []*models.Area
}

func (m *mockAreaRepository) Create(area *models.Area) error {
	m.createCalled = true
	return nil
}

func (m *mockAreaRepository) ListByPlan(planID string) ([]*models.Area, error) {
	m.listCalled = true
	return m.areas, nil
}

func (m *mockAreaRepository) FindByID(id string) (*models.Area, error) {
	return nil, nil
}

func (m *mockAreaRepository) List() ([]*models.Area, error) {
	return nil, nil
}

func (m *mockAreaRepository) Update(area *models.Area) error {
	m.updateCalled = true
	return nil
}

func (m *mockAreaRepository) Delete(id string) error {
	m.deleteCalled = true
	return nil
}


func TestAreaService_CreateArea_OK(t *testing.T) {
	repo := &mockAreaRepository{}
	service := NewAreaService(repo)

	area := &models.Area{
		PlanID: "plan123",
		Name:   "Health",
	}

	err := service.CreateArea(area)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !repo.createCalled {
		t.Fatal("expected Create to be called")
	}
}

func TestAreaService_CreateArea_Invalid(t *testing.T) {
	repo := &mockAreaRepository{}
	service := NewAreaService(repo)

	area := &models.Area{
		PlanID: "",
		Name:   "",
	}

	err := service.CreateArea(area)

	if err == nil {
		t.Fatal("expected validation error")
	}

	if repo.createCalled {
		t.Fatal("Create should NOT be called when validation fails")
	}
}

func TestAreaService_ListAreasByPlan_OK(t *testing.T) {
	repo := &mockAreaRepository{
		areas: []*models.Area{
			{ID: "1", Name: "Area 1"},
			{ID: "2", Name: "Area 2"},
		},
	}

	service := NewAreaService(repo)

	areas, err := service.ListAreasByPlan("plan123")

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !repo.listCalled {
		t.Fatal("expected ListByPlan to be called")
	}

	if len(areas) != 2 {
		t.Fatalf("expected 2 areas, got %d", len(areas))
	}
}

func TestAreaService_ListAreasByPlan_InvalidPlanID(t *testing.T) {
	repo := &mockAreaRepository{}
	service := NewAreaService(repo)

	_, err := service.ListAreasByPlan("")

	if err == nil {
		t.Fatal("expected validation error")
	}

	if repo.listCalled {
		t.Fatal("ListByPlan should NOT be called with invalid planID")
	}
}

func TestAreaService_UpdateArea_OK(t *testing.T) {
	repo := &mockAreaRepository{}
	service := NewAreaService(repo)

	area := &models.Area{
		ID:     "area123",
		PlanID: "plan123",
		Name:   "Updated",
	}

	err := service.UpdateArea(area)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !repo.updateCalled {
		t.Fatal("expected Update to be called")
	}
}

func TestAreaService_UpdateArea_Invalid(t *testing.T) {
	repo := &mockAreaRepository{}
	service := NewAreaService(repo)

	area := &models.Area{
		ID: "",
	}

	err := service.UpdateArea(area)

	if err == nil {
		t.Fatal("expected validation error")
	}

	if repo.updateCalled {
		t.Fatal("Update should NOT be called when validation fails")
	}
}

func TestAreaService_DeleteArea_InvalidID(t *testing.T) {
	repo := &mockAreaRepository{}
	service := NewAreaService(repo)

	err := service.DeleteArea("")

	if err == nil {
		t.Fatal("expected validation error")
	}

	if repo.deleteCalled {
		t.Fatal("Delete should NOT be called when ID is invalid")
	}
}