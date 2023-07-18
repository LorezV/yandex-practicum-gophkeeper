package models

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Email     string    `gorm:"unique;uniqueIndex;not null"`
	Password  string    `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Cards     []Card   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Logins    []Login  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Notes     []Note   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Bynary    []Binary `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (user *User) ToString() string {
	return fmt.Sprintf("id: %v\nemail: %s", user.ID, user.Email)
}
