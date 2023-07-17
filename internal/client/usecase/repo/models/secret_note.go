package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MetaNote struct {
	gorm.Model
	ID     uuid.UUID
	Name   string
	Value  string
	NoteID uuid.UUID
}
type Note struct {
	gorm.Model
	ID     uuid.UUID `gorm:"type:uuid;primary_key"`
	Name   string    `gorm:"size:100"`
	Note   string
	UserID uint
	Meta   []MetaNote `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
