package usecase

import (
	"context"

	"github.com/LorezV/gophkeeper/internal/entity"
	"github.com/google/uuid"
)

func (uc *GophKeeperUseCase) GetNotes(ctx context.Context, user entity.User) ([]entity.SecretNote, error) {
	return uc.repo.GetNotes(ctx, user)
}

func (uc *GophKeeperUseCase) AddNote(ctx context.Context, note *entity.SecretNote, userID uuid.UUID) error {
	return uc.repo.AddNote(ctx, note, userID)
}

func (uc *GophKeeperUseCase) DelNote(ctx context.Context, noteID, userID uuid.UUID) error {
	return uc.repo.DelNote(ctx, noteID, userID)
}

func (uc *GophKeeperUseCase) UpdateNote(ctx context.Context, note *entity.SecretNote, userID uuid.UUID) error {
	return uc.repo.UpdateNote(ctx, note, userID)
}
