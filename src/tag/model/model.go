package tagModel

import (
	"github.com/google/uuid"
)

type Tag struct {
	ID   uuid.UUID `json:"id" db:"id"`
	Nome string    `json:"nome" db:"nome"`
	Cor  string    `json:"cor" db:"cor"`
}
