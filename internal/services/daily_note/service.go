package dailynote

import (
	dc "github.com/mobml/ant/internal/domain/common"
	dn "github.com/mobml/ant/internal/domain/daily_note"
	"github.com/mobml/ant/internal/models"
	repo "github.com/mobml/ant/internal/repositories"
)

type DailyNoteService interface {
	CreateDailyNote(dailyNote *models.DailyNote) error
	ListDailyNotes() ([]*models.DailyNote, error)
	UpdateDailyNote(dailyNote *models.DailyNote) error
	DeleteDailyNote(id string) error
}

type dailyNoteService struct {
	dailyNoteRepo repo.DailyNoteRepository
}

func NewDailyNoteService(dailyNoteRepo repo.DailyNoteRepository) DailyNoteService {
	return &dailyNoteService{
		dailyNoteRepo: dailyNoteRepo,
	}
}

func (dns *dailyNoteService) CreateDailyNote(dailyNote *models.DailyNote) error {

	if err := dn.ValidateDailyNote(dailyNote); err != nil {
		return err
	}

	return dns.dailyNoteRepo.Create(dailyNote)
}

func (dns *dailyNoteService) ListDailyNotes() ([]*models.DailyNote, error) {
	return dns.dailyNoteRepo.List()
}

func (dns *dailyNoteService) UpdateDailyNote(dailyNote *models.DailyNote) error {

	if err := dn.ValidateDailyNote(dailyNote); err != nil {
		return err
	}

	if err := dc.ValidateID(dailyNote.ID); err != nil {
		return err
	}

	return dns.dailyNoteRepo.Update(dailyNote)
}

func (dns *dailyNoteService) DeleteDailyNote(id string) error {

	if err := dc.ValidateID(id); err != nil {
		return err
	}
	return dns.dailyNoteRepo.Delete(id)
}
