package repositories

import (
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/mobml/ant/internal/models"
)

func TestDailyNoteRepository_Create(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	repo := NewDailyNoteRepository(db)

	n := &models.DailyNote{
		NoteDate: time.Now(),
		Content:  "My note content",
	}

	mock.ExpectExec(regexp.QuoteMeta(`
		INSERT INTO daily_notes (id, note_date, content)
		VALUES (?, ?, ?)
	`)).
		WithArgs(
			sqlmock.AnyArg(),
			n.NoteDate,
			n.Content,
		).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := repo.Create(n)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestDailyNoteRepository_List(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	repo := NewDailyNoteRepository(db)

	rows := sqlmock.NewRows([]string{
		"id", "note_date", "content", "created_at", "updated_at",
	}).AddRow(
		"1", time.Now(), "Note 1", time.Now(), time.Now(),
	).AddRow(
		"2", time.Now(), "Note 2", time.Now(), time.Now(),
	)

	mock.ExpectQuery("SELECT \\* FROM daily_notes").
		WillReturnRows(rows)

	notes, err := repo.List()

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(notes) != 2 {
		t.Fatalf("expected 2 notes, got %d", len(notes))
	}
}

func TestDailyNoteRepository_FindByID(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	repo := NewDailyNoteRepository(db)

	row := sqlmock.NewRows([]string{
		"id", "note_date", "content", "created_at", "updated_at",
	}).AddRow(
		"abc123", time.Now(), "Hello", time.Now(), time.Now(),
	)

	mock.ExpectQuery("SELECT \\* FROM daily_notes WHERE id = \\?").
		WithArgs("abc123").
		WillReturnRows(row)

	note, err := repo.FindByID("abc123")

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if note.ID != "abc123" {
		t.Fatalf("expected id abc123, got %s", note.ID)
	}
}

func TestDailyNoteRepository_Update(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	repo := NewDailyNoteRepository(db)

	n := &models.DailyNote{
		ID:       "abc123",
		NoteDate: time.Now(),
		Content:  "Updated content",
	}

	mock.ExpectExec("UPDATE daily_notes").
		WithArgs(
			n.NoteDate,
			n.Content,
			sqlmock.AnyArg(),
			n.ID,
		).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := repo.Update(n)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestDailyNoteRepository_Delete(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	repo := NewDailyNoteRepository(db)

	mock.ExpectExec("DELETE FROM daily_notes WHERE id = \\?").
		WithArgs("abc123").
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := repo.Delete("abc123")

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}
