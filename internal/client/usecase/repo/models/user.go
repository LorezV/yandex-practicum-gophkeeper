package models

import (
	"fmt"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email        string `gorm:"unique;uniqueIndex;not null"`
	Password     string `gorm:"not null"`
	AccessToken  string
	RefreshToken string
	Cards        []Card  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Logins       []Login `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Notes        []Note  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Binary       []Note  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (user *User) ToString() string {
	return fmt.Sprintf("id: %v\nemail: %s", user.ID, user.Email)
}
