package habit

import (
	"testing"

	"github.com/mobml/ant/internal/models"
)

type mockHabitRepository struct {
	createCalled bool
	listCalled   bool
	updateCalled bool
	deleteCalled bool

	habits []*models.Habit
}

func (m *mockHabitRepository) Create(habit *models.Habit) error {
	m.createCalled = true
	return nil
}

func (m *mockHabitRepository) CreateHabit(habit *models.Habit, days []int) error {
	m.createCalled = true
	return nil
}

func (m *mockHabitRepository) CreateHabitSchedule(id string, days []int) error {
	m.createCalled = true
	return nil
}

func (m *mockHabitRepository) List() ([]*models.Habit, error) {
	m.listCalled = true
	return m.habits, nil
}

func (m *mockHabitRepository) HabitsForToday(day int) ([]models.HabitWithStatus, error) {
	return []models.HabitWithStatus{}, nil
}

func (m *mockHabitRepository) FindByID(id string) (*models.Habit, error) {
	return nil, nil
}

func (m *mockHabitRepository) Update(habit *models.Habit) error {
	m.updateCalled = true
	return nil
}

func (m *mockHabitRepository) Delete(id string) error {
	m.deleteCalled = true
	return nil
}

func (m *mockHabitRepository) DeleteHabitSchedules(id string) error {
	m.deleteCalled = true
	return nil
}

func TestHabitService_CreateHabit_OK(t *testing.T) {
	repo := &mockHabitRepository{}
	service := NewHabitService(repo)

	habit := &models.Habit{
		GoalID:      "goal123",
		Name:        "Drink Water",
		MeasureType: models.MeasureInteger,
	}

	err := service.CreateHabit(habit)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !repo.createCalled {
		t.Fatal("expected Create to be called")
	}
}

func TestHabitService_CreateHabitWithSchedule_OK(t *testing.T) {
	repo := &mockHabitRepository{}
	service := NewHabitService(repo)

	habit := &models.Habit{
		GoalID:      "goal123",
		Name:        "Exercise",
		MeasureType: models.MeasureInteger,
	}

	days := []int{1, 3, 5} // Monday, Wednesday, Friday

	err := service.CreateHabitWithSchedule(habit, days)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !repo.createCalled {
		t.Fatal("expected CreateHabit to be called")
	}
}

func TestHabitService_CreateHabit_Invalid(t *testing.T) {
	repo := &mockHabitRepository{}
	service := NewHabitService(repo)

	habit := &models.Habit{
		Name: "",
	}

	err := service.CreateHabit(habit)

	if err == nil {
		t.Fatal("expected validation error")
	}

	if repo.createCalled {
		t.Fatal("Create should NOT be called when validation fails")
	}
}

func TestHabitService_GetHabitByID_OK(t *testing.T) {
	repo := &mockHabitRepository{}
	service := NewHabitService(repo)

	_, err := service.GetHabitByID("habit123")

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestHabitService_GetHabitByID_InvalidID(t *testing.T) {
	repo := &mockHabitRepository{}
	service := NewHabitService(repo)

	_, err := service.GetHabitByID("")

	if err == nil {
		t.Fatal("expected validation error")
	}
}

func TestHabitService_ListHabits_OK(t *testing.T) {
	repo := &mockHabitRepository{
		habits: []*models.Habit{
			{ID: "1", Name: "Habit 1"},
			{ID: "2", Name: "Habit 2"},
		},
	}

	service := NewHabitService(repo)

	habits, err := service.ListHabits()

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !repo.listCalled {
		t.Fatal("expected List to be called")
	}

	if len(habits) != 2 {
		t.Fatalf("expected 2 habits, got %d", len(habits))
	}
}

func TestHabitService_UpdateHabit_OK(t *testing.T) {
	repo := &mockHabitRepository{}
	service := NewHabitService(repo)

	habit := &models.Habit{
		ID:          "habit123",
		Name:        "Updated Habit",
		MeasureType: models.MeasureFloat,
		GoalID:      "goal123",
	}

	err := service.UpdateHabit(habit)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !repo.updateCalled {
		t.Fatal("expected Update to be called")
	}
}

func TestHabitService_UpdateHabit_Invalid(t *testing.T) {
	repo := &mockHabitRepository{}
	service := NewHabitService(repo)

	habit := &models.Habit{
		ID: "",
	}

	err := service.UpdateHabit(habit)

	if err == nil {
		t.Fatal("expected validation error")
	}

	if repo.updateCalled {
		t.Fatal("Update should NOT be called when validation fails")
	}
}

func TestHabitService_DeleteHabit_InvalidID(t *testing.T) {
	repo := &mockHabitRepository{}
	service := NewHabitService(repo)

	err := service.DeleteHabit("")

	if err == nil {
		t.Fatal("expected validation error")
	}

	if repo.deleteCalled {
		t.Fatal("Delete should NOT be called when ID is invalid")
	}
}
