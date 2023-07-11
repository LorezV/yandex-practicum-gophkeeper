package usecase

import (
	"log"

	"github.com/fatih/color"

	"github.com/LorezV/gophkeeper/internal/entity"
	"github.com/LorezV/gophkeeper/internal/utils"
)

func (uc *GophKeeperClientUseCase) Login(user *entity.User) {
	token, err := uc.clientAPI.Login(user)
	if err != nil {
		return
	}

	if !uc.repo.UserExistsByEmail(user.Email) {
		err = uc.repo.AddUser(user)
		if err != nil {
			log.Fatal(err)
		}
	}
	if err = uc.repo.UpdateUserToken(user, &token); err != nil {
		log.Fatal(err)
	}
	color.Green("Got authorization token for %q", user.Email)
}

func (uc *GophKeeperClientUseCase) Register(user *entity.User) {
	if err := uc.clientAPI.Register(user); err != nil {
		return
	}

	if err := uc.repo.AddUser(user); err != nil {
		color.Red("Internal error: %v", err)

		return
	}

	color.Green("User registered")
	color.Green("ID: %v", user.ID)
	color.Green("Email: %s", user.Email)
}

func (uc *GophKeeperClientUseCase) Logout() {
	if err := uc.repo.DropUserToken(); err != nil {
		color.Red("Internal error: %v", err)

		return
	}

	color.Green("Users tokens were dropped")
}

func (uc *GophKeeperClientUseCase) Sync(userPassword string) {
	if !uc.verifyPassword(userPassword) {
		return
	}
	accessToken, err := uc.repo.GetSavedAccessToken()
	if err != nil {
		color.Red("Internal error: %v", err)

		return
	}
	uc.loadCards(accessToken)
	uc.loadLogins(accessToken)
	uc.loadNotes(accessToken)
	uc.loadBinaries(accessToken)
}

func (uc *GophKeeperClientUseCase) verifyPassword(userPassword string) bool {
	if err := utils.VerifyPassword(uc.repo.GetUserPasswordHash(), userPassword); err != nil {
		color.Red("Password check status: failed")

		return false
	}

	return true
}
