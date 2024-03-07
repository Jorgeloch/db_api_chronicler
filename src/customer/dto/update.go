package customerDTO

import (
	"time"

	"github.com/google/uuid"
)

type UpdateCustomerDTO struct {
	Nome           string    `validate:"max=20" json:"nome,omitempty" db:"nome"`
	Profissao      uuid.UUID `json:"profissao,omitempty" db:"profissao"`
	DataNascimento time.Time `json:"data_nascimento,omitempty" db:"data_nascimento"`
	Telefone       []string  `json:"telefone,omitempty" db:"telefone"`
}
