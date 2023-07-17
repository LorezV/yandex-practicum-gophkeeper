package app

import (
	"log"

	"github.com/LorezV/gophkeeper/internal/client/usecase"
	"github.com/spf13/cobra"
)

var DelNote = &cobra.Command{ //nolint:gochecknoglobals // cobra style guide
	Use:   "delnote",
	Short: "Delete user note by id",
	Long: `
This command remove note
Usage: delnote -i \"note_id\" 
Flags:
  -i, --id string Card id
  -p, --password string   User password value.`,
	Run: func(cmd *cobra.Command, args []string) {
		usecase.GetClientUseCase().DelNote(userPassword, delNoteID)
	},
}

var delNoteID string //nolint:gochecknoglobals // cobra style guide

func init() {
	DelNote.Flags().StringVarP(&userPassword, "password", "p", "", "User password value.")
	DelNote.Flags().StringVarP(&delNoteID, "id", "i", "", "Note id")

	if err := DelNote.MarkFlagRequired("password"); err != nil {
		log.Fatal(err)
	}
	if err := DelNote.MarkFlagRequired("id"); err != nil {
		log.Fatal(err)
	}
}
