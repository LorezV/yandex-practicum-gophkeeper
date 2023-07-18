package app

import (
	"github.com/LorezV/gophkeeper/internal/client/usecase"
	"github.com/spf13/cobra"
)

var LogoutUser = &cobra.Command{ //nolint:gochecknoglobals // cobra style guide
	Use:   "logout",
	Short: "Logout user",
	Long: `
This command drops users tokens`,
	Run: func(cmd *cobra.Command, args []string) {
		usecase.GetClientUseCase().Logout()
	},
}
