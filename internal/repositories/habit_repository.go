package repositories

import (
	"database/sql"
	"fmt"
	"time"

	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/mobml/ant/internal/models"
)

type HabitRepository interface {
	Create(h *models.Habit) error
	CreateHabit(h *models.Habit, days []int) error
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
		INSERT INTO habits (id, goal_id, name, description, measure_type, measure_unit)
		VALUES (?, ?, ?, ?, ?, ?)
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
	)

	if err != nil {
		return fmt.Errorf("failed to create habit: %w", err)
	}

	return nil
}

func (r *habitRepository) CreateHabit(h *models.Habit, days []int) error {
	tx, err := r.db.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}

	defer tx.Rollback()

	id, err := gonanoid.New(8)
	if err != nil {
		return fmt.Errorf("failed to create id")
	}
	h.ID = id

	query := `
		INSERT INTO habits (id, goal_id, name, description, measure_type, measure_unit)
		VALUES (?, ?, ?, ?, ?, ?)
	`

	_, err = tx.Exec(
		query,
		h.ID,
		h.GoalID,
		h.Name,
		h.Description,
		h.MeasureType,
		h.MeasureUnit,
	)

	if err != nil {
		return fmt.Errorf("failed to create habit: %w", err)
	}

	stmt, err := tx.Prepare(`
		INSERT INTO habit_schedules (id, habit_id, day_of_week)
		VALUES (?, ?, ?)
	`)
	if err != nil {
		return fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close()

	for _, day := range days {
		sID, err := gonanoid.New(8)
		if err != nil {
			return fmt.Errorf("failed to generate schedule id: %w", err)
		}

		if _, err = stmt.Exec(sID, h.ID, day); err != nil {
			return fmt.Errorf("failed to insert schedule for day %d: %w", day, err)
		}
	}

	return tx.Commit()
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
		time.Now(),
		h.ID,
	)

	if err != nil {
		return fmt.Errorf("failed updating habit '%s': %w", h.ID, err)
	}

	return nil
}

//make this code transactional to delete associated habit schedules as well

func (r *habitRepository) Delete(id string) error {
	tx, err := r.db.Begin()
	if err != nil {
		return fmt.Errorf("begin tx: %w", err)
	}
	defer tx.Rollback()

	deletes := []struct {
		query string
		label string
	}{
		{"DELETE FROM habit_schedules WHERE habit_id = ?", "habit_schedules"},
		{"DELETE FROM habit_logs WHERE habit_id = ?", "habit_logs"},
		{"DELETE FROM habits WHERE id = ?", "habits"},
	}

	for _, d := range deletes {
		_, err := tx.Exec(d.query, id)
		if err != nil {
			return fmt.Errorf("delete from %s failed: %w", d.label, err)
		}
	}

	return tx.Commit()
}
