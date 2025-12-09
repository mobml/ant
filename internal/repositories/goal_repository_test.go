package repositories

import (
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/mobml/ant/internal/models"
)

func TestGoalRepository_Create(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	repo := NewGoalRepository(db)

	g := &models.Goal{
		ID:          "goal1",
		AreaID:      "area1",
		Name:        "My Goal",
		Description: "Desc",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	mock.ExpectExec(regexp.QuoteMeta(`
		INSERT INTO goals (id, area_id, name, description, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?, ?)
	`)).
		WithArgs(
			g.ID,
			g.AreaID,
			g.Name,
			g.Description,
			g.CreatedAt,
			g.UpdatedAt,
		).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := repo.Create(g)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestGoalRepository_List(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	repo := NewGoalRepository(db)

	rows := sqlmock.NewRows([]string{
		"id", "area_id", "name", "description", "created_at", "updated_at",
	}).AddRow(
		"1", "area1", "Goal 1", "Desc 1", time.Now(), time.Now(),
	).AddRow(
		"2", "area2", "Goal 2", "Desc 2", time.Now(), time.Now(),
	)

	mock.ExpectQuery("SELECT \\* FROM goals").
		WillReturnRows(rows)

	goals, err := repo.List()

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(goals) != 2 {
		t.Fatalf("expected 2 goals, got %d", len(goals))
	}
}

func TestGoalRepository_FindByID(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	repo := NewGoalRepository(db)

	row := sqlmock.NewRows([]string{
		"id", "area_id", "name", "description", "created_at", "updated_at",
	}).AddRow(
		"goal1", "area1", "My Goal", "Desc", time.Now(), time.Now(),
	)

	mock.ExpectQuery("SELECT \\* FROM goals WHERE id = \\?").
		WithArgs("goal1").
		WillReturnRows(row)

	goal, err := repo.FindByID("goal1")

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if goal.ID != "goal1" {
		t.Fatalf("expected ID goal1, got %s", goal.ID)
	}
}

func TestGoalRepository_Update(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	repo := NewGoalRepository(db)

	g := &models.Goal{
		ID:          "goal1",
		AreaID:      "areaUpdated",
		Name:        "Updated",
		Description: "Updated Desc",
		UpdatedAt:   time.Now(),
	}

	mock.ExpectExec("UPDATE goals").
		WithArgs(
			g.AreaID,
			g.Name,
			g.Description,
			g.UpdatedAt,
			g.ID,
		).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := repo.Update(g)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestGoalRepository_Delete(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	repo := NewGoalRepository(db)

	mock.ExpectExec("DELETE FROM goals WHERE id = \\?").
		WithArgs("goal1").
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := repo.Delete("goal1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}
