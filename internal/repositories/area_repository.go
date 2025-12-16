package repositories

import (
	"database/sql"
	"fmt"
	"time"

	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/mobml/ant/internal/models"
)

type AreaRepository interface {
	Create(area *models.Area) error
	List() ([]*models.Area, error)
	FindByID(id string) (*models.Area, error)
	ListByPlan(planID string) ([]*models.Area, error)
	Update(area *models.Area) error
	Delete(id string) error
}

type areaRepository struct {
	db *sql.DB
}

func NewAreaRepository(db *sql.DB) AreaRepository {
	return &areaRepository{db: db}
}

func (r *areaRepository) Create(a *models.Area) error {
	query := `
		INSERT INTO areas (id, plan_id, name, description)
		VALUES (?, ?, ?, ?)
	`
	id, err := gonanoid.New(8)

	if err != nil {
		return fmt.Errorf("failed to create id")
	}

	_, err = r.db.Exec(
		query,
		id,
		a.PlanID,
		a.Name,
		a.Description,
	)

	if err != nil {
		return fmt.Errorf("failed to create area: %w", err)
	}

	return nil
}

func (r *areaRepository) List() ([]*models.Area, error) {
	query := "SELECT * FROM areas;"

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to list areas: %w", err)
	}
	defer rows.Close()

	var areas []*models.Area

	for rows.Next() {
		var a models.Area

		if err := rows.Scan(
			&a.ID,
			&a.PlanID,
			&a.Name,
			&a.Description,
			&a.CreatedAt,
			&a.UpdatedAt,
		); err != nil {
			return nil, fmt.Errorf("failed to scan area: %w", err)
		}

		areas = append(areas, &a)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}

	return areas, nil
}

func (r *areaRepository) FindByID(id string) (*models.Area, error) {
	query := "SELECT * FROM areas WHERE id = ?"

	row := r.db.QueryRow(query, id)

	var a models.Area

	if err := row.Scan(
		&a.ID,
		&a.PlanID,
		&a.Name,
		&a.Description,
		&a.CreatedAt,
		&a.UpdatedAt,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to find area: %w", err)
	}

	return &a, nil
}

//write a method that lists areas by plan id
func (r *areaRepository) ListByPlan(planID string) ([]*models.Area, error) {
	query := "SELECT * FROM areas WHERE plan_id = ?;"
	
	rows, err := r.db.Query(query, planID)
	if err != nil {
		return nil, fmt.Errorf("failed to list areas by plan: %w", err)
	}
	defer rows.Close()
	
	var areas []*models.Area
	for rows.Next() {
		var a models.Area
		if err := rows.Scan(
			&a.ID,
			&a.PlanID,
			&a.Name,
			&a.Description,
			&a.CreatedAt,
			&a.UpdatedAt,
		); err != nil {
			return nil, fmt.Errorf("failed to scan area: %w", err)
		}
		areas = append(areas, &a)
	}
	
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}
	
	return areas, nil
}

func (r *areaRepository) Update(a *models.Area) error {
	query := `
		UPDATE areas 
		SET plan_id = ?, name = ?, description = ?, updated_at = ?
		WHERE id = ?
	`

	_, err := r.db.Exec(
		query,
		a.PlanID,
		a.Name,
		a.Description,
		time.Now(),
		a.ID,
	)

	if err != nil {
		return fmt.Errorf("failed updating area '%s': %w", a.ID, err)
	}

	return nil
}

func (r *areaRepository) Delete(id string) error {
	query := "DELETE FROM areas WHERE id = ?"

	_, err := r.db.Exec(query, id)

	if err != nil {
		return fmt.Errorf("failed deleting area '%s': %w", id, err)
	}

	return nil
}
