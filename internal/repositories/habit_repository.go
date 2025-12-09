package repositories

import (
	"database/sql"
	"fmt"

	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/mobml/ant/internal/models"
)

type HabitRepository interface {
	Create(h *models.Habit) error
	List() ([]*models.Habit, error)
	FindByID(id string) (*models.Habit, error)
	Update(h *models.Habit) error
	Delete(id string) error
}

type habitRepository struct {
	db *sql.DB
}

func NewHabitRepository(db *sql.DB) HabitRepository {
	return &habitRepository{db: db}
}

func (r *habitRepository) Create(h *models.Habit) error {
	query := `
		INSERT INTO habits (id, goal_id, name, description, measure_type, measure_unit, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?)
	`

	id, err := gonanoid.New(8)

	if err != nil {
		return fmt.Errorf("failed to create id")
	}

	_, err = r.db.Exec(
		query,
		id,
		h.GoalID,
		h.Name,
		h.Description,
		h.MeasureType,
		h.MeasureUnit,
		h.CreatedAt,
		h.UpdatedAt,
	)

	if err != nil {
		return fmt.Errorf("failed to create habit: %w", err)
	}

	return nil
}

func (r *habitRepository) List() ([]*models.Habit, error) {
	query := "SELECT * FROM habits;"

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to list habits: %w", err)
	}
	defer rows.Close()

	var habits []*models.Habit

	for rows.Next() {
		var h models.Habit

		if err := rows.Scan(
			&h.ID,
			&h.GoalID,
			&h.Name,
			&h.Description,
			&h.MeasureType,
			&h.MeasureUnit,
			&h.CreatedAt,
			&h.UpdatedAt,
		); err != nil {
			return nil, fmt.Errorf("failed scanning habit: %w", err)
		}

		habits = append(habits, &h)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration failed: %w", err)
	}

	return habits, nil
}

func (r *habitRepository) FindByID(id string) (*models.Habit, error) {
	query := "SELECT * FROM habits WHERE id = ?"

	row := r.db.QueryRow(query, id)

	var h models.Habit

	if err := row.Scan(
		&h.ID,
		&h.GoalID,
		&h.Name,
		&h.Description,
		&h.MeasureType,
		&h.MeasureUnit,
		&h.CreatedAt,
		&h.UpdatedAt,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, fmt.Errorf("failed to find habit: %w", err)
	}

	return &h, nil
}

func (r *habitRepository) Update(h *models.Habit) error {
	query := `
		UPDATE habits
		SET goal_id = ?, name = ?, description = ?, measure_type = ?, measure_unit = ?, updated_at = ?
		WHERE id = ?
	`

	_, err := r.db.Exec(
		query,
		h.GoalID,
		h.Name,
		h.Description,
		h.MeasureType,
		h.MeasureUnit,
		h.UpdatedAt,
		h.ID,
	)

	if err != nil {
		return fmt.Errorf("failed updating habit '%s': %w", h.ID, err)
	}

	return nil
}

func (r *habitRepository) Delete(id string) error {
	query := "DELETE FROM habits WHERE id = ?"

	_, err := r.db.Exec(query, id)

	if err != nil {
		return fmt.Errorf("failed deleting habit '%s': %w", id, err)
	}

	return nil
}
