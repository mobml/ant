package repositories

import (
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/mobml/ant/internal/models"
)

func TestHabitLogRepository_Create(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	repo := NewHabitLogRepository(db)

	h := &models.HabitLog{
		ID:        "log1",
		HabitID:   "habit1",
		LogDate:   time.Now(),
		Value:     5.5,
		Note:      "Good day",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	mock.ExpectExec(regexp.QuoteMeta(`
        INSERT INTO habit_logs (id, habit_id, log_date, value, note, created_at, updated_at)
        VALUES (?, ?, ?, ?, ?, ?, ?)
    `)).
		WithArgs(
			h.ID,
			h.HabitID,
			h.LogDate,
			h.Value,
			h.Note,
			h.CreatedAt,
			h.UpdatedAt,
		).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := repo.Create(h)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestHabitLogRepository_List(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	repo := NewHabitLogRepository(db)

	rows := sqlmock.NewRows([]string{
		"id", "habit_id", "log_date", "value", "note", "created_at", "updated_at",
	}).AddRow(
		"log1", "habit1", time.Now(), 5.5, "Note 1", time.Now(), time.Now(),
	).AddRow(
		"log2", "habit2", time.Now(), 7.0, "Note 2", time.Now(), time.Now(),
	)

	mock.ExpectQuery("SELECT \\* FROM habit_logs").
		WillReturnRows(rows)

	logs, err := repo.List()

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(logs) != 2 {
		t.Fatalf("expected 2 logs, got %d", len(logs))
	}
}

func TestHabitLogRepository_FindByID(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	repo := NewHabitLogRepository(db)

	row := sqlmock.NewRows([]string{
		"id", "habit_id", "log_date", "value", "note", "created_at", "updated_at",
	}).AddRow(
		"log1", "habit1", time.Now(), 3.0, "Test", time.Now(), time.Now(),
	)

	mock.ExpectQuery("SELECT \\* FROM habit_logs WHERE id = \\?").
		WithArgs("log1").
		WillReturnRows(row)

	log, err := repo.FindByID("log1")

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if log.ID != "log1" {
		t.Fatalf("expected ID log1, got %s", log.ID)
	}
}

func TestHabitLogRepository_FindByHabitID(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	repo := NewHabitLogRepository(db)

	rows := sqlmock.NewRows([]string{
		"id", "habit_id", "log_date", "value", "note", "created_at", "updated_at",
	}).AddRow(
		"log1", "habitX", time.Now(), 1.0, "Note A", time.Now(), time.Now(),
	).AddRow(
		"log2", "habitX", time.Now(), 2.0, "Note B", time.Now(), time.Now(),
	)

	mock.ExpectQuery("SELECT \\* FROM habit_logs WHERE habit_id = \\?").
		WithArgs("habitX").
		WillReturnRows(rows)

	logs, err := repo.FindByHabitID("habitX")

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(logs) != 2 {
		t.Fatalf("expected 2 logs, got %d", len(logs))
	}
}

func TestHabitLogRepository_Update(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	repo := NewHabitLogRepository(db)

	h := &models.HabitLog{
		ID:        "log1",
		HabitID:   "habitUpdated",
		LogDate:   time.Now(),
		Value:     10.0,
		Note:      "Updated note",
		UpdatedAt: time.Now(),
	}

	mock.ExpectExec("UPDATE habit_logs").
		WithArgs(
			h.Value,
			h.Note,
			h.UpdatedAt,
			h.ID,
		).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := repo.Update(h)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestHabitLogRepository_Delete(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	repo := NewHabitLogRepository(db)

	mock.ExpectExec("DELETE FROM habit_logs WHERE id = \\?").
		WithArgs("log1").
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := repo.Delete("log1")

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}
