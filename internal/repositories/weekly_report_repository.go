package repositories

import (
	"database/sql"
	"fmt"
	"github.com/matoous/go-nanoid/v2"
	"github.com/mobml/ant/internal/models"
)

type WeeklyReportRepository interface {
	Create(r *models.WeeklyReport) error
	List() ([]*models.WeeklyReport, error)
	FindByID(id string) (*models.WeeklyReport, error)
	Delete(id string) error
}

type weeklyReportRepository struct {
	db *sql.DB
}

func NewWeeklyReportRepository(db *sql.DB) WeeklyReportRepository {
	return &weeklyReportRepository{db: db}
}

func (r *weeklyReportRepository) Create(w *models.WeeklyReport) error {
	query := `
		INSERT INTO weekly_reports (id, plan_id, week_start, week_end, report_md)
		VALUES (?, ?, ?, ?, ?, ?)
	`

	id, err := gonanoid.New(8)
	if err != nil {
		return fmt.Errorf("failed to generate id")
	}

	_, err = r.db.Exec(
		query,
		id,
		w.PlanID,
		w.WeekStart,
		w.WeekEnd,
		w.ReportMD,
	)

	if err != nil {
		return fmt.Errorf("failed to create weekly report: %w", err)
	}

	return nil
}

func (r *weeklyReportRepository) List() ([]*models.WeeklyReport, error) {
	query := "SELECT * FROM weekly_reports"

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to list weekly reports: %w", err)
	}
	defer rows.Close()

	var reports []*models.WeeklyReport

	for rows.Next() {
		var w models.WeeklyReport

		if err := rows.Scan(
			&w.ID,
			&w.PlanID,
			&w.WeekStart,
			&w.WeekEnd,
			&w.ReportMD,
			&w.GeneratedAt,
		); err != nil {
			return nil, fmt.Errorf("failed reading rows: %w", err)
		}

		reports = append(reports, &w)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("failed during rows iteration: %w", err)
	}

	return reports, nil
}

func (r *weeklyReportRepository) FindByID(id string) (*models.WeeklyReport, error) {
	query := "SELECT * FROM weekly_reports WHERE id = ?"

	row := r.db.QueryRow(query, id)

	var w models.WeeklyReport

	if err := row.Scan(
		&w.ID,
		&w.PlanID,
		&w.WeekStart,
		&w.WeekEnd,
		&w.ReportMD,
		&w.GeneratedAt,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to find weekly report: %w", err)
	}

	return &w, nil
}

func (r *weeklyReportRepository) Delete(id string) error {
	query := `DELETE FROM weekly_reports WHERE id = ?`

	_, err := r.db.Exec(query, id)

	if err != nil {
		return fmt.Errorf("failed deleting weekly report '%s': %w", id, err)
	}

	return nil
}
