package repositories

import (
	"database/sql"
	"fmt"

	"github.com/matoous/go-nanoid/v2"
	"github.com/mobml/ant/internal/models"
)

type HabitLogRepository interface {
	Create(h *models.HabitLog) error
	List() ([]*models.HabitLog, error)
	FindByID(id string) (*models.HabitLog, error)
	FindByHabitID(habitID string) ([]*models.HabitLog, error)
	Update(h *models.HabitLog) error
	Delete(id string) error
}

type habitLogRepository struct {
	db *sql.DB
}

func NewHabitLogRepository(db *sql.DB) HabitLogRepository {
	return &habitLogRepository{db: db}
}

func (r *habitLogRepository) Create(h *models.HabitLog) error {
	query := `
        INSERT INTO habit_logs (id, habit_id, log_date, value, note, created_at, updated_at)
        VALUES (?, ?, ?, ?, ?, ?, ?)
    `
	id, err := gonanoid.New(8)

	if err != nil {
		return fmt.Errorf("failed to create id")
	}

	_, err = r.db.Exec(
		query,
		id,
		h.HabitID,
		h.LogDate,
		h.Value,
		h.Note,
		h.CreatedAt,
		h.UpdatedAt,
	)

	if err != nil {
		return fmt.Errorf("failed to create habit log: %w", err)
	}

	return nil
}

func (r *habitLogRepository) List() ([]*models.HabitLog, error) {
	query := "SELECT * FROM habit_logs"

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to list habit logs: %w", err)
	}
	defer rows.Close()

	var logs []*models.HabitLog

	for rows.Next() {
		var h models.HabitLog

		if err := rows.Scan(
			&h.ID,
			&h.HabitID,
			&h.LogDate,
			&h.Value,
			&h.Note,
			&h.CreatedAt,
			&h.UpdatedAt,
		); err != nil {
			return nil, fmt.Errorf("failed to scan rows: %w", err)
		}

		logs = append(logs, &h)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error during row iteration: %w", err)
	}

	return logs, nil
}

func (r *habitLogRepository) FindByID(id string) (*models.HabitLog, error) {
	query := "SELECT * FROM habit_logs WHERE id = ?"

	row := r.db.QueryRow(query, id)

	var h models.HabitLog

	if err := row.Scan(
		&h.ID,
		&h.HabitID,
		&h.LogDate,
		&h.Value,
		&h.Note,
		&h.CreatedAt,
		&h.UpdatedAt,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to scan habit log: %w", err)
	}

	if err := row.Err(); err != nil {
		return nil, fmt.Errorf("error during row iteration: %w", err)
	}

	return &h, nil
}

func (r *habitLogRepository) FindByHabitID(habitID string) ([]*models.HabitLog, error) {
	query := "SELECT * FROM habit_logs WHERE habit_id = ?"

	rows, err := r.db.Query(query, habitID)
	if err != nil {
		return nil, fmt.Errorf("failed to list logs by habit ID: %w", err)
	}
	defer rows.Close()

	var logs []*models.HabitLog

	for rows.Next() {
		var h models.HabitLog

		if err := rows.Scan(
			&h.ID,
			&h.HabitID,
			&h.LogDate,
			&h.Value,
			&h.Note,
			&h.CreatedAt,
			&h.UpdatedAt,
		); err != nil {
			return nil, fmt.Errorf("failed to scan logs: %w", err)
		}

		logs = append(logs, &h)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error during row iteration: %w", err)
	}

	return logs, nil
}

func (r *habitLogRepository) Update(h *models.HabitLog) error {
	query := `
        UPDATE habit_logs
        SET value = ?, note = ?, updated_at = ?
        WHERE id = ?
    `

	_, err := r.db.Exec(
		query,
		h.Value,
		h.Note,
		h.UpdatedAt,
		h.ID,
	)

	if err != nil {
		return fmt.Errorf("failed to update habit log '%s': %w", h.ID, err)
	}

	return nil
}

func (r *habitLogRepository) Delete(id string) error {
	query := `DELETE FROM habit_logs WHERE id = ?`

	_, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete habit log '%s': %w", id, err)
	}

	return nil
}
