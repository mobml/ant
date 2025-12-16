package goal

import (
	"testing"

	"github.com/mobml/ant/internal/models"
)

type mockGoalRepository struct {
	createCalled bool
	listCalled   bool
	updateCalled bool
	deleteCalled bool

	goals []*models.Goal
}

func (m *mockGoalRepository) Create(goal *models.Goal) error {
	m.createCalled = true
	return nil
}

func (m *mockGoalRepository) ListByArea(areaID string) ([]*models.Goal, error) {
	m.listCalled = true
	return m.goals, nil
}

func (m *mockGoalRepository) FindByID(id string) (*models.Goal, error) {
	return nil, nil
}

func (m *mockGoalRepository) List() ([]*models.Goal, error) {
	return nil, nil
}

func (m *mockGoalRepository) Update(goal *models.Goal) error {
	m.updateCalled = true
	return nil
}

func (m *mockGoalRepository) Delete(id string) error {
	m.deleteCalled = true
	return nil
}

func TestGoalService_CreateGoal_OK(t *testing.T) {
	repo := &mockGoalRepository{}
	service := NewGoalService(repo)

	goal := &models.Goal{
		AreaID: "area123",
		Name:   "Run 5km",
	}

	err := service.CreateGoal(goal)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !repo.createCalled {
		t.Fatal("expected Create to be called")
	}
}

func TestGoalService_CreateGoal_Invalid(t *testing.T) {
	repo := &mockGoalRepository{}
	service := NewGoalService(repo)

	goal := &models.Goal{
		AreaID: "",
		Name:   "",
	}

	err := service.CreateGoal(goal)

	if err == nil {
		t.Fatal("expected validation error")
	}

	if repo.createCalled {
		t.Fatal("Create should NOT be called when validation fails")
	}
}

func TestGoalService_ListGoalsByArea_OK(t *testing.T) {
	repo := &mockGoalRepository{
		goals: []*models.Goal{
			{ID: "1", Name: "Goal 1"},
			{ID: "2", Name: "Goal 2"},
		},
	}

	service := NewGoalService(repo)

	goals, err := service.ListGoalsByArea("area123")

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !repo.listCalled {
		t.Fatal("expected ListByArea to be called")
	}

	if len(goals) != 2 {
		t.Fatalf("expected 2 goals, got %d", len(goals))
	}
}

func TestGoalService_ListGoalsByArea_InvalidAreaID(t *testing.T) {
	repo := &mockGoalRepository{}
	service := NewGoalService(repo)

	_, err := service.ListGoalsByArea("")

	if err == nil {
		t.Fatal("expected validation error")
	}

	if repo.listCalled {
		t.Fatal("ListByArea should NOT be called with invalid areaID")
	}
}

func TestGoalService_UpdateGoal_OK(t *testing.T) {
	repo := &mockGoalRepository{}
	service := NewGoalService(repo)

	goal := &models.Goal{
		ID:     "goal123",
		AreaID: "area123",
		Name:   "Updated Goal",
	}

	err := service.UpdateGoal(goal)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !repo.updateCalled {
		t.Fatal("expected Update to be called")
	}
}

func TestGoalService_UpdateGoal_Invalid(t *testing.T) {
	repo := &mockGoalRepository{}
	service := NewGoalService(repo)

	goal := &models.Goal{
		ID: "",
	}

	err := service.UpdateGoal(goal)

	if err == nil {
		t.Fatal("expected validation error")
	}

	if repo.updateCalled {
		t.Fatal("Update should NOT be called when validation fails")
	}
}
