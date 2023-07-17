// Package usecase implements application business logic. Each logic group in own file.
package usecase

import (
	"context"
	"mime/multipart"

	"github.com/LorezV/gophkeeper/internal/entity"
	"github.com/google/uuid"
)

type (
	// GophKeeper - use cases.
	GophKeeper interface {
		HealthCheck() error
		SignUpUser(ctx context.Context, email, password string) (entity.User, error)
		SignInUser(ctx context.Context, email, password string) (entity.JWT, error)
		RefreshAccessToken(ctx context.Context, refreshToken string) (entity.JWT, error)
		GetDomainName() string
		CheckAccessToken(ctx context.Context, accessToken string) (entity.User, error)

		GetLogins(ctx context.Context, user entity.User) ([]entity.Login, error)
		AddLogin(ctx context.Context, login *entity.Login, userID uuid.UUID) error
		DelLogin(ctx context.Context, loginID, userID uuid.UUID) error
		UpdateLogin(ctx context.Context, login *entity.Login, userID uuid.UUID) error

		GetCards(ctx context.Context, user entity.User) ([]entity.Card, error)
		AddCard(ctx context.Context, card *entity.Card, userID uuid.UUID) error
		DelCard(ctx context.Context, cardUUID, userID uuid.UUID) error
		UpdateCard(ctx context.Context, card *entity.Card, userID uuid.UUID) error

		GetNotes(ctx context.Context, user entity.User) ([]entity.SecretNote, error)
		AddNote(ctx context.Context, note *entity.SecretNote, userID uuid.UUID) error
		DelNote(ctx context.Context, noteID, userID uuid.UUID) error
		UpdateNote(ctx context.Context, note *entity.SecretNote, userID uuid.UUID) error

		GetBinaries(ctx context.Context, user entity.User) ([]entity.Binary, error)
		AddBinary(ctx context.Context, binary *entity.Binary, file *multipart.FileHeader, userID uuid.UUID) error
		GetUserBinary(ctx context.Context, currentUser *entity.User, binaryUUID uuid.UUID) (string, error)
		DelUserBinary(ctx context.Context, currentUser *entity.User, binaryUUID uuid.UUID) error
		AddBinaryMeta(
			ctx context.Context,
			currentUser *entity.User,
			binaryUUID uuid.UUID,
			meta []entity.Meta) (*entity.Binary, error)
	}

	// GophKeeperRepo - db logic.
	GophKeeperRepo interface {
		DBHealthCheck() error
		AddUser(ctx context.Context, email, hashedPassword string) (entity.User, error)
		GetUserByEmail(ctx context.Context, email, hashedPassword string) (entity.User, error)
		GetUserByID(ctx context.Context, id string) (entity.User, error)

		GetLogins(ctx context.Context, user entity.User) ([]entity.Login, error)
		AddLogin(ctx context.Context, login *entity.Login, userID uuid.UUID) error
		DelLogin(ctx context.Context, loginID, userID uuid.UUID) error
		UpdateLogin(ctx context.Context, login *entity.Login, userID uuid.UUID) error
		IsLoginOwner(ctx context.Context, loginID, userID uuid.UUID) bool

		GetCards(ctx context.Context, user entity.User) ([]entity.Card, error)
		AddCard(ctx context.Context, card *entity.Card, userID uuid.UUID) error
		DelCard(ctx context.Context, cardUUID, userID uuid.UUID) error
		UpdateCard(ctx context.Context, card *entity.Card, userID uuid.UUID) error
		IsCardOwner(ctx context.Context, cardUUID, userID uuid.UUID) bool

		GetNotes(ctx context.Context, user entity.User) ([]entity.SecretNote, error)
		AddNote(ctx context.Context, note *entity.SecretNote, userID uuid.UUID) error
		DelNote(ctx context.Context, noteID, userID uuid.UUID) error
		UpdateNote(ctx context.Context, note *entity.SecretNote, userID uuid.UUID) error
		IsNoteOwner(ctx context.Context, noteID, userID uuid.UUID) bool

		GetBinaries(ctx context.Context, user entity.User) ([]entity.Binary, error)
		AddBinary(ctx context.Context, binary *entity.Binary, userID uuid.UUID) error
		GetBinary(ctx context.Context, binaryID, userID uuid.UUID) (*entity.Binary, error)
		DelUserBinary(ctx context.Context, currentUser *entity.User, binaryUUID uuid.UUID) error
		AddBinaryMeta(
			ctx context.Context,
			currentUser *entity.User,
			binaryUUID uuid.UUID,
			meta []entity.Meta) (*entity.Binary, error)
	}
)
