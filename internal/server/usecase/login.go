package usecase

import (
	"context"

	"github.com/LorezV/gophkeeper/internal/entity"
	"github.com/google/uuid"
)

func (uc *GophKeeperUseCase) GetLogins(ctx context.Context, user entity.User) ([]entity.Login, error) {
	return uc.repo.GetLogins(ctx, user)
}

func (uc *GophKeeperUseCase) AddLogin(ctx context.Context, login *entity.Login, userID uuid.UUID) error {
	return uc.repo.AddLogin(ctx, login, userID)
}

func (uc *GophKeeperUseCase) DelLogin(ctx context.Context, loginID, userID uuid.UUID) error {
	return uc.repo.DelLogin(ctx, loginID, userID)
}

func (uc *GophKeeperUseCase) UpdateLogin(ctx context.Context, login *entity.Login, userID uuid.UUID) error {
	return uc.repo.UpdateLogin(ctx, login, userID)
}
