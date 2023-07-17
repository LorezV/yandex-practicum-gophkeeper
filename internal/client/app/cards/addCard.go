package app

import (
	"log"

	"github.com/LorezV/gophkeeper/internal/client/usecase"
	"github.com/LorezV/gophkeeper/internal/entity"
	utils "github.com/LorezV/gophkeeper/internal/utils/client"
	"github.com/spf13/cobra"
)

var AddCard = &cobra.Command{ //nolint:gochecknoglobals // cobra style guide
	Use:   "addcard",
	Short: "Add card",
	Long: `
This command add card
Usage: addcard -p \"user_password\" 
Flags:
  -b, --brand string      Card brand
  -c, --code string       Card code
  -h, --help              help for addcard
  -m, --month string      Card expiration month
  -n, --number string     Card namber
  -o, --owner string      Card holder name
  -p, --password string   User password value.
  -t, --title string      Card title
  -y, --year string       Card expiration year
  --meta 				  Add meta data for entiry
  example: --meta'[{"name":"some_meta","value":"some_meta_value"},{"name":"some_meta2","value":"some_meta_value2"}]'
  `,
	Run: func(cmd *cobra.Command, args []string) {
		usecase.GetClientUseCase().AddCard(userPassword, &cardForAdditing)
	},
}

var (
	cardForAdditing entity.Card //nolint:gochecknoglobals // cobra style guide
	userPassword    string      //nolint:gochecknoglobals // cobra style guide
)

func init() {
	AddCard.Flags().StringVarP(&userPassword, "password", "p", "", "User password value.")
	AddCard.Flags().StringVarP(&cardForAdditing.Name, "title", "t", "", "Card title")
	AddCard.Flags().StringVarP(&cardForAdditing.Number, "number", "n", "", "Card namber")
	AddCard.Flags().StringVarP(&cardForAdditing.CardHolderName, "owner", "o", "", "Card holder name")
	AddCard.Flags().StringVarP(&cardForAdditing.Brand, "brand", "b", "", "Card brand")
	AddCard.Flags().StringVarP(&cardForAdditing.SecurityCode, "code", "c", "", "Card code")
	AddCard.Flags().StringVarP(&cardForAdditing.ExpirationMonth, "month", "m", "", "Card expiration month")
	AddCard.Flags().StringVarP(&cardForAdditing.ExpirationYear, "year", "y", "", "Card expiration year")
	AddCard.Flags().Var(&utils.JSONFlag{Target: &cardForAdditing.Meta}, "meta", `Add meta fields for entity`)

	if err := AddCard.MarkFlagRequired("password"); err != nil {
		log.Fatal(err)
	}
	if err := AddCard.MarkFlagRequired("title"); err != nil {
		log.Fatal(err)
	}
	if err := AddCard.MarkFlagRequired("number"); err != nil {
		log.Fatal(err)
	}
	if err := AddCard.MarkFlagRequired("password"); err != nil {
		log.Fatal(err)
	}
}
