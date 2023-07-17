package app

import (
	"log"

	"github.com/LorezV/gophkeeper/internal/client/usecase"
	"github.com/spf13/cobra"
)

var GetNote = &cobra.Command{ //nolint:gochecknoglobals // cobra style guide
	Use:   "getnote",
	Short: "Show user note by id",
	Long: `
This command show user note
Usage: getnote -i \"note_id\" 
Flags:
  -i, --id string Note id
  -p, --password string   User password value.`,
	Run: func(cmd *cobra.Command, args []string) {
		usecase.GetClientUseCase().ShowNote(userPassword, getNoteID)
	},
}

var getNoteID string //nolint:gochecknoglobals // cobra style guide

func init() {
	GetNote.Flags().StringVarP(&userPassword, "password", "p", "", "User password value.")
	GetNote.Flags().StringVarP(&getNoteID, "id", "i", "", "Note id")

	if err := GetNote.MarkFlagRequired("password"); err != nil {
		log.Fatal(err)
	}
	if err := GetNote.MarkFlagRequired("id"); err != nil {
		log.Fatal(err)
	}
}
