package app

import (
	"fmt"
	"log"
	"os"

	config "github.com/LorezV/gophkeeper/config/client"
	auth "github.com/LorezV/gophkeeper/internal/client/app/auth"
	binary "github.com/LorezV/gophkeeper/internal/client/app/binary"
	"github.com/LorezV/gophkeeper/internal/client/app/build"
	cards "github.com/LorezV/gophkeeper/internal/client/app/cards"
	logins "github.com/LorezV/gophkeeper/internal/client/app/logins"
	notes "github.com/LorezV/gophkeeper/internal/client/app/notes"
	vault "github.com/LorezV/gophkeeper/internal/client/app/vault"
	"github.com/LorezV/gophkeeper/internal/client/usecase"
	api "github.com/LorezV/gophkeeper/internal/client/usecase/client_api"
	"github.com/LorezV/gophkeeper/internal/client/usecase/repo"
	"github.com/spf13/cobra"
)

var (
	cfg           *config.Config           //nolint:gochecknoglobals // cobra style guide
	ClientUseCase usecase.GophKeeperClient //nolint:gochecknoglobals // cobra style guide

	rootCmd = &cobra.Command{ //nolint:gochecknoglobals // cobra style guide
		Use:   config.LoadConfig().App.Name,
		Short: "App for storing private data",
		Long:  `User can save cards, note and logins`,
		Run: func(cmd *cobra.Command, args []string) {
			build.PrintBulidInfo()
		},
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initApp)
	commands := []*cobra.Command{
		auth.LoginUserCmd,
		auth.LogoutUser,
		auth.RegisterUserCmd,

		vault.RegisterInitLocalStorage,
		vault.ShowVault,
		vault.SyncUserData,

		cards.AddCard,
		cards.DelCard,
		cards.GetCard,

		logins.AddLogin,
		logins.DelLogin,
		logins.GetLogin,

		notes.AddNote,
		notes.DelNote,
		notes.GetNote,

		binary.AddBinary,
		binary.DelBinary,
		binary.GetBinary,
	}

	rootCmd.AddCommand(commands...)
}

func initApp() {
	cfg = config.LoadConfig()

	log.Println(cfg)

	uc := usecase.GetClientUseCase()
	clientOpts := []usecase.GophKeeperUseCaseOpts{
		usecase.SetAPI(api.New(cfg.Server.URL)),
		usecase.SetConfig(cfg),
		usecase.SetRepo(repo.New(cfg.SQLite.DSN)),
	}

	for _, opt := range clientOpts {
		opt(uc)
	}

	if _, err := os.Stat(cfg.FilesStorage.Location); os.IsNotExist(err) {
		err = os.MkdirAll(cfg.FilesStorage.Location, os.ModePerm)
		if err != nil {
			log.Fatalf("App.Init - os.MkdirAll - %v", err)
		}
	}
}
