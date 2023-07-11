package entity

import "github.com/google/uuid"

type Login struct {
	ID       uuid.UUID `json:"uuid" swaggerignore:"true"`
	Name     string    `json:"name"`
	Login    string    `json:"login"`
	Password string    `json:"password"`
	URI      string    `json:"uri"`
	Meta     []Meta    `json:"meta"`
}
