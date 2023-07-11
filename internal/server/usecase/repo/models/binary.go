package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MetaBinary struct {
	gorm.Model
	ID       uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Name     string
	Value    string
	BinaryID uuid.UUID `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type Binary struct {
	gorm.Model
	ID       uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Name     string
	FileName string
	UserID   uuid.UUID
	Meta     []MetaBinary
}
