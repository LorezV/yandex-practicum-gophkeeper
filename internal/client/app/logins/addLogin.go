package app

import (
	"log"

	"github.com/LorezV/gophkeeper/internal/client/usecase"
	"github.com/LorezV/gophkeeper/internal/entity"
	utils "github.com/LorezV/gophkeeper/internal/utils/client"
	"github.com/spf13/cobra"
)

var AddLogin = &cobra.Command{ //nolint:gochecknoglobals // cobra style guide
	Use:   "addlogin",
	Short: "Add login",
	Long: `
This command add logit for site
Usage: addlogin -p \"user_password\" 
Flags:
  -h, --help              help for addlogin
  -l, --login string      Site login
  -p, --password string   User password value.
  -s, --secret string     Site password|secret
  -t, --title string      Login title
  -u, --uri string        Site endloint  
  --meta 				  Add meta data for entiry
  example: --meta'[{"name":"some_meta","value":"some_meta_value"},{"name":"some_meta2","value":"some_meta_value2"}]'
  `,
	Run: func(cmd *cobra.Command, args []string) {
		usecase.GetClientUseCase().AddLogin(userPassword, &loginForAdditing)
	},
}

var (
	loginForAdditing entity.Login //nolint:gochecknoglobals // cobra style guide
	userPassword     string       //nolint:gochecknoglobals // cobra style guide
)

func init() {
	AddLogin.Flags().StringVarP(&userPassword, "password", "p", "", "User password value.")

	AddLogin.Flags().StringVarP(&loginForAdditing.Name, "title", "t", "", "Login title")
	AddLogin.Flags().StringVarP(&loginForAdditing.Login, "login", "l", "", "Site login")
	AddLogin.Flags().StringVarP(&loginForAdditing.Password, "secret", "s", "", "Site password|secret")
	AddLogin.Flags().StringVarP(&loginForAdditing.URI, "uri", "u", "", "Site endloint")
	AddLogin.Flags().Var(&utils.JSONFlag{Target: &loginForAdditing.Meta}, "meta", `Add meta fields for entity`)

	if err := AddLogin.MarkFlagRequired("password"); err != nil {
		log.Fatal(err)
	}
	if err := AddLogin.MarkFlagRequired("title"); err != nil {
		log.Fatal(err)
	}
}
