package repositories

import (
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/mobml/ant/internal/models"
)

func TestWeeklyReportRepository_Create(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	repo := NewWeeklyReportRepository(db)

	w := &models.WeeklyReport{
		PlanID:      "plan123",
		WeekStart:   time.Now(),
		WeekEnd:     time.Now().Add(7 * 24 * time.Hour),
		ReportMD:    "# Report",
		GeneratedAt: time.Now(),
	}

	mock.ExpectExec(regexp.QuoteMeta(`
		INSERT INTO weekly_reports (id, plan_id, week_start, week_end, report_md)
		VALUES (?, ?, ?, ?, ?, ?)
	`)).
		WithArgs(
			sqlmock.AnyArg(),
			w.PlanID,
			w.WeekStart,
			w.WeekEnd,
			w.ReportMD,
		).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := repo.Create(w)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestWeeklyReportRepository_List(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	repo := NewWeeklyReportRepository(db)

	rows := sqlmock.NewRows([]string{
		"id", "plan_id", "week_start", "week_end", "report_md", "generated_at",
	}).AddRow(
		"1", "plan1", time.Now(), time.Now(), "# Report 1", time.Now(),
	).AddRow(
		"2", "plan2", time.Now(), time.Now(), "# Report 2", time.Now(),
	)

	mock.ExpectQuery("SELECT \\* FROM weekly_reports").
		WillReturnRows(rows)

	result, err := repo.List()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(result) != 2 {
		t.Fatalf("expected 2 reports, got %d", len(result))
	}
}

func TestWeeklyReportRepository_FindByID(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	repo := NewWeeklyReportRepository(db)

	row := sqlmock.NewRows([]string{
		"id", "plan_id", "week_start", "week_end", "report_md", "generated_at",
	}).AddRow(
		"abc123", "plan1", time.Now(), time.Now(), "# Report", time.Now(),
	)

	mock.ExpectQuery("SELECT \\* FROM weekly_reports WHERE id = \\?").
		WithArgs("abc123").
		WillReturnRows(row)

	report, err := repo.FindByID("abc123")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if report.ID != "abc123" {
		t.Fatalf("expected ID abc123, got %s", report.ID)
	}
}

func TestWeeklyReportRepository_Delete(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	repo := NewWeeklyReportRepository(db)

	mock.ExpectExec("DELETE FROM weekly_reports WHERE id = \\?").
		WithArgs("abc123").
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := repo.Delete("abc123")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}
