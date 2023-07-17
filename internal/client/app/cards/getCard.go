package app

import (
	"log"

	"github.com/LorezV/gophkeeper/internal/client/usecase"
	"github.com/spf13/cobra"
)

var GetCard = &cobra.Command{ //nolint:gochecknoglobals // cobra style guide
	Use:   "getcard",
	Short: "Show user card by id",
	Long: `
This command add card
Usage: getcard -i \"card_id\" 
Flags:
  -i, --id string Card id
  -p, --password string   User password value.`,
	Run: func(cmd *cobra.Command, args []string) {
		usecase.GetClientUseCase().ShowCard(userPassword, getCardID)
	},
}

var getCardID string //nolint:gochecknoglobals // cobra style guide

func init() {
	GetCard.Flags().StringVarP(&userPassword, "password", "p", "", "User password value.")
	GetCard.Flags().StringVarP(&getCardID, "id", "i", "", "Card id")

	if err := GetCard.MarkFlagRequired("password"); err != nil {
		log.Fatal(err)
	}
	if err := GetCard.MarkFlagRequired("id"); err != nil {
		log.Fatal(err)
	}
}
