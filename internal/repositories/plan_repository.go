package repositories

import (
	"database/sql"
	"fmt"
	"github.com/matoous/go-nanoid/v2"
	"github.com/mobml/ant/internal/models"
	"time"
)

type PlanRepository interface {
	Create(plan *models.Plan) error
	List() ([]*models.Plan, error)
	FindByID(id string) (*models.Plan, error)
	Update(p *models.Plan) error
	Delete(id string) error
}

type planRepository struct {
	db *sql.DB
}

func NewPlanRepository(db *sql.DB) PlanRepository {
	return &planRepository{db: db}
}

func (r *planRepository) Create(p *models.Plan) error {
	query := `
		INSERT INTO plans (id, name, description, start_date, duration)
		VALUES (?, ?, ?, ?, ?)
	`
	id, err := gonanoid.New(8)

	if err != nil {
		return fmt.Errorf("failed to create id")
	}
	_, err = r.db.Exec(
		query,
		id,
		p.Name,
		p.Description,
		p.StartDate,
		p.Duration,
	)

	if err != nil {
		return fmt.Errorf("failed to create plan: %w", err)
	}
	return nil
}

func (r *planRepository) List() ([]*models.Plan, error) {
	query := "SELECT * FROM plans;"

	rows, err := r.db.Query(query)

	if err != nil {
		return nil, fmt.Errorf("failed to list plans: %w", err)
	}
	defer rows.Close()

	var plans []*models.Plan

	for rows.Next() {
		var p models.Plan

		if err := rows.Scan(
			&p.ID,
			&p.Name,
			&p.Description,
			&p.StartDate,
			&p.Duration,
			&p.CreatedAt,
			&p.UpdatedAt,
		); err != nil {
			return nil, fmt.Errorf("failed to reads rows: %w", err)
		}
		plans = append(plans, &p)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("failed during the rows interations: %w", err)
	}

	return plans, nil
}

func (r *planRepository) FindByID(id string) (*models.Plan, error) {
	query := "SELECT * from plans WHERE id = ?"

	row := r.db.QueryRow(query, id)

	var p models.Plan

	if err := row.Scan(
		&p.ID,
		&p.Name,
		&p.Description,
		&p.StartDate,
		&p.Duration,
		&p.CreatedAt,
		&p.UpdatedAt,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, fmt.Errorf("failed to find plan: %w", err)
	}

	if err := row.Err(); err != nil {
		return nil, fmt.Errorf("failed during the row interations: %w", err)
	}

	return &p, nil
}

func (r *planRepository) Update(p *models.Plan) error {
	query := `
		UPDATE plans 
		SET name = ?, description = ?, start_date = ?, duration = ?, updated_at = ? 
		WHERE id = ?
	`

	_, err := r.db.Exec(
		query,
		p.Name,
		p.Description,
		p.StartDate,
		p.Duration,
		time.Now(),
		p.ID,
	)

	if err != nil {
		return fmt.Errorf("failed updating plan '%s': %w", p.ID, err)
	}
	return nil
}

func (r *planRepository) Delete(id string) error {
	query := `DELETE FROM plans WHERE id = ?`

	_, err := r.db.Exec(query, id)

	if err != nil {
		return fmt.Errorf("failed deleting plan '%s': %w", id, err)
	}

	return nil
}
