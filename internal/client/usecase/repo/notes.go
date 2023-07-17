package repo

import (
	"errors"

	"github.com/LorezV/gophkeeper/internal/client/usecase/repo/models"
	"github.com/LorezV/gophkeeper/internal/client/usecase/viewsets"
	"github.com/LorezV/gophkeeper/internal/entity"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

var errNoteNotFound = errors.New("note not found")

func (r *GophKeeperRepo) AddNote(note *entity.SecretNote) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		noteForSaving := models.Note{
			ID:     note.ID,
			Name:   note.Name,
			Note:   note.Note,
			UserID: r.getUserID(),
		}
		if err := tx.Save(&noteForSaving).Error; err != nil {
			return err
		}
		for _, meta := range note.Meta {
			metaForLogin := models.MetaNote{
				Name:   meta.Name,
				Value:  meta.Value,
				NoteID: noteForSaving.ID,
				ID:     meta.ID,
			}
			if err := tx.Create(&metaForLogin).Error; err != nil {
				return err
			}
		}

		return nil
	})
}

func (r *GophKeeperRepo) LoadNotes() []viewsets.NoteForList {
	userID := r.getUserID()
	var notes []models.Note
	r.db.
		Model(&models.Note{}).
		Preload("Meta").
		Where("user_id", userID).Find(&notes)
	if len(notes) == 0 {
		return nil
	}

	notesViewSet := make([]viewsets.NoteForList, len(notes))

	for index := range notes {
		notesViewSet[index].ID = notes[index].ID
		notesViewSet[index].Name = notes[index].Name
	}

	return notesViewSet
}

func (r *GophKeeperRepo) SaveNotes(notes []entity.SecretNote) error {
	if len(notes) == 0 {
		return nil
	}
	userID := r.getUserID()
	notesForDB := make([]models.Note, len(notes))
	for index := range notes {
		notesForDB[index].ID = notes[index].ID
		notesForDB[index].Name = notes[index].Name
		notesForDB[index].Note = notes[index].Note
		notesForDB[index].UserID = userID
	}

	return r.db.Save(notesForDB).Error
}

func (r *GophKeeperRepo) GetNoteByID(noteID uuid.UUID) (note entity.SecretNote, err error) {
	var noteFromDB models.Note
	if err = r.db.
		Model(&models.Note{}).
		Preload("Meta").
		Find(&noteFromDB, noteID).Error; noteFromDB.ID == uuid.Nil || err != nil {
		return note, errNoteNotFound
	}

	note.ID = noteFromDB.ID
	note.Note = noteFromDB.Note
	note.Name = noteFromDB.Name
	for index := range noteFromDB.Meta {
		note.Meta = append(
			note.Meta,
			entity.Meta{
				ID:    noteFromDB.Meta[index].ID,
				Name:  noteFromDB.Meta[index].Name,
				Value: noteFromDB.Meta[index].Value,
			})
	}

	return
}

func (r *GophKeeperRepo) DelNote(noteID uuid.UUID) error {
	return r.db.Unscoped().Delete(&models.Note{}, noteID).Error
}
