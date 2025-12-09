package repositories

import (
	"database/sql"
	"fmt"
	"github.com/matoous/go-nanoid/v2"
	"github.com/mobml/ant/internal/models"
	"time"
)

type DailyNoteRepository interface {
	Create(note *models.DailyNote) error
	List() ([]*models.DailyNote, error)
	FindByID(id string) (*models.DailyNote, error)
	Update(note *models.DailyNote) error
	Delete(id string) error
}

type dailyNoteRepository struct {
	db *sql.DB
}

func NewDailyNoteRepository(db *sql.DB) DailyNoteRepository {
	return &dailyNoteRepository{db: db}
}

func (r *dailyNoteRepository) Create(n *models.DailyNote) error {
	query := `
		INSERT INTO daily_notes (id, note_date, content)
		VALUES (?, ?, ?)
	`

	id, err := gonanoid.New(8)
	if err != nil {
		return fmt.Errorf("failed generating id: %w", err)
	}

	_, err = r.db.Exec(query,
		id,
		n.NoteDate,
		n.Content,
	)

	if err != nil {
		return fmt.Errorf("failed creating daily note: %w", err)
	}

	return nil
}

func (r *dailyNoteRepository) List() ([]*models.DailyNote, error) {
	query := "SELECT * FROM daily_notes"

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed listing notes: %w", err)
	}
	defer rows.Close()

	var notes []*models.DailyNote

	for rows.Next() {
		var n models.DailyNote

		if err := rows.Scan(
			&n.ID,
			&n.NoteDate,
			&n.Content,
			&n.CreatedAt,
			&n.UpdatedAt,
		); err != nil {
			return nil, fmt.Errorf("failed scanning rows: %w", err)
		}
		notes = append(notes, &n)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("failed during rows iteration: %w", err)
	}

	return notes, nil
}

func (r *dailyNoteRepository) FindByID(id string) (*models.DailyNote, error) {
	query := "SELECT * FROM daily_notes WHERE id = ?"

	row := r.db.QueryRow(query, id)

	var n models.DailyNote

	if err := row.Scan(
		&n.ID,
		&n.NoteDate,
		&n.Content,
		&n.CreatedAt,
		&n.UpdatedAt,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("failed reading daily note: %w", err)
	}

	return &n, nil
}

func (r *dailyNoteRepository) Update(n *models.DailyNote) error {
	query := `
		UPDATE daily_notes
		SET note_date = ?, content = ?, updated_at = ?
		WHERE id = ?
	`

	_, err := r.db.Exec(
		query,
		n.NoteDate,
		n.Content,
		time.Now(),
		n.ID,
	)

	if err != nil {
		return fmt.Errorf("failed updating note '%s': %w", n.ID, err)
	}

	return nil
}

func (r *dailyNoteRepository) Delete(id string) error {
	query := "DELETE FROM daily_notes WHERE id = ?"

	_, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed deleting note '%s': %w", id, err)
	}

	return nil
}
