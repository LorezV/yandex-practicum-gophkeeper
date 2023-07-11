package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MetaBinary struct {
	gorm.Model
	ID       uuid.UUID `gorm:"type:uuid;primary_key"`
	Name     string
	Value    string
	BinaryID uuid.UUID `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type Binary struct {
	gorm.Model
	ID       uuid.UUID `gorm:"type:uuid;primary_key"`
	Name     string
	FileName string
	UserID   uint
	Meta     []MetaBinary
}
