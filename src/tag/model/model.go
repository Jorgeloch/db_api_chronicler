package tagModel

import (
	"time"

	"github.com/google/uuid"
)

type Tag struct {
	Id        uuid.UUID `json:"id" db:"id"`
	Nome      string    `json:"nome" db:"nome"`
	Cor       string    `json:"cor" db:"cor"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}
