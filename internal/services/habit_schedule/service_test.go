package habitschedule

import (
	"testing"

	"github.com/mobml/ant/internal/models"
)

type mockHabitScheduleRepository struct {
	createCalled bool
	listCalled   bool
	updateCalled bool
	deleteCalled bool

	habitSchedules []*models.HabitSchedule
}

func (m *mockHabitScheduleRepository) Create(habitSchedule *models.HabitSchedule) error {
	m.createCalled = true
	return nil
}

func (m *mockHabitScheduleRepository) List() ([]*models.HabitSchedule, error) {
	m.listCalled = true
	return m.habitSchedules, nil
}

func (m *mockHabitScheduleRepository) FindByID(id string) (*models.HabitSchedule, error) {
	return nil, nil
}

func (m *mockHabitScheduleRepository) Update(habitSchedule *models.HabitSchedule) error {
	m.updateCalled = true
	return nil
}

func (m *mockHabitScheduleRepository) Delete(id string) error {
	m.deleteCalled = true
	return nil
}

func TestHabitScheduleService_CreateHabitSchedule_OK(t *testing.T) {
	repo := &mockHabitScheduleRepository{}
	service := NewHabitScheduleService(repo)

	habitSchedule := &models.HabitSchedule{
		HabitID: "habit123",
	}

	err := service.CreateHabitSchedule(habitSchedule)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !repo.createCalled {
		t.Fatal("expected Create to be called")
	}
}

func TestHabitScheduleService_CreateHabitSchedule_Invalid(t *testing.T) {
	repo := &mockHabitScheduleRepository{}
	service := NewHabitScheduleService(repo)

	habitSchedule := &models.HabitSchedule{}

	err := service.CreateHabitSchedule(habitSchedule)

	if err == nil {
		t.Fatal("expected validation error")
	}

	if repo.createCalled {
		t.Fatal("Create should NOT be called when validation fails")
	}
}

func TestHabitScheduleService_ListHabitSchedules_OK(t *testing.T) {
	repo := &mockHabitScheduleRepository{
		habitSchedules: []*models.HabitSchedule{
			{ID: "1"},
			{ID: "2"},
		},
	}

	service := NewHabitScheduleService(repo)

	schedules, err := service.ListHabitSchedules()

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !repo.listCalled {
		t.Fatal("expected List to be called")
	}

	if len(schedules) != 2 {
		t.Fatalf("expected 2 habit schedules, got %d", len(schedules))
	}
}

func TestHabitScheduleService_UpdateHabitSchedule_OK(t *testing.T) {
	repo := &mockHabitScheduleRepository{}
	service := NewHabitScheduleService(repo)

	habitSchedule := &models.HabitSchedule{
		ID:      "schedule123",
		HabitID: "habit123",
	}

	err := service.UpdateHabitSchedule(habitSchedule)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !repo.updateCalled {
		t.Fatal("expected Update to be called")
	}
}

func TestHabitScheduleService_UpdateHabitSchedule_Invalid(t *testing.T) {
	repo := &mockHabitScheduleRepository{}
	service := NewHabitScheduleService(repo)

	habitSchedule := &models.HabitSchedule{
		ID: "",
	}

	err := service.UpdateHabitSchedule(habitSchedule)

	if err == nil {
		t.Fatal("expected validation error")
	}

	if repo.updateCalled {
		t.Fatal("Update should NOT be called when validation fails")
	}
}

func TestHabitScheduleService_DeleteHabitSchedule_InvalidID(t *testing.T) {
	repo := &mockHabitScheduleRepository{}
	service := NewHabitScheduleService(repo)

	err := service.DeleteHabitSchedule("")

	if err == nil {
		t.Fatal("expected validation error")
	}

	if repo.deleteCalled {
		t.Fatal("Delete should NOT be called when ID is invalid")
	}
}
