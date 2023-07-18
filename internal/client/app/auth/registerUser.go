package app

import (
	"github.com/LorezV/gophkeeper/internal/client/usecase"
	"github.com/LorezV/gophkeeper/internal/entity"
	"github.com/spf13/cobra"
)

var RegisterUserCmd = &cobra.Command{ //nolint:gochecknoglobals // cobra style guide
	Use:   "register",
	Short: "Register user to the service",
	Long: `
This command register new user.
Usage: gophkeeperclient register user_login user_password`,
	Args: cobra.MinimumNArgs(RequiredUserArgs),
	Run: func(cmd *cobra.Command, args []string) {
		account := entity.User{
			Email:    args[0],
			Password: args[1],
		}

		usecase.GetClientUseCase().Register(&account)
	},
}
