package repositories

import (
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/mobml/ant/internal/models"
)

func TestPlanRepository_Create(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	repo := NewPlanRepository(db)

	p := &models.Plan{
		Name:        "My Plan",
		Description: "Desc",
		StartDate:   time.Now(),
		Duration:    30,
	}

	mock.ExpectExec(regexp.QuoteMeta(`
		INSERT INTO plans (id, name, description, start_date, duration)
		VALUES (?, ?, ?, ?, ?)
	`)).
		WithArgs(
			sqlmock.AnyArg(),
			p.Name,
			p.Description,
			p.StartDate,
			p.Duration,
		).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := repo.Create(p)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestPlanRepository_List(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	repo := NewPlanRepository(db)

	rows := sqlmock.NewRows([]string{
		"id", "name", "description", "start_date", "duration", "created_at", "updated_at",
	}).AddRow(
		"1", "Plan 1", "Desc 1", time.Now(), 10, time.Now(), time.Now(),
	).AddRow(
		"2", "Plan 2", "Desc 2", time.Now(), 20, time.Now(), time.Now(),
	)

	mock.ExpectQuery("SELECT \\* FROM plans").
		WillReturnRows(rows)

	plans, err := repo.List()

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(plans) != 2 {
		t.Fatalf("expected 2 plans, got %d", len(plans))
	}
}

func TestPlanRepository_FindByID(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	repo := NewPlanRepository(db)

	row := sqlmock.NewRows([]string{
		"id", "name", "description", "start_date", "duration", "created_at", "updated_at",
	}).AddRow(
		"abc123", "My Plan", "Desc", time.Now(), 30, time.Now(), time.Now(),
	)

	mock.ExpectQuery("SELECT \\* from plans WHERE id = \\?").
		WithArgs("abc123").
		WillReturnRows(row)

	plan, err := repo.FindByID("abc123")

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if plan.ID != "abc123" {
		t.Fatalf("expected ID abc123, got %s", plan.ID)
	}
}

func TestPlanRepository_Update(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	repo := NewPlanRepository(db)

	p := &models.Plan{
		ID:          "abc123",
		Name:        "Updated",
		Description: "Updated Desc",
		StartDate:   time.Now(),
		Duration:    40,
	}

	mock.ExpectExec("UPDATE plans").
		WithArgs(
			p.Name,
			p.Description,
			p.StartDate,
			p.Duration,
			sqlmock.AnyArg(),
			p.ID,
		).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := repo.Update(p)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestPlanRepository_Delete(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	repo := NewPlanRepository(db)

	mock.ExpectExec("DELETE FROM plans WHERE id = \\?").
		WithArgs("abc123").
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := repo.Delete("abc123")

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}
