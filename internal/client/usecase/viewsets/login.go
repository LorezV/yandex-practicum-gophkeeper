package viewsets

import "github.com/google/uuid"

type LoginForList struct {
	ID   uuid.UUID
	Name string
	URI  string
}
