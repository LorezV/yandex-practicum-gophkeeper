package app

import (
	"github.com/LorezV/gophkeeper/internal/client/usecase"
	"github.com/LorezV/gophkeeper/internal/entity"
	"github.com/spf13/cobra"
)

var RequiredUserArgs = 2 //nolint:gochecknoglobals // cobra style guide

var LoginUserCmd = &cobra.Command{ //nolint:gochecknoglobals // cobra style guide
	Use:   "login",
	Short: "Login user to the service",
	Long: `
This command login user.
Usage: gophkeeperclient login user_login user_password`,
	Args: cobra.MinimumNArgs(RequiredUserArgs),
	Run: func(cmd *cobra.Command, args []string) {
		account := entity.User{
			Email:    args[0],
			Password: args[1],
		}

		usecase.GetClientUseCase().Login(&account)
	},
}
