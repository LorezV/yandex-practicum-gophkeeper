package repo

import (
	"fmt"

	"github.com/LorezV/gophkeeper/internal/client/usecase/repo/models"
	"github.com/LorezV/gophkeeper/internal/entity"
	"github.com/LorezV/gophkeeper/internal/utils"
)

func (r *GophKeeperRepo) RemoveUsers() {
	r.db.Exec("DELETE FROM users")
}

func (r *GophKeeperRepo) AddUser(user *entity.User) error {
	r.RemoveUsers()
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return fmt.Errorf("GophKeeperRepo - AddUser - HashPassword - %w", err)
	}

	newUser := models.User{
		Email:    user.Email,
		Password: hashedPassword,
	}

	return r.db.Create(&newUser).Error
}

func (r *GophKeeperRepo) UpdateUserToken(user *entity.User, token *entity.JWT) error {
	var existedUser models.User

	r.db.Where("email", user.Email).First(&existedUser)
	existedUser.AccessToken = token.AccessToken
	existedUser.RefreshToken = token.RefreshToken

	return r.db.Save(&existedUser).Error
}

func (r *GophKeeperRepo) UserExistsByEmail(email string) bool {
	var user models.User

	r.db.Where("email = ?", email).First(&user)

	return user.ID != 0
}

func (r *GophKeeperRepo) DropUserToken() error {
	var existedUser models.User

	r.db.First(&existedUser)
	existedUser.AccessToken = ""
	existedUser.RefreshToken = ""

	return r.db.Save(&existedUser).Error
}

func (r *GophKeeperRepo) GetUserPasswordHash() string {
	var existedUser models.User

	r.db.First(&existedUser)

	return existedUser.Password
}

func (r *GophKeeperRepo) GetSavedAccessToken() (accessToken string, err error) {
	var user models.User
	err = r.db.First(&user).Error

	return user.AccessToken, err
}

func (r *GophKeeperRepo) getUserID() uint {
	var user models.User
	r.db.First(&user)

	return user.ID
}
