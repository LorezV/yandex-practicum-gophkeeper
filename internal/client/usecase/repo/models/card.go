package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MetaCard struct {
	gorm.Model
	ID     uuid.UUID
	Name   string
	Value  string
	CardID uuid.UUID
}

type Card struct {
	gorm.Model
	ID              uuid.UUID `gorm:"type:uuid;primary_key"`
	Name            string    `gorm:"size:100"`
	CardHolderName  string
	Number          string
	Brand           string
	ExpirationMonth string
	ExpirationYear  string
	SecurityCode    string
	UserID          uint
	Meta            []MetaCard `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
