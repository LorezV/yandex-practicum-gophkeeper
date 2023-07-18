package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MetaCard struct {
	gorm.Model
	ID     uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Name   string
	Value  string
	CardID uuid.UUID
}
type Card struct {
	gorm.Model
	ID              uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Name            string    `gorm:"size:100"`
	CardHolderName  string
	Number          string
	Brand           string
	ExpirationMonth string
	ExpirationYear  string
	SecurityCode    string
	UserID          uuid.UUID
	Meta            []MetaCard `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
