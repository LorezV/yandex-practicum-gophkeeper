package repo

import (
	"github.com/LorezV/gophkeeper/internal/client/usecase/repo/models"
	"github.com/LorezV/gophkeeper/internal/client/usecase/viewsets"
	"github.com/LorezV/gophkeeper/internal/entity"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (r *GophKeeperRepo) LoadBinaries() []viewsets.BinaryForList {
	userID := r.getUserID()
	var binaries []models.Binary
	r.db.
		Model(&models.Binary{}).
		Preload("Meta").
		Where("user_id", userID).Find(&binaries)
	if len(binaries) == 0 {
		return nil
	}

	binariesViewSet := make([]viewsets.BinaryForList, len(binaries))

	for index := range binaries {
		binariesViewSet[index].ID = binaries[index].ID
		binariesViewSet[index].Name = binaries[index].Name
		binariesViewSet[index].FileName = binaries[index].FileName
	}

	return binariesViewSet
}

func (r *GophKeeperRepo) SaveBinaries(binaries []entity.Binary) error {
	if len(binaries) == 0 {
		return nil
	}
	userID := r.getUserID()
	binariesForDB := make([]models.Binary, len(binaries))
	for index := range binaries {
		binariesForDB[index].ID = binaries[index].ID
		binariesForDB[index].Name = binaries[index].Name
		binariesForDB[index].FileName = binaries[index].FileName
		binariesForDB[index].UserID = userID
		for _, meta := range binaries[index].Meta {
			binariesForDB[index].Meta = append(binariesForDB[index].Meta,
				models.MetaBinary{
					ID:    meta.ID,
					Name:  meta.Name,
					Value: meta.Value,
				})
		}
	}

	return r.db.Save(binariesForDB).Error
}

func (r *GophKeeperRepo) AddBinary(binary *entity.Binary) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		binaryForSaving := models.Binary{
			ID:       binary.ID,
			Name:     binary.Name,
			FileName: binary.FileName,
			UserID:   r.getUserID(),
		}
		if err := tx.Save(&binaryForSaving).Error; err != nil {
			return err
		}
		for _, meta := range binary.Meta {
			metaForBinary := models.MetaNote{
				Name:   meta.Name,
				Value:  meta.Value,
				NoteID: binaryForSaving.ID,
				ID:     meta.ID,
			}
			if err := tx.Create(&metaForBinary).Error; err != nil {
				return err
			}
		}

		return nil
	})
}

func (r *GophKeeperRepo) GetBinaryByID(binaryID uuid.UUID) (binary entity.Binary, err error) {
	var binaryFromDB models.Binary
	if err = r.db.
		Model(&models.Binary{}).
		Preload("Meta").
		Find(&binaryFromDB, binaryID).Error; binaryFromDB.ID == uuid.Nil || err != nil {
		return binary, errNoteNotFound
	}

	binary.ID = binaryFromDB.ID
	binary.Name = binaryFromDB.Name
	binary.FileName = binaryFromDB.FileName
	for index := range binaryFromDB.Meta {
		binary.Meta = append(
			binary.Meta,
			entity.Meta{
				ID:    binaryFromDB.Meta[index].ID,
				Name:  binaryFromDB.Meta[index].Name,
				Value: binaryFromDB.Meta[index].Value,
			})
	}

	return binary, nil
}

func (r *GophKeeperRepo) DelBinary(binaryID uuid.UUID) error {
	return r.db.Unscoped().Delete(&models.Note{}, binaryID).Error
}
