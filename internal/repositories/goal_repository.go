package repositories

import (
	"database/sql"
	"fmt"

	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/mobml/ant/internal/models"
)

type GoalRepository interface {
	Create(goal *models.Goal) error
	List() ([]*models.Goal, error)
	FindByID(id string) (*models.Goal, error)
	Update(goal *models.Goal) error
	Delete(id string) error
}

type goalRepository struct {
	db *sql.DB
}

func NewGoalRepository(db *sql.DB) GoalRepository {
	return &goalRepository{db: db}
}

func (r *goalRepository) Create(g *models.Goal) error {
	query := `
		INSERT INTO goals (id, area_id, name, description, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?, ?)
	`
	id, err := gonanoid.New(8)

	if err != nil {
		return fmt.Errorf("failed to create id")
	}

	_, err = r.db.Exec(
		query,
		id,
		g.AreaID,
		g.Name,
		g.Description,
		g.CreatedAt,
		g.UpdatedAt,
	)

	if err != nil {
		return fmt.Errorf("failed to create goal: %w", err)
	}

	return nil
}

func (r *goalRepository) List() ([]*models.Goal, error) {
	query := "SELECT * FROM goals;"

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to list goals: %w", err)
	}
	defer rows.Close()

	var goals []*models.Goal

	for rows.Next() {
		var g models.Goal

		if err := rows.Scan(
			&g.ID,
			&g.AreaID,
			&g.Name,
			&g.Description,
			&g.CreatedAt,
			&g.UpdatedAt,
		); err != nil {
			return nil, fmt.Errorf("failed scanning goal: %w", err)
		}

		goals = append(goals, &g)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration failed: %w", err)
	}

	return goals, nil
}

func (r *goalRepository) FindByID(id string) (*models.Goal, error) {
	query := "SELECT * FROM goals WHERE id = ?"

	row := r.db.QueryRow(query, id)

	var g models.Goal

	if err := row.Scan(
		&g.ID,
		&g.AreaID,
		&g.Name,
		&g.Description,
		&g.CreatedAt,
		&g.UpdatedAt,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to find goal: %w", err)
	}

	return &g, nil
}

func (r *goalRepository) Update(g *models.Goal) error {
	query := `
		UPDATE goals 
		SET area_id = ?, name = ?, description = ?, updated_at = ?
		WHERE id = ?
	`

	_, err := r.db.Exec(
		query,
		g.AreaID,
		g.Name,
		g.Description,
		g.UpdatedAt,
		g.ID,
	)

	if err != nil {
		return fmt.Errorf("failed updating goal '%s': %w", g.ID, err)
	}

	return nil
}

func (r *goalRepository) Delete(id string) error {
	query := "DELETE FROM goals WHERE id = ?"

	_, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed deleting goal '%s': %w", id, err)
	}

	return nil
}
