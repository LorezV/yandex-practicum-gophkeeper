package app

import (
	"log"

	"github.com/LorezV/gophkeeper/internal/client/usecase"
	"github.com/spf13/cobra"
)

var DelLogin = &cobra.Command{ //nolint:gochecknoglobals // cobra style guide
	Use:   "dellogin",
	Short: "Delete user login by id",
	Long: `
This command remove login
Usage: delcard -i \"login_id\" 
Flags:
  -i, --id string Card id
  -p, --password string   User password value.`,
	Run: func(cmd *cobra.Command, args []string) {
		usecase.GetClientUseCase().DelLogin(userPassword, delLoginID)
	},
}

var delLoginID string //nolint:gochecknoglobals // cobra style guide

func init() {
	DelLogin.Flags().StringVarP(&userPassword, "password", "p", "", "User password value.")
	DelLogin.Flags().StringVarP(&delLoginID, "id", "i", "", "Card id")

	if err := DelLogin.MarkFlagRequired("password"); err != nil {
		log.Fatal(err)
	}
	if err := DelLogin.MarkFlagRequired("id"); err != nil {
		log.Fatal(err)
	}
}
