package viewsets

import "github.com/google/uuid"

type NoteForList struct {
	ID   uuid.UUID
	Name string
}
