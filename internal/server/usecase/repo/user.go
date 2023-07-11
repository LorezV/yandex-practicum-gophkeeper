package repo

import (
	"context"
	"fmt"

	"github.com/LorezV/gophkeeper/internal/entity"
	"github.com/LorezV/gophkeeper/internal/server/usecase/repo/models"
	"github.com/LorezV/gophkeeper/internal/utils"
	"github.com/LorezV/gophkeeper/internal/utils/errs"
	"github.com/google/uuid"
)

func (r *GophKeeperRepo) GetUserByEmail(ctx context.Context, email, hashedPassword string) (user entity.User, err error) {
	var userFromDB models.User

	r.db.WithContext(ctx).Where("email = ?", email).First(&userFromDB)

	if userFromDB.ID == uuid.Nil {
		err = errs.ErrWrongCredentials

		return
	}

	if err = utils.VerifyPassword(userFromDB.Password, hashedPassword); err != nil {
		err = errs.ErrWrongCredentials

		return
	}

	user.ID = userFromDB.ID
	user.Email = userFromDB.Email

	return
}

func (r *GophKeeperRepo) AddUser(ctx context.Context, email, hashedPassword string) (user entity.User, err error) {
	newUser := models.User{
		Email:    email,
		Password: hashedPassword,
	}
	result := r.db.WithContext(ctx).Create(&newUser)

	if result.Error == nil {
		user.ID = newUser.ID
		user.Email = newUser.Email

		return
	}

	switch errs.ParsePostgresErr(result.Error).Code {
	case "23505":
		r.l.Debug("AddUser - %w", result.Error)

		err = errs.ErrEmailAlreadyExists

		return
	default:
		err = fmt.Errorf("AddUser - %w", result.Error)

		return
	}
}

func (r *GophKeeperRepo) GetUserByID(ctx context.Context, id string) (user entity.User, err error) {
	var userFromDB models.User

	r.db.WithContext(ctx).Where("id = ?", id).First(&userFromDB)

	if userFromDB.ID == uuid.Nil {
		err = errs.ErrWrongCredentials

		return
	}

	user.ID = userFromDB.ID
	user.Email = userFromDB.Email

	return
}
