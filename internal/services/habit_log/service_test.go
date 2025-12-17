package habitlog

import (
	"testing"
	"time"

	"github.com/mobml/ant/internal/models"
)

type mockHabitLogRepository struct {
	createCalled bool
	listCalled   bool
	listByHabit  bool
	updateCalled bool
	deleteCalled bool

	habitLogs []*models.HabitLog
}

func (m *mockHabitLogRepository) Create(habitLog *models.HabitLog) error {
	m.createCalled = true
	return nil
}

func (m *mockHabitLogRepository) List() ([]*models.HabitLog, error) {
	m.listCalled = true
	return m.habitLogs, nil
}

func (m *mockHabitLogRepository) FindByHabitID(habitID string) ([]*models.HabitLog, error) {
	m.listByHabit = true
	return m.habitLogs, nil
}

func (m *mockHabitLogRepository) FindByID(id string) (*models.HabitLog, error) {
	return nil, nil
}

func (m *mockHabitLogRepository) Update(habitLog *models.HabitLog) error {
	m.updateCalled = true
	return nil
}

func (m *mockHabitLogRepository) Delete(id string) error {
	m.deleteCalled = true
	return nil
}

func TestHabitLogService_CreateHabitLog_OK(t *testing.T) {
	repo := &mockHabitLogRepository{}
	service := NewHabitLogService(repo)

	habitLog := &models.HabitLog{
		HabitID: "habit123",
		Value:   10,
		LogDate: time.Now(),
	}

	err := service.CreateHabitLog(habitLog)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !repo.createCalled {
		t.Fatal("expected Create to be called")
	}
}

func TestHabitLogService_CreateHabitLog_Invalid(t *testing.T) {
	repo := &mockHabitLogRepository{}
	service := NewHabitLogService(repo)

	habitLog := &models.HabitLog{
		HabitID: "",
	}

	err := service.CreateHabitLog(habitLog)

	if err == nil {
		t.Fatal("expected validation error")
	}

	if repo.createCalled {
		t.Fatal("Create should NOT be called when validation fails")
	}
}

func TestHabitLogService_ListHabitLogs_OK(t *testing.T) {
	repo := &mockHabitLogRepository{
		habitLogs: []*models.HabitLog{
			{ID: "1"},
			{ID: "2"},
		},
	}

	service := NewHabitLogService(repo)

	logs, err := service.ListHabitLogs()

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !repo.listCalled {
		t.Fatal("expected List to be called")
	}

	if len(logs) != 2 {
		t.Fatalf("expected 2 habit logs, got %d", len(logs))
	}
}

func TestHabitLogService_ListHabitLogsByHabitID_OK(t *testing.T) {
	repo := &mockHabitLogRepository{
		habitLogs: []*models.HabitLog{
			{ID: "1"},
			{ID: "2"},
		},
	}

	service := NewHabitLogService(repo)

	logs, err := service.ListHabitLogsByHabitID("habit123")

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !repo.listByHabit {
		t.Fatal("expected FindByHabitID to be called")
	}

	if len(logs) != 2 {
		t.Fatalf("expected 2 habit logs, got %d", len(logs))
	}
}

func TestHabitLogService_ListHabitLogsByHabitID_InvalidID(t *testing.T) {
	repo := &mockHabitLogRepository{}
	service := NewHabitLogService(repo)

	_, err := service.ListHabitLogsByHabitID("")

	if err == nil {
		t.Fatal("expected validation error")
	}

	if repo.listByHabit {
		t.Fatal("FindByHabitID should NOT be called with invalid habitID")
	}
}

func TestHabitLogService_UpdateHabitLog_OK(t *testing.T) {
	repo := &mockHabitLogRepository{}
	service := NewHabitLogService(repo)

	habitLog := &models.HabitLog{
		ID:      "log123",
		HabitID: "habit123",
		Value:   20,
		LogDate: time.Now(),
	}

	err := service.UpdateHabitLog(habitLog)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !repo.updateCalled {
		t.Fatal("expected Update to be called")
	}
}

func TestHabitLogService_UpdateHabitLog_Invalid(t *testing.T) {
	repo := &mockHabitLogRepository{}
	service := NewHabitLogService(repo)

	habitLog := &models.HabitLog{
		ID: "",
	}

	err := service.UpdateHabitLog(habitLog)

	if err == nil {
		t.Fatal("expected validation error")
	}

	if repo.updateCalled {
		t.Fatal("Update should NOT be called when validation fails")
	}
}

func TestHabitLogService_DeleteHabitLog_InvalidID(t *testing.T) {
	repo := &mockHabitLogRepository{}
	service := NewHabitLogService(repo)

	err := service.DeleteHabitLog("")

	if err == nil {
		t.Fatal("expected validation error")
	}

	if repo.deleteCalled {
		t.Fatal("Delete should NOT be called when ID is invalid")
	}
}
