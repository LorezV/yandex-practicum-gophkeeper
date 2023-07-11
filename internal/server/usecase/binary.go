package usecase

import (
	"context"
	"fmt"
	"mime/multipart"
	"os"

	"github.com/LorezV/gophkeeper/internal/entity"
	"github.com/LorezV/gophkeeper/internal/utils"
	"github.com/google/uuid"
)

func (uc *GophKeeperUseCase) GetBinaries(ctx context.Context, user entity.User) ([]entity.Binary, error) {
	return uc.repo.GetBinaries(ctx, user)
}

func (uc *GophKeeperUseCase) AddBinary(
	ctx context.Context,
	binary *entity.Binary,
	file *multipart.FileHeader,
	userID uuid.UUID,
) error {
	userDirectory := uc.cfg.FilesStorage.Location + "/" + userID.String()

	if err := uc.repo.AddBinary(ctx, binary, userID); err != nil {
		return err
	}

	if err := utils.SaveUploadedFile(file, binary.ID.String(), userDirectory); err != nil {
		uc.l.Debug("GophKeeperUseCase - AddBinary - SaveUploadedFile - %w", err)

		return err
	}

	return nil
}

func (uc *GophKeeperUseCase) GetUserBinary(
	ctx context.Context,
	currentUser *entity.User,
	binaryUUID uuid.UUID,
) (string, error) {
	binary, err := uc.repo.GetBinary(ctx, binaryUUID, currentUser.ID)
	if err != nil {
		return "", err
	}
	filePath := fmt.Sprintf(
		"%s/%s/%s",
		uc.cfg.FilesStorage.Location,
		currentUser.ID.String(),
		binary.ID)

	return filePath, nil
}

func (uc *GophKeeperUseCase) DelUserBinary(
	ctx context.Context,
	currentUser *entity.User,
	binaryUUID uuid.UUID,
) error {
	err := uc.repo.DelUserBinary(ctx, currentUser, binaryUUID)
	if err != nil {
		return err
	}

	filePath := fmt.Sprintf(
		"%s/%s/%s",
		uc.cfg.FilesStorage.Location,
		currentUser.ID.String(),
		binaryUUID.String())

	return os.Remove(filePath)
}

func (uc *GophKeeperUseCase) AddBinaryMeta(
	ctx context.Context,
	currentUser *entity.User,
	binaryUUID uuid.UUID,
	meta []entity.Meta,
) (*entity.Binary, error) {
	return uc.repo.AddBinaryMeta(
		ctx,
		currentUser,
		binaryUUID,
		meta,
	)
}
