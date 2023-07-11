package usecase

import (
	"fmt"
	"log"

	"github.com/fatih/color"
	"github.com/google/uuid"

	"github.com/LorezV/gophkeeper/internal/entity"
	"github.com/LorezV/gophkeeper/internal/utils"
)

func (uc *GophKeeperClientUseCase) AddCard(userPassword string, card *entity.Card) {
	accessToken, err := uc.authorisationCheck(userPassword)
	if err != nil {
		return
	}

	uc.encryptCard(userPassword, card)

	if err = uc.clientAPI.AddCard(accessToken, card); err != nil {
		return
	}

	if err = uc.repo.AddCard(card); err != nil {
		log.Fatal(err)
	}

	color.Green("Card %q added, id: %v", card.Name, card.ID)
}

func (uc *GophKeeperClientUseCase) encryptCard(userPassword string, card *entity.Card) {
	card.Number = utils.Encrypt(userPassword, card.Number)
	card.SecurityCode = utils.Encrypt(userPassword, card.SecurityCode)
	card.ExpirationMonth = utils.Encrypt(userPassword, card.ExpirationMonth)
	card.ExpirationYear = utils.Encrypt(userPassword, card.ExpirationYear)
	card.CardHolderName = utils.Encrypt(userPassword, card.CardHolderName)
}

func (uc *GophKeeperClientUseCase) decryptCard(userPassword string, card *entity.Card) {
	card.Number = utils.Decrypt(userPassword, card.Number)
	card.SecurityCode = utils.Decrypt(userPassword, card.SecurityCode)
	card.ExpirationMonth = utils.Decrypt(userPassword, card.ExpirationMonth)
	card.ExpirationYear = utils.Decrypt(userPassword, card.ExpirationYear)
	card.CardHolderName = utils.Decrypt(userPassword, card.CardHolderName)
}

func (uc *GophKeeperClientUseCase) loadCards(accessToken string) {
	cards, err := uc.clientAPI.GetCards(accessToken)
	if err != nil {
		color.Red("Connection error: %v", err)

		return
	}

	if err = uc.repo.SaveCards(cards); err != nil {
		log.Println(err)

		return
	}
	color.Green("Loaded %v cards", len(cards))
}

func (uc *GophKeeperClientUseCase) ShowCard(userPassword, cardID string) {
	if !uc.verifyPassword(userPassword) {
		return
	}
	cardUUID, err := uuid.Parse(cardID)
	if err != nil {
		color.Red(err.Error())

		return
	}
	card, err := uc.repo.GetCardByID(cardUUID)
	if err != nil {
		color.Red(err.Error())

		return
	}
	uc.decryptCard(userPassword, &card)
	yellow := color.New(color.FgYellow).SprintFunc()
	fmt.Printf("ID: %s\nname:%s\nCardHolderName:%s\nNumber:%s\nBrand:%s\nExpiration: %s/%s\nCode%s\n%v\n", //nolint:forbidigo // cli printing
		yellow(card.ID),
		yellow(card.Name),
		yellow(card.CardHolderName),
		yellow(card.Number),
		yellow(card.Brand),
		yellow(card.ExpirationMonth),
		yellow(card.ExpirationYear),
		yellow(card.SecurityCode),
		yellow(card.Meta),
	)
}

func (uc *GophKeeperClientUseCase) DelCard(userPassword, cardID string) {
	accessToken, err := uc.authorisationCheck(userPassword)
	if err != nil {
		return
	}
	cardUUID, err := uuid.Parse(cardID)
	if err != nil {
		color.Red(err.Error())
		log.Fatalf("GophKeeperClientUseCase - uuid.Parse - %v", err)
	}

	if err := uc.repo.DelCard(cardUUID); err != nil {
		log.Fatalf("GophKeeperClientUseCase - repo.DelCard - %v", err)
	}

	if err := uc.clientAPI.DelCard(accessToken, cardID); err != nil {
		log.Fatalf("GophKeeperClientUseCase - repo.DelCard - %v", err)
	}

	color.Green("Card %q removed", cardID)
}
