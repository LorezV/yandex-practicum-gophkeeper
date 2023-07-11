package errs

import (
	"encoding/json"
	"errors"
)

var (
	ErrWrongEmail           = errors.New("incorrect email given")
	ErrEmailAlreadyExists   = errors.New("given email already exists")
	ErrWrongCredentials     = errors.New("wrong credentials have been given")
	ErrTokenValidation      = errors.New("token validation error")
	ErrUnexpectedError      = errors.New("some unexpected error")
	ErrWrongOwnerOrNotFound = errors.New("wrong owner or not found")
)

type GormErr struct {
	Code    string `json:"Code"`
	Message string `json:"Message"`
}

func ParsePostgresErr(dbErr error) (newError GormErr) {
	byteErr, err := json.Marshal(dbErr)
	if err != nil {
		return
	}

	if err = json.Unmarshal((byteErr), &newError); err != nil {
		return GormErr{}
	}

	return
}

func ParseServerError(body []byte) string {
	var errMessage struct {
		Message string `json:"error"`
	}

	if err := json.Unmarshal(body, &errMessage); err == nil {
		return errMessage.Message
	}

	return ""
}
