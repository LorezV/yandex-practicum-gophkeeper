package app

import (
	"log"

	"github.com/LorezV/gophkeeper/internal/client/usecase"
	"github.com/spf13/cobra"
)

var DelBinary = &cobra.Command{ //nolint:gochecknoglobals // cobra style guide
	Use:   "delnote",
	Short: "Delete user file by id",
	Long: `
This command remove file
Usage: delnote -i \"binary_id\" 
Flags:
  -i, --id string Card id
  -p, --password string   User password value.`,
	Run: func(cmd *cobra.Command, args []string) {
		usecase.GetClientUseCase().DelBinary(userPassword, delBinaryID)
	},
}

var delBinaryID string //nolint:gochecknoglobals // cobra style guide

func init() {
	DelBinary.Flags().StringVarP(&userPassword, "password", "p", "", "User password value.")
	DelBinary.Flags().StringVarP(&delBinaryID, "id", "i", "", "Binary id")

	if err := DelBinary.MarkFlagRequired("password"); err != nil {
		log.Fatal(err)
	}
	if err := DelBinary.MarkFlagRequired("id"); err != nil {
		log.Fatal(err)
	}
}
