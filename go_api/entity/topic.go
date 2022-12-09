package entity

import "github.com/google/uuid"

type Topic struct {
	ID uuid.UUID
	Name string
	Description string
}
