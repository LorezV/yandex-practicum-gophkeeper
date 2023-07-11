package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MetaLogin struct {
	gorm.Model
	ID      uuid.UUID
	Name    string
	Value   string
	LoginID uuid.UUID
}
type Login struct {
	gorm.Model
	ID       uuid.UUID `gorm:"type:uuid;primary_key"`
	Name     string    `gorm:"size:100"`
	URI      string
	Login    string
	Password string
	UserID   uint
	Meta     []MetaLogin `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
