package repo

import (
	"errors"

	"github.com/LorezV/gophkeeper/internal/client/usecase/repo/models"
	"github.com/LorezV/gophkeeper/internal/client/usecase/viewsets"
	"github.com/LorezV/gophkeeper/internal/entity"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

var errCardNotFound = errors.New("card not found")

func (r *GophKeeperRepo) AddCard(card *entity.Card) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		cardForSaving := models.Card{
			ID:              card.ID,
			Brand:           card.Brand,
			Name:            card.Name,
			Number:          card.Number,
			SecurityCode:    card.SecurityCode,
			CardHolderName:  card.CardHolderName,
			ExpirationMonth: card.ExpirationMonth,
			ExpirationYear:  card.ExpirationYear,
			UserID:          r.getUserID(),
		}
		if err := tx.Save(&cardForSaving).Error; err != nil {
			return nil
		}
		for _, meta := range card.Meta {
			metaForCard := models.MetaCard{
				Name:   meta.Name,
				Value:  meta.Value,
				CardID: cardForSaving.ID,
				ID:     meta.ID,
			}
			if err := tx.Create(&metaForCard).Error; err != nil {
				return err
			}
		}

		return nil
	})
}

func (r *GophKeeperRepo) SaveCards(cards []entity.Card) error {
	if len(cards) == 0 {
		return nil
	}
	userID := r.getUserID()
	cardsForDB := make([]models.Card, len(cards))
	for index := range cards {
		cardsForDB[index].ID = cards[index].ID
		cardsForDB[index].Brand = cards[index].Brand
		cardsForDB[index].CardHolderName = cards[index].CardHolderName
		cardsForDB[index].ExpirationMonth = cards[index].ExpirationMonth
		cardsForDB[index].ExpirationYear = cards[index].ExpirationYear
		cardsForDB[index].Name = cards[index].Name
		cardsForDB[index].Number = cards[index].Number
		cardsForDB[index].SecurityCode = cards[index].SecurityCode
		cardsForDB[index].UserID = userID
		for _, meta := range cards[index].Meta {
			cardsForDB[index].Meta = append(cardsForDB[index].Meta, models.MetaCard{
				Name:   meta.Name,
				Value:  meta.Value,
				CardID: cards[index].ID,
				ID:     meta.ID,
			})
		}
	}

	return r.db.Save(cardsForDB).Error
}

func (r *GophKeeperRepo) LoadCards() []viewsets.CardForList {
	userID := r.getUserID()
	var cards []models.Card
	r.db.Model(&models.Card{}).
		Preload("Meta").
		Where("user_id", userID).
		Find(&cards)
	if len(cards) == 0 {
		return nil
	}

	cardsViewSet := make([]viewsets.CardForList, len(cards))

	for index := range cards {
		cardsViewSet[index].ID = cards[index].ID
		cardsViewSet[index].Name = cards[index].Name
		cardsViewSet[index].Brand = cards[index].Brand
	}

	return cardsViewSet
}

func (r *GophKeeperRepo) GetCardByID(cardID uuid.UUID) (card entity.Card, err error) {
	var cardFromDB models.Card
	if err = r.db.
		Model(&models.Card{}).
		Preload("Meta").
		Find(&cardFromDB, cardID).Error; cardFromDB.ID == uuid.Nil || err != nil {
		return card, errCardNotFound
	}

	card.ID = cardFromDB.ID
	card.Brand = cardFromDB.Brand
	card.Number = cardFromDB.Number
	card.Name = cardFromDB.Name
	card.CardHolderName = cardFromDB.CardHolderName
	card.ExpirationMonth = cardFromDB.ExpirationMonth
	card.ExpirationYear = cardFromDB.ExpirationYear

	for index := range cardFromDB.Meta {
		card.Meta = append(card.Meta, entity.Meta{
			ID:    cardFromDB.Meta[index].ID,
			Name:  cardFromDB.Meta[index].Name,
			Value: cardFromDB.Meta[index].Value,
		})
	}

	return
}

func (r *GophKeeperRepo) DelCard(cardID uuid.UUID) error {
	return r.db.Unscoped().Delete(&models.Card{}, cardID).Error
}
