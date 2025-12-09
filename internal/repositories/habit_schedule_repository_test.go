package repositories

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/mobml/ant/internal/models"
	"regexp"
	"testing"
	"time"
)

func TestHabitScheduleRepository_Create(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	repo := NewHabitScheduleRepository(db)

	hs := &models.HabitSchedule{
		HabitID:   "habit123",
		DayOfWeek: 3,
	}

	mock.ExpectExec(regexp.QuoteMeta(`
		INSERT INTO habit_schedules (id, habit_id, day_of_week)
		VALUES (?, ?, ?)
	`)).
		WithArgs(
			sqlmock.AnyArg(),
			hs.HabitID,
			hs.DayOfWeek,
		).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := repo.Create(hs)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestHabitScheduleRepository_List(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	repo := NewHabitScheduleRepository(db)

	rows := sqlmock.NewRows([]string{
		"id", "habit_id", "day_of_week", "created_at", "updated_at",
	}).AddRow(
		"1", "habit1", 1, time.Now(), time.Now(),
	).AddRow(
		"2", "habit2", 5, time.Now(), time.Now(),
	)

	mock.ExpectQuery("SELECT \\* FROM habit_schedules").
		WillReturnRows(rows)

	schedules, err := repo.List()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(schedules) != 2 {
		t.Fatalf("expected 2 schedules, got %d", len(schedules))
	}
}

func TestHabitScheduleRepository_FindByID(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	repo := NewHabitScheduleRepository(db)

	row := sqlmock.NewRows([]string{
		"id", "habit_id", "day_of_week", "created_at", "updated_at",
	}).AddRow(
		"abc123", "habitA", 4, time.Now(), time.Now(),
	)

	mock.ExpectQuery("SELECT \\* FROM habit_schedules WHERE id = \\?").
		WithArgs("abc123").
		WillReturnRows(row)

	hs, err := repo.FindByID("abc123")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if hs.ID != "abc123" {
		t.Fatalf("expected id abc123, got %s", hs.ID)
	}
}

func TestHabitScheduleRepository_Update(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	repo := NewHabitScheduleRepository(db)

	hs := &models.HabitSchedule{
		ID:        "abc123",
		HabitID:   "habitX",
		DayOfWeek: 6,
	}

	mock.ExpectExec("UPDATE habit_schedules").
		WithArgs(
			hs.DayOfWeek,
			sqlmock.AnyArg(),
			hs.ID,
		).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := repo.Update(hs)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestHabitScheduleRepository_Delete(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	repo := NewHabitScheduleRepository(db)

	mock.ExpectExec("DELETE FROM habit_schedules WHERE id = \\?").
		WithArgs("abc123").
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := repo.Delete("abc123")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}
