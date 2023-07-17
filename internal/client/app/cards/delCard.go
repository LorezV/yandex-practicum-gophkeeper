package app

import (
	"log"

	"github.com/LorezV/gophkeeper/internal/client/usecase"
	"github.com/spf13/cobra"
)

var DelCard = &cobra.Command{ //nolint:gochecknoglobals // cobra style guide
	Use:   "delcard",
	Short: "Delete user card by id",
	Long: `
This command remove card
Usage: delcard -i \"card_id\" 
Flags:
  -i, --id string Card id
  -p, --password string   User password value.`,
	Run: func(cmd *cobra.Command, args []string) {
		usecase.GetClientUseCase().DelCard(userPassword, delCardID)
	},
}

var delCardID string //nolint:gochecknoglobals // cobra style guide

func init() {
	DelCard.Flags().StringVarP(&userPassword, "password", "p", "", "User password value.")
	DelCard.Flags().StringVarP(&delCardID, "id", "i", "", "Card id")

	if err := DelCard.MarkFlagRequired("password"); err != nil {
		log.Fatal(err)
	}
	if err := DelCard.MarkFlagRequired("id"); err != nil {
		log.Fatal(err)
	}
}
