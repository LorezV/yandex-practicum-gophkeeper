package usecase

import (
	"errors"
	"sync"

	config "github.com/LorezV/gophkeeper/config/client"
	"github.com/fatih/color"
)

type GophKeeperClientUseCase struct {
	repo      GophKeeperClientRepo
	clientAPI GophKeeperClientAPI
	cfg       *config.Config
}

var (
	clientUseCase *GophKeeperClientUseCase //nolint:gochecknoglobals // pattern singleton
	once          sync.Once                //nolint:gochecknoglobals // pattern singleton
)

func GetClientUseCase() *GophKeeperClientUseCase {
	once.Do(func() {
		clientUseCase = &GophKeeperClientUseCase{}
	})

	return clientUseCase
}

type GophKeeperUseCaseOpts func(*GophKeeperClientUseCase)

func SetRepo(r GophKeeperClientRepo) GophKeeperUseCaseOpts {
	return func(gkuc *GophKeeperClientUseCase) {
		gkuc.repo = r
	}
}

func SetAPI(clientAPI GophKeeperClientAPI) GophKeeperUseCaseOpts {
	return func(gkuc *GophKeeperClientUseCase) {
		gkuc.clientAPI = clientAPI
	}
}

func SetConfig(cfg *config.Config) GophKeeperUseCaseOpts {
	return func(gkuc *GophKeeperClientUseCase) {
		gkuc.cfg = cfg
	}
}

func (uc *GophKeeperClientUseCase) InitDB() {
	uc.repo.MigrateDB()
}

var (
	errPasswordCheck = errors.New("wrong password")
	errToken         = errors.New("user token erroe")
)

func (uc *GophKeeperClientUseCase) authorisationCheck(userPassword string) (string, error) {
	if !uc.verifyPassword(userPassword) {
		return "", errPasswordCheck
	}
	accessToken, err := uc.repo.GetSavedAccessToken()
	if err != nil || accessToken == "" {
		color.Red("User should be logged")

		return "", errToken
	}

	return accessToken, nil
}
