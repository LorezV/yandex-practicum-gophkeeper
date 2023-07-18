package repo

import (
	"context"
	"errors"

	"github.com/LorezV/gophkeeper/internal/entity"
	"github.com/LorezV/gophkeeper/internal/server/usecase/repo/models"
	"github.com/google/uuid"
)

var errWrongBinaryOwner = errors.New("wrong binary owner or not found")

func (r *GophKeeperRepo) GetBinaries(ctx context.Context, user entity.User) ([]entity.Binary, error) {
	var binariesFromDB []models.Binary

	if err := r.db.WithContext(ctx).
		Model(&models.Binary{}).
		Preload("Meta").
		Find(&binariesFromDB, "user_id = ?", user.ID).Error; err != nil {
		return nil, err
	}

	if len(binariesFromDB) == 0 {
		return nil, nil
	}

	binaries := make([]entity.Binary, len(binariesFromDB))

	for index := range binariesFromDB {
		binaries[index].ID = binariesFromDB[index].ID
		binaries[index].Name = binariesFromDB[index].Name
		binaries[index].FileName = binariesFromDB[index].FileName
		for metaIndex := range binariesFromDB[index].Meta {
			binaries[index].Meta = append(binaries[index].Meta, entity.Meta{
				ID:    binariesFromDB[index].Meta[metaIndex].ID,
				Name:  binariesFromDB[index].Meta[metaIndex].Name,
				Value: binariesFromDB[index].Meta[metaIndex].Value,
			})
		}
	}

	return binaries, nil
}

func (r *GophKeeperRepo) AddBinary(ctx context.Context, binary *entity.Binary, userID uuid.UUID) error {
	newBinaryToDB := models.Binary{
		Name:     binary.Name,
		FileName: binary.FileName,
		UserID:   userID,
	}

	if err := r.db.WithContext(ctx).Create(&newBinaryToDB).Error; err != nil {
		r.l.Debug("GophKeeperRepo - AddBinary - Create - %w", err)

		return err
	}
	binary.ID = newBinaryToDB.ID

	return nil
}

func (r *GophKeeperRepo) GetBinary(ctx context.Context, binaryID, userID uuid.UUID) (*entity.Binary, error) {
	var binaryFromDB models.Binary
	if err := r.db.WithContext(ctx).
		Model(&models.Binary{}).
		Preload("Meta").
		Find(&binaryFromDB, binaryID).Error; err != nil {
		return nil, err
	}

	if binaryFromDB.UserID != userID {
		return nil, errWrongBinaryOwner
	}
	var meta []entity.Meta
	if len(binaryFromDB.Meta) > 0 {
		meta = make([]entity.Meta, len(binaryFromDB.Meta))
		for index := range binaryFromDB.Meta {
			meta[index].ID = binaryFromDB.Meta[index].ID
			meta[index].Name = binaryFromDB.Meta[index].Name
			meta[index].Value = binaryFromDB.Meta[index].Value
		}
	}

	return &entity.Binary{
		ID:       binaryFromDB.ID,
		FileName: binaryFromDB.FileName,
		Meta:     meta,
	}, nil
}

func (r *GophKeeperRepo) DelUserBinary(ctx context.Context, currentUser *entity.User, binaryUUID uuid.UUID) error {
	var binaryFromDB models.Binary
	r.db.WithContext(ctx).Find(&binaryFromDB, binaryUUID)
	if binaryFromDB.UserID != currentUser.ID {
		return errWrongBinaryOwner
	}

	return r.db.Delete(&binaryFromDB).Error
}

func (r *GophKeeperRepo) AddBinaryMeta(
	ctx context.Context,
	currentUser *entity.User,
	binaryUUID uuid.UUID,
	meta []entity.Meta,
) (*entity.Binary, error) {
	metaForDB := make([]models.MetaBinary, len(meta))
	for index := range meta {
		metaForDB[index].BinaryID = binaryUUID
		metaForDB[index].Name = meta[index].Name
		metaForDB[index].Value = meta[index].Value
	}

	if err := r.db.WithContext(ctx).Save(&metaForDB).Error; err != nil {
		return nil, err
	}

	return r.GetBinary(ctx, binaryUUID, currentUser.ID)
}
