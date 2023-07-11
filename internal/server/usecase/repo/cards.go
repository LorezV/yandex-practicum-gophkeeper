package repo

import (
	"context"

	"github.com/LorezV/gophkeeper/internal/entity"
	"github.com/LorezV/gophkeeper/internal/server/usecase/repo/models"
	"github.com/LorezV/gophkeeper/internal/utils/errs"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (r *GophKeeperRepo) GetCards(ctx context.Context, user entity.User) ([]entity.Card, error) {
	var cardsFromDB []models.Card

	err := r.db.WithContext(ctx).
		Model(&models.Card{}).
		Preload("Meta").
		Find(&cardsFromDB, "user_id = ?", user.ID).Error
	if err != nil {
		return nil, err
	}

	if len(cardsFromDB) == 0 {
		return nil, nil
	}

	cards := make([]entity.Card, len(cardsFromDB))

	for index := range cardsFromDB {
		cards[index].ID = cardsFromDB[index].ID
		cards[index].Brand = cardsFromDB[index].Brand
		cards[index].CardHolderName = cardsFromDB[index].CardHolderName
		cards[index].ExpirationMonth = cardsFromDB[index].ExpirationMonth
		cards[index].ExpirationYear = cardsFromDB[index].ExpirationYear
		cards[index].Name = cardsFromDB[index].Name
		cards[index].Number = cardsFromDB[index].Number
		cards[index].SecurityCode = cardsFromDB[index].SecurityCode
	}

	return cards, nil
}

func (r *GophKeeperRepo) AddCard(ctx context.Context, card *entity.Card, userID uuid.UUID) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		cardToDB := models.Card{
			ID:              uuid.New(),
			UserID:          userID,
			Name:            card.Name,
			Brand:           card.Brand,
			CardHolderName:  card.CardHolderName,
			Number:          card.Number,
			ExpirationMonth: card.ExpirationMonth,
			ExpirationYear:  card.ExpirationYear,
			SecurityCode:    card.SecurityCode,
		}

		if err := tx.WithContext(ctx).Create(&cardToDB).Error; err != nil {
			return err
		}
		card.ID = cardToDB.ID
		for index, meta := range card.Meta {
			metaForCard := models.MetaCard{
				Name:   meta.Name,
				Value:  meta.Value,
				CardID: cardToDB.ID,
			}
			if err := tx.WithContext(ctx).Create(&metaForCard).Error; err != nil {
				return err
			}
			card.Meta[index].ID = metaForCard.ID
		}

		return nil
	})
}

func (r *GophKeeperRepo) IsCardOwner(ctx context.Context, cardUUID, userID uuid.UUID) bool {
	var cardFromDB models.Card

	r.db.WithContext(ctx).Where("id = ?", cardUUID).First(&cardFromDB)

	return cardFromDB.UserID == userID
}

func (r *GophKeeperRepo) DelCard(ctx context.Context, cardUUID, userID uuid.UUID) error {
	if !r.IsCardOwner(ctx, cardUUID, userID) {
		return errs.ErrWrongOwnerOrNotFound
	}

	return r.db.WithContext(ctx).Delete(&models.Card{}, cardUUID).Error
}

func (r *GophKeeperRepo) UpdateCard(ctx context.Context, card *entity.Card, userID uuid.UUID) error {
	if !r.IsCardOwner(ctx, card.ID, userID) {
		return errs.ErrWrongOwnerOrNotFound
	}

	return r.db.Transaction(func(tx *gorm.DB) error {
		cardToDB := models.Card{
			ID:              card.ID,
			UserID:          userID,
			Name:            card.Name,
			Brand:           card.Brand,
			CardHolderName:  card.CardHolderName,
			Number:          card.Number,
			ExpirationMonth: card.ExpirationMonth,
			ExpirationYear:  card.ExpirationYear,
			SecurityCode:    card.SecurityCode,
		}
		if err := tx.WithContext(ctx).Save(&cardToDB).Error; err != nil {
			return err
		}
		for _, meta := range card.Meta {
			metaForCard := models.MetaCard{
				Name:   meta.Name,
				Value:  meta.Value,
				CardID: cardToDB.ID,
				ID:     meta.ID,
			}
			if err := tx.WithContext(ctx).Create(&metaForCard).Error; err != nil {
				return err
			}
		}

		return nil
	})
}
