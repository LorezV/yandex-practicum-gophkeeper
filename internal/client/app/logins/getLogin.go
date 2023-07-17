package app

import (
	"log"

	"github.com/LorezV/gophkeeper/internal/client/usecase"
	"github.com/spf13/cobra"
)

var GetLogin = &cobra.Command{ //nolint:gochecknoglobals // cobra style guide
	Use:   "getlogin",
	Short: "Show user login by id",
	Long: `
This command getlogin
Usage: getlogin -i \"login_id\" 
Flags:
  -i, --id string Login id
  -p, --password string   User password value.`,
	Run: func(cmd *cobra.Command, args []string) {
		usecase.GetClientUseCase().ShowLogin(userPassword, getLoginID)
	},
}

var getLoginID string //nolint:gochecknoglobals // cobra style guide

func init() {
	GetLogin.Flags().StringVarP(&userPassword, "password", "p", "", "User password value.")
	GetLogin.Flags().StringVarP(&getLoginID, "id", "i", "", "Card id")

	if err := GetLogin.MarkFlagRequired("password"); err != nil {
		log.Fatal(err)
	}
	if err := GetLogin.MarkFlagRequired("id"); err != nil {
		log.Fatal(err)
	}
}
