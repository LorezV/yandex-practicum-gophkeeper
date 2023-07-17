package usecase

import (
	"github.com/LorezV/gophkeeper/internal/client/usecase/viewsets"
	"github.com/LorezV/gophkeeper/internal/entity"
	"github.com/google/uuid"
)

type (
	// GophKeeperClient - use cases.
	GophKeeperClient interface {
		InitDB()

		Register(user *entity.User)
		Login(user *entity.User)
		Logout()
		Sync(userPassword string)

		ShowVault(userPassword, showVaultOption string)

		AddCard(userPassword string, card *entity.Card)
		ShowCard(userPassword, cardID string)
		DelCard(userPassword, cardID string)

		AddLogin(userPassword string, login *entity.Login)
		ShowLogin(userPassword, loginID string)
		DelLogin(userPassword, loginID string)

		AddNote(userPassword string, note *entity.SecretNote)
		ShowNote(userPassword, noteID string)
		DelNote(userPassword, noteID string)

		AddBinary(userPassword string, binary *entity.Binary)
		DelBinary(userPassword, binaryID string)
		GetBinary(userPassword, getBinaryID, filePath string)
	}
	GophKeeperClientRepo interface {
		MigrateDB()
		AddUser(user *entity.User) error
		UpdateUserToken(user *entity.User, token *entity.JWT) error
		DropUserToken() error
		GetSavedAccessToken() (string, error)
		RemoveUsers()
		UserExistsByEmail(email string) bool
		GetUserPasswordHash() string

		AddCard(*entity.Card) error
		SaveCards([]entity.Card) error
		LoadCards() []viewsets.CardForList
		GetCardByID(cardID uuid.UUID) (entity.Card, error)
		DelCard(cardID uuid.UUID) error

		AddLogin(*entity.Login) error
		SaveLogins([]entity.Login) error
		LoadLogins() []viewsets.LoginForList
		GetLoginByID(loginID uuid.UUID) (entity.Login, error)
		DelLogin(loginID uuid.UUID) error

		LoadNotes() []viewsets.NoteForList
		SaveNotes([]entity.SecretNote) error
		AddNote(*entity.SecretNote) error
		GetNoteByID(notedID uuid.UUID) (entity.SecretNote, error)
		DelNote(noteID uuid.UUID) error

		LoadBinaries() []viewsets.BinaryForList
		SaveBinaries([]entity.Binary) error
		AddBinary(*entity.Binary) error
		GetBinaryByID(binarydID uuid.UUID) (entity.Binary, error)
		DelBinary(binaryID uuid.UUID) error
	}
	GophKeeperClientAPI interface {
		Login(user *entity.User) (entity.JWT, error)
		Register(user *entity.User) error

		GetCards(accessToken string) ([]entity.Card, error)
		AddCard(accessToken string, card *entity.Card) error
		DelCard(accessToken, cardID string) error

		GetLogins(accessToken string) ([]entity.Login, error)
		AddLogin(accessToken string, login *entity.Login) error
		DelLogin(accessToken, loginID string) error

		GetNotes(accessToken string) ([]entity.SecretNote, error)
		AddNote(accessToken string, note *entity.SecretNote) error
		DelNote(accessToken, noteID string) error

		GetBinaries(accessToken string) ([]entity.Binary, error)
		AddBinary(accessToken string, binary *entity.Binary, tmpFilePath string) error
		DelBinary(accessToken, binaryID string) error
		DownloadBinary(accessToken, outpuFilePath string, binary *entity.Binary) error
	}
)
