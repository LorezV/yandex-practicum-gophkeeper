package entity

import "github.com/google/uuid"

type Card struct {
	ID              uuid.UUID `json:"uuid" swaggerignore:"true"`
	Name            string    `json:"name"`
	CardHolderName  string    `json:"card_holder_hame"`
	Number          string    `json:"number"`
	Brand           string    `json:"brand"`
	ExpirationMonth string    `json:"expiration_month"`
	ExpirationYear  string    `json:"expiration_year"`
	SecurityCode    string    `json:"security_code"`
	Meta            []Meta    `json:"meta"`
}
