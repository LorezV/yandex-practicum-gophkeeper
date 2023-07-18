package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MetaLogin struct {
	gorm.Model
	ID      uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Name    string
	Value   string
	LoginID uuid.UUID
}
type Login struct {
	gorm.Model
	ID       uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Name     string    `gorm:"size:100"`
	URI      string
	Login    string
	Password string
	UserID   uuid.UUID
	Meta     []MetaLogin `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
