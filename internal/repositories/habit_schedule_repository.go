package repositories

import (
	"database/sql"
	"fmt"
	"time"

	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/mobml/ant/internal/models"
)

type HabitScheduleRepository interface {
	Create(hs *models.HabitSchedule) error
	List() ([]*models.HabitSchedule, error)
	FindByID(id string) (*models.HabitSchedule, error)
	Update(hs *models.HabitSchedule) error
	Delete(id string) error
}

type habitScheduleRepository struct {
	db *sql.DB
}

func NewHabitScheduleRepository(db *sql.DB) HabitScheduleRepository {
	return &habitScheduleRepository{db: db}
}

func (r *habitScheduleRepository) Create(hs *models.HabitSchedule) error {
	query := `
		INSERT INTO habit_schedules (id, habit_id, day_of_week)
		VALUES (?, ?, ?)
	`

	id, err := gonanoid.New(8)
	if err != nil {
		return fmt.Errorf("failed generating id: %w", err)
	}

	_, err = r.db.Exec(
		query,
		id,
		hs.HabitID,
		hs.DayOfWeek,
	)

	if err != nil {
		return fmt.Errorf("failed creating habit schedule: %w", err)
	}

	return nil
}

func (r *habitScheduleRepository) List() ([]*models.HabitSchedule, error) {
	query := "SELECT * FROM habit_schedules;"

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to list habit schedules: %w", err)
	}
	defer rows.Close()

	var schedules []*models.HabitSchedule

	for rows.Next() {
		var hs models.HabitSchedule

		if err := rows.Scan(
			&hs.ID,
			&hs.HabitID,
			&hs.DayOfWeek,
			&hs.CreatedAt,
			&hs.UpdatedAt,
		); err != nil {
			return nil, fmt.Errorf("failed scanning rows: %w", err)
		}

		schedules = append(schedules, &hs)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("failed during row iteration: %w", err)
	}

	return schedules, nil
}

func (r *habitScheduleRepository) FindByID(id string) (*models.HabitSchedule, error) {
	query := "SELECT * FROM habit_schedules WHERE id = ?"

	row := r.db.QueryRow(query, id)

	var hs models.HabitSchedule

	if err := row.Scan(
		&hs.ID,
		&hs.HabitID,
		&hs.DayOfWeek,
		&hs.CreatedAt,
		&hs.UpdatedAt,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to find habit schedule: %w", err)
	}

	return &hs, nil
}

func (r *habitScheduleRepository) Update(hs *models.HabitSchedule) error {
	query := `
		UPDATE habit_schedules
		SET day_of_week = ?, updated_at = ?
		WHERE id = ?
	`

	_, err := r.db.Exec(
		query,
		hs.DayOfWeek,
		time.Now(),
		hs.ID,
	)

	if err != nil {
		return fmt.Errorf("failed updating habit schedule '%s': %w", hs.ID, err)
	}

	return nil
}

func (r *habitScheduleRepository) Delete(id string) error {
	query := "DELETE FROM habit_schedules WHERE id = ?"

	_, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed deleting habit schedule '%s': %w", id, err)
	}

	return nil
}
