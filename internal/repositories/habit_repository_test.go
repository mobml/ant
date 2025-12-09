package repositories

import (
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/mobml/ant/internal/models"
)

func TestHabitRepository_Create(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	repo := NewHabitRepository(db)

	h := &models.Habit{
		GoalID:      "goal1",
		Name:        "Drink water",
		Description: "8 glasses a day",
		MeasureType: models.MeasureInteger,
		MeasureUnit: "glasses",
	}

	mock.ExpectExec(regexp.QuoteMeta(`
		INSERT INTO habits (id, goal_id, name, description, measure_type, measure_unit)
		VALUES (?, ?, ?, ?, ?, ?)
	`)).
		WithArgs(
			sqlmock.AnyArg(),
			h.GoalID,
			h.Name,
			h.Description,
			h.MeasureType,
			h.MeasureUnit,
		).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := repo.Create(h)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestHabitRepository_List(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	repo := NewHabitRepository(db)

	rows := sqlmock.NewRows([]string{
		"id", "goal_id", "name", "description", "measure_type", "measure_unit", "created_at", "updated_at",
	}).AddRow(
		"1", "goal1", "Habit 1", "Desc 1", models.MeasureBoolean, "yes/no", time.Now(), time.Now(),
	).AddRow(
		"2", "goal2", "Habit 2", "Desc 2", models.MeasureFloat, "liters", time.Now(), time.Now(),
	)

	mock.ExpectQuery("SELECT \\* FROM habits").
		WillReturnRows(rows)

	habits, err := repo.List()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(habits) != 2 {
		t.Fatalf("expected 2 habits, got %d", len(habits))
	}
}

func TestHabitRepository_FindByID(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	repo := NewHabitRepository(db)

	row := sqlmock.NewRows([]string{
		"id", "goal_id", "name", "description", "measure_type", "measure_unit", "created_at", "updated_at",
	}).AddRow(
		"habit1", "goal1", "Drink water", "Desc", models.MeasureInteger, "glasses", time.Now(), time.Now(),
	)

	mock.ExpectQuery("SELECT \\* FROM habits WHERE id = \\?").
		WithArgs("habit1").
		WillReturnRows(row)

	habit, err := repo.FindByID("habit1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if habit.ID != "habit1" {
		t.Fatalf("expected ID habit1, got %s", habit.ID)
	}
}

func TestHabitRepository_Update(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	repo := NewHabitRepository(db)

	h := &models.Habit{
		ID:          "habit1",
		GoalID:      "goalUpdated",
		Name:        "Updated Habit",
		Description: "Updated Desc",
		MeasureType: models.MeasureHours,
		MeasureUnit: "hours",
	}

	mock.ExpectExec("UPDATE habits").
		WithArgs(
			h.GoalID,
			h.Name,
			h.Description,
			h.MeasureType,
			h.MeasureUnit,
			sqlmock.AnyArg(),
			h.ID,
		).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := repo.Update(h)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestHabitRepository_Delete(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	repo := NewHabitRepository(db)

	mock.ExpectExec("DELETE FROM habits WHERE id = \\?").
		WithArgs("habit1").
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := repo.Delete("habit1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}
