package viewsets

import "github.com/google/uuid"

type BinaryForList struct {
	ID       uuid.UUID
	Name     string
	FileName string
}
