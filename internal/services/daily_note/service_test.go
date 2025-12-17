package dailynote

import (
	"testing"
	"time"

	"github.com/mobml/ant/internal/models"
)

type mockDailyNoteRepository struct {
	createCalled bool
	listCalled   bool
	updateCalled bool
	deleteCalled bool

	dailyNotes []*models.DailyNote
}

func (m *mockDailyNoteRepository) Create(dailyNote *models.DailyNote) error {
	m.createCalled = true
	return nil
}

func (m *mockDailyNoteRepository) List() ([]*models.DailyNote, error) {
	m.listCalled = true
	return m.dailyNotes, nil
}

func (m *mockDailyNoteRepository) FindByID(id string) (*models.DailyNote, error) {
	return nil, nil
}

func (m *mockDailyNoteRepository) Update(dailyNote *models.DailyNote) error {
	m.updateCalled = true
	return nil
}

func (m *mockDailyNoteRepository) Delete(id string) error {
	m.deleteCalled = true
	return nil
}

func TestDailyNoteService_CreateDailyNote_OK(t *testing.T) {
	repo := &mockDailyNoteRepository{}
	service := NewDailyNoteService(repo)

	dailyNote := &models.DailyNote{
		NoteDate: time.Now(),
		Content:  "Today was a good day",
	}

	err := service.CreateDailyNote(dailyNote)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !repo.createCalled {
		t.Fatal("expected Create to be called")
	}
}

func TestDailyNoteService_CreateDailyNote_Invalid(t *testing.T) {
	repo := &mockDailyNoteRepository{}
	service := NewDailyNoteService(repo)

	dailyNote := &models.DailyNote{
		Content: "",
	}

	err := service.CreateDailyNote(dailyNote)

	if err == nil {
		t.Fatal("expected validation error")
	}

	if repo.createCalled {
		t.Fatal("Create should NOT be called when validation fails")
	}
}

func TestDailyNoteService_ListDailyNotes_OK(t *testing.T) {
	repo := &mockDailyNoteRepository{
		dailyNotes: []*models.DailyNote{
			{ID: "1", Content: "Note 1"},
			{ID: "2", Content: "Note 2"},
		},
	}

	service := NewDailyNoteService(repo)

	notes, err := service.ListDailyNotes()

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !repo.listCalled {
		t.Fatal("expected List to be called")
	}

	if len(notes) != 2 {
		t.Fatalf("expected 2 daily notes, got %d", len(notes))
	}
}

func TestDailyNoteService_UpdateDailyNote_OK(t *testing.T) {
	repo := &mockDailyNoteRepository{}
	service := NewDailyNoteService(repo)

	dailyNote := &models.DailyNote{
		ID:       "note123",
		NoteDate: time.Now(),
		Content:  "Updated note",
	}

	err := service.UpdateDailyNote(dailyNote)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !repo.updateCalled {
		t.Fatal("expected Update to be called")
	}
}

func TestDailyNoteService_UpdateDailyNote_Invalid(t *testing.T) {
	repo := &mockDailyNoteRepository{}
	service := NewDailyNoteService(repo)

	dailyNote := &models.DailyNote{
		ID: "",
	}

	err := service.UpdateDailyNote(dailyNote)

	if err == nil {
		t.Fatal("expected validation error")
	}

	if repo.updateCalled {
		t.Fatal("Update should NOT be called when validation fails")
	}
}

func TestDailyNoteService_DeleteDailyNote_InvalidID(t *testing.T) {
	repo := &mockDailyNoteRepository{}
	service := NewDailyNoteService(repo)

	err := service.DeleteDailyNote("")

	if err == nil {
		t.Fatal("expected validation error")
	}

	if repo.deleteCalled {
		t.Fatal("Delete should NOT be called when ID is invalid")
	}
}
