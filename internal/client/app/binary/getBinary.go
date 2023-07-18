package app

import (
	"log"

	"github.com/LorezV/gophkeeper/internal/client/usecase"
	"github.com/spf13/cobra"
)

var GetBinary = &cobra.Command{ //nolint:gochecknoglobals // cobra style guide
	Use:   "getbinary",
	Short: "Get user file by id",
	Long: `
This command get user binary info and encode it for path
Usage: getbinary -i \"binary_id\" -f some_file -p pass 
Flags:
  -i, --id string Note id
  -p, --password string   User password value
  -f --file string File path.`,
	Run: func(cmd *cobra.Command, args []string) {
		usecase.GetClientUseCase().GetBinary(userPassword, getBinaryID, filePath)
	},
}

var (
	getBinaryID string //nolint:gochecknoglobals // cobra style guide
	filePath    string //nolint:gochecknoglobals // cobra style guide
)

func init() {
	GetBinary.Flags().StringVarP(&userPassword, "password", "p", "", "User password value.")
	GetBinary.Flags().StringVarP(&getBinaryID, "id", "i", "", "Binary id")
	GetBinary.Flags().StringVarP(&filePath, "file", "f", "", "User file")

	if err := GetBinary.MarkFlagRequired("password"); err != nil {
		log.Fatal(err)
	}
	if err := GetBinary.MarkFlagRequired("id"); err != nil {
		log.Fatal(err)
	}
	if err := AddBinary.MarkFlagRequired("file"); err != nil {
		log.Fatal(err)
	}
}
