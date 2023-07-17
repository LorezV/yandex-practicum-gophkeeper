package usecase

import (
	"context"

	"github.com/LorezV/gophkeeper/internal/entity"
	"github.com/google/uuid"
)

func (uc *GophKeeperUseCase) GetCards(ctx context.Context, user entity.User) ([]entity.Card, error) {
	return uc.repo.GetCards(ctx, user)
}

func (uc *GophKeeperUseCase) AddCard(ctx context.Context, card *entity.Card, userID uuid.UUID) error {
	return uc.repo.AddCard(ctx, card, userID)
}

func (uc *GophKeeperUseCase) DelCard(ctx context.Context, cardUUID, userID uuid.UUID) error {
	return uc.repo.DelCard(ctx, cardUUID, userID)
}

func (uc *GophKeeperUseCase) UpdateCard(ctx context.Context, card *entity.Card, userID uuid.UUID) error {
	return uc.repo.UpdateCard(ctx, card, userID)
}
