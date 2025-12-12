package plan

import (
	"github.com/mobml/ant/internal/models"
	"testing"
	"time"
)

type mockPlanRepository struct {
	createCalled bool
	listCalled   bool
	updateCalled bool
	deleteCalled bool

	createErr error
	listErr   error
	updateErr error
	deleteErr error

	plans []*models.Plan
}

func (m *mockPlanRepository) Create(plan *models.Plan) error {
	m.createCalled = true
	return m.createErr
}

func (m *mockPlanRepository) List() ([]*models.Plan, error) {
	m.listCalled = true
	return m.plans, m.listErr
}

func (m *mockPlanRepository) FindByID(id string) (*models.Plan, error) {
	return nil, nil
}

func (m *mockPlanRepository) Update(plan *models.Plan) error {
	m.updateCalled = true
	return m.updateErr
}

func (m *mockPlanRepository) Delete(id string) error {
	m.deleteCalled = true
	return m.deleteErr
}

func TestPlanService_CreatePlan_OK(t *testing.T) {
	repo := &mockPlanRepository{}
	service := NewPlanService(repo)

	plan := &models.Plan{
		Name:        "My Plan",
		Description: "Desc",
		StartDate:   time.Now().Add(1 * time.Hour),
		Duration:    30,
	}

	err := service.CreatePlan(plan)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !repo.createCalled {
		t.Fatal("expected Create to be called")
	}
}

func TestPlanService_CreatePlan_Invalid(t *testing.T) {
	repo := &mockPlanRepository{}
	service := NewPlanService(repo)

	plan := &models.Plan{
		Name:     "",
		Duration: 0,
	}

	err := service.CreatePlan(plan)

	if err == nil {
		t.Fatalf("expected validation error, got nil")
	}

	if repo.createCalled {
		t.Fatalf("repository Create should not be called when validation fails")
	}
}

func TestPlanService_ListPlans(t *testing.T) {
	repo := &mockPlanRepository{
		plans: []*models.Plan{
			{ID: "1", Name: "Plan 1"},
			{ID: "2", Name: "Plan 2"},
		},
	}

	service := NewPlanService(repo)

	plans, err := service.ListPlans()

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !repo.listCalled {
		t.Fatal("expected List to be called")
	}

	if len(plans) != 2 {
		t.Fatalf("expected 2 plans, got %d", len(plans))
	}
}

func TestPlanService_UpdatePlan_OK(t *testing.T) {
	repo := &mockPlanRepository{}
	service := NewPlanService(repo)

	plan := &models.Plan{
		ID:          "abc123",
		Name:        "Updated",
		Description: "Desc",
		StartDate:   time.Now().Add(1 * time.Hour),
		Duration:    10,
	}

	err := service.UpdatePlan(plan)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !repo.updateCalled {
		t.Fatal("expected Update to be called")
	}
}

func TestPlanService_DeletePlan_InvalidID(t *testing.T) {
	repo := &mockPlanRepository{}
	service := NewPlanService(repo)

	err := service.DeletePlan("")

	if err == nil {
		t.Fatal("expected validation error")
	}

	if repo.deleteCalled {
		t.Fatal("Delete should NOT be called with invalid ID")
	}
}
