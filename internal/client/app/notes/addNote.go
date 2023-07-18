package app

import (
	"log"

	"github.com/LorezV/gophkeeper/internal/client/usecase"
	"github.com/LorezV/gophkeeper/internal/entity"
	utils "github.com/LorezV/gophkeeper/internal/utils/client"
	"github.com/spf13/cobra"
)

var AddNote = &cobra.Command{ //nolint:gochecknoglobals // cobra style guide
	Use:   "addnote",
	Short: "Add note",
	Long: `
This command add user note
Usage: addnote -p \"user_password\" 
Flags:
  -h, --help              help for addlogin
  -p, --password string   User password value.
  -n, --note string     User note  
  --meta 				  Add meta data for entiry
  example: --meta'[{"name":"some_meta","value":"some_meta_value"},{"name":"some_meta2","value":"some_meta_value2"}]'
  `,
	Run: func(cmd *cobra.Command, args []string) {
		usecase.GetClientUseCase().AddNote(userPassword, &noteForAdditing)
	},
}

var (
	noteForAdditing entity.SecretNote //nolint:gochecknoglobals // cobra style guide
	userPassword    string            //nolint:gochecknoglobals // cobra style guide
)

func init() {
	AddNote.Flags().StringVarP(&userPassword, "password", "p", "", "User password value.")

	AddNote.Flags().StringVarP(&noteForAdditing.Name, "title", "t", "", "Login title")
	AddNote.Flags().StringVarP(&noteForAdditing.Note, "note", "n", "", "User note")
	AddNote.Flags().Var(&utils.JSONFlag{Target: &noteForAdditing.Meta}, "meta", `Add meta fields for entity`)

	if err := AddNote.MarkFlagRequired("password"); err != nil {
		log.Fatal(err)
	}
	if err := AddNote.MarkFlagRequired("title"); err != nil {
		log.Fatal(err)
	}
}
