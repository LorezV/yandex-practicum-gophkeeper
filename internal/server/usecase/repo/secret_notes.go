package repo

import (
	"context"

	"github.com/LorezV/gophkeeper/internal/entity"
)

func (r *GophKeeperRepo) GetSecretNotes(ctx context.Context, user entity.User) (notes []entity.SecretNote, err error) {
	return
}
