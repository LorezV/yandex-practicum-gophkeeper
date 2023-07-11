package app

import (
	"log"

	"github.com/LorezV/gophkeeper/internal/client/usecase"
	"github.com/spf13/cobra"
)

var SyncUserData = &cobra.Command{ //nolint:gochecknoglobals // cobra style guide
	Use:   "sync",
	Short: "Sync user`s data",
	Long: `
This command update users private data from server
Usage: gophkeeperclient sync -p \"user_password\"`,
	Run: func(cmd *cobra.Command, args []string) {
		usecase.GetClientUseCase().Sync(userPassword)
	},
}

var userPassword string //nolint:gochecknoglobals // cobra style guide

func init() {
	SyncUserData.Flags().StringVarP(&userPassword, "password", "p", "", "User password value.")
	if err := SyncUserData.MarkFlagRequired("password"); err != nil {
		log.Fatal(err)
	}
}
