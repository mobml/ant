package dailynote

import (
	"strings"

	"github.com/mobml/ant/internal/models"
)

func ValidateDailyNote(dn *models.DailyNote) error {
	if strings.TrimSpace(dn.Content) == "" {
		return ErrContentRequired
	}
	if dn.NoteDate.IsZero() {
		return ErrNoteDateRequired
	}
	return nil
}
