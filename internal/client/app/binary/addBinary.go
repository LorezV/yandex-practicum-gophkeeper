package app

import (
	"log"

	"github.com/LorezV/gophkeeper/internal/client/usecase"
	"github.com/LorezV/gophkeeper/internal/entity"
	utils "github.com/LorezV/gophkeeper/internal/utils/client"
	"github.com/spf13/cobra"
)

var AddBinary = &cobra.Command{ //nolint:gochecknoglobals // cobra style guide
	Use:   "addbinary",
	Short: "Add binary",
	Long: `
This command add user file
Usage: addnote -p \"user_password\" -f "file_location" 
Flags:
  -h, --help              help for addlogin
  -p, --password string   User password value.
  -f, --file string       User file  
  --meta 				  Add meta data for entiry
  example: --meta'[{"name":"some_meta","value":"some_meta_value"},{"name":"some_meta2","value":"some_meta_value2"}]'
  `,
	Run: func(cmd *cobra.Command, args []string) {
		usecase.GetClientUseCase().AddBinary(userPassword, &binaryForAdditing)
	},
}

var (
	binaryForAdditing entity.Binary //nolint:gochecknoglobals // cobra style guide
	userPassword      string        //nolint:gochecknoglobals // cobra style guide
)

func init() {
	AddBinary.Flags().StringVarP(&userPassword, "password", "p", "", "User password value.")

	AddBinary.Flags().StringVarP(&binaryForAdditing.Name, "title", "t", "", "Login title")
	AddBinary.Flags().StringVarP(&binaryForAdditing.FileName, "file", "f", "", "User file")
	AddBinary.Flags().Var(&utils.JSONFlag{Target: &binaryForAdditing.Meta}, "meta", `Add meta fields for entity`)

	if err := AddBinary.MarkFlagRequired("password"); err != nil {
		log.Fatal(err)
	}
	if err := AddBinary.MarkFlagRequired("title"); err != nil {
		log.Fatal(err)
	}
	if err := AddBinary.MarkFlagRequired("file"); err != nil {
		log.Fatal(err)
	}
}
