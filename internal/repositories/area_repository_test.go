package repositories

import (
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/mobml/ant/internal/models"
)

func TestAreaRepository_Create(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	repo := NewAreaRepository(db)

	a := &models.Area{
		PlanID:      "plan1",
		Name:        "Area 1",
		Description: "Desc",
	}

	mock.ExpectExec(regexp.QuoteMeta(`
		INSERT INTO areas (id, plan_id, name, description)
		VALUES (?, ?, ?, ?)
	`)).
		WithArgs(
			sqlmock.AnyArg(),
			a.PlanID,
			a.Name,
			a.Description,
		).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := repo.Create(a)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestAreaRepository_List(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	repo := NewAreaRepository(db)

	rows := sqlmock.NewRows([]string{
		"id", "plan_id", "name", "description", "created_at", "updated_at",
	}).AddRow(
		"1", "plan1", "Area 1", "Desc 1", time.Now(), time.Now(),
	).AddRow(
		"2", "plan2", "Area 2", "Desc 2", time.Now(), time.Now(),
	)

	mock.ExpectQuery("SELECT \\* FROM areas").
		WillReturnRows(rows)

	areas, err := repo.List()

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(areas) != 2 {
		t.Fatalf("expected 2 areas, got %d", len(areas))
	}
}

func TestAreaRepository_FindByID(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	repo := NewAreaRepository(db)

	row := sqlmock.NewRows([]string{
		"id", "plan_id", "name", "description", "created_at", "updated_at",
	}).AddRow(
		"area1", "plan1", "Area 1", "Desc", time.Now(), time.Now(),
	)

	mock.ExpectQuery("SELECT \\* FROM areas WHERE id = \\?").
		WithArgs("area1").
		WillReturnRows(row)

	area, err := repo.FindByID("area1")

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if area.ID != "area1" {
		t.Fatalf("expected ID area1, got %s", area.ID)
	}
}

func TestAreaRepository_ListByPlanID(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()
	
	repo := NewAreaRepository(db)
	rows := sqlmock.NewRows([]string{
		"id", "plan_id", "name", "description", "created_at", "updated_at",
	}).AddRow(
		"area1", "plan1", "Area 1", "Desc 1", time.Now(), time.Now(),
	).AddRow(
		"area2", "plan1", "Area 2", "Desc 2", time.Now(), time.Now(),
	)
	
	mock.ExpectQuery("SELECT \\* FROM areas WHERE plan_id = \\?").
		WithArgs("plan1").
		WillReturnRows(rows)
			
	areas, err := repo.ListByPlan("plan1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	
	if len(areas) != 2 {
		t.Fatalf("expected 2 areas, got %d", len(areas))
	}
}

func TestAreaRepository_Update(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	repo := NewAreaRepository(db)

	a := &models.Area{
		ID:          "area1",
		PlanID:      "planUpdated",
		Name:        "Updated Area",
		Description: "Updated Desc",
	}

	mock.ExpectExec("UPDATE areas").
		WithArgs(
			a.PlanID,
			a.Name,
			a.Description,
			sqlmock.AnyArg(),
			a.ID,
		).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := repo.Update(a)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestAreaRepository_Delete(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	repo := NewAreaRepository(db)

	mock.ExpectExec("DELETE FROM areas WHERE id = \\?").
		WithArgs("area1").
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := repo.Delete("area1")

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}
