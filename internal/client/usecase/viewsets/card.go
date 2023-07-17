package viewsets

import "github.com/google/uuid"

type CardForList struct {
	ID    uuid.UUID
	Name  string
	Brand string
}
