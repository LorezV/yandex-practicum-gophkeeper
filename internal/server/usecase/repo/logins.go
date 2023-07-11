package repo

import (
	"context"

	"github.com/LorezV/gophkeeper/internal/entity"
	"github.com/LorezV/gophkeeper/internal/server/usecase/repo/models"
	"github.com/LorezV/gophkeeper/internal/utils/errs"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (r *GophKeeperRepo) GetLogins(ctx context.Context, user entity.User) (logins []entity.Login, err error) {
	var loginsFromDB []models.Login

	err = r.db.WithContext(ctx).
		Model(&models.Login{}).
		Preload("Meta").
		Find(&loginsFromDB, "user_id = ?", user.ID).Error
	if err != nil {
		return nil, err
	}

	if len(loginsFromDB) == 0 {
		return nil, nil
	}

	logins = make([]entity.Login, len(loginsFromDB))

	for index := range loginsFromDB {
		logins[index].ID = loginsFromDB[index].ID
		logins[index].Name = loginsFromDB[index].Name
		logins[index].Password = loginsFromDB[index].Password
		logins[index].URI = loginsFromDB[index].URI
		logins[index].Login = loginsFromDB[index].Login
	}

	return logins, nil
}

func (r *GophKeeperRepo) AddLogin(ctx context.Context, login *entity.Login, userID uuid.UUID) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		loginToDB := models.Login{
			ID:       uuid.New(),
			UserID:   userID,
			Name:     login.Name,
			Password: login.Password,
			URI:      login.URI,
			Login:    login.Login,
		}

		if err := tx.WithContext(ctx).Create(&loginToDB).Error; err != nil {
			return err
		}

		login.ID = loginToDB.ID
		for index, meta := range login.Meta {
			metaForLogin := models.MetaLogin{
				Name:    meta.Name,
				Value:   meta.Value,
				LoginID: loginToDB.ID,
			}
			if err := tx.WithContext(ctx).Create(&metaForLogin).Error; err != nil {
				return err
			}
			login.Meta[index].ID = metaForLogin.ID
		}

		return nil
	})
}

func (r *GophKeeperRepo) IsLoginOwner(ctx context.Context, loginID, userID uuid.UUID) bool {
	var loginFromDB models.Login

	r.db.WithContext(ctx).Where("id = ?", loginID).First(&loginFromDB)

	return loginFromDB.UserID == userID
}

func (r *GophKeeperRepo) DelLogin(ctx context.Context, loginID, userID uuid.UUID) error {
	if !r.IsLoginOwner(ctx, loginID, userID) {
		return errs.ErrWrongOwnerOrNotFound
	}

	return r.db.WithContext(ctx).Delete(&models.Login{}, loginID).Error
}

func (r *GophKeeperRepo) UpdateLogin(ctx context.Context, login *entity.Login, userID uuid.UUID) error {
	if !r.IsLoginOwner(ctx, login.ID, userID) {
		return errs.ErrWrongOwnerOrNotFound
	}

	return r.db.Transaction(func(tx *gorm.DB) error {
		loginToDB := models.Login{
			ID:       login.ID,
			Name:     login.Name,
			Password: login.Password,
			URI:      login.URI,
			Login:    login.Login,
			UserID:   userID,
		}

		if err := tx.WithContext(ctx).Save(&loginToDB).Error; err != nil {
			return err
		}
		login.ID = loginToDB.ID
		for _, meta := range login.Meta {
			metaForLogin := models.MetaLogin{
				Name:    meta.Name,
				Value:   meta.Value,
				LoginID: loginToDB.ID,
				ID:      meta.ID,
			}
			if err := tx.WithContext(ctx).Create(&metaForLogin).Error; err != nil {
				return err
			}
		}

		return nil
	})
}
