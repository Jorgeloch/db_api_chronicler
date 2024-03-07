package customerModel

import (
	"time"

	"github.com/google/uuid"
)

type Customer struct {
	CPF            string    `json:"cpf" db:"cpf"`
	Nome           string    `json:"nome" db:"nome"`
	Profissao      uuid.UUID `json:"profissao" db:"profissao"`
	DataNascimento time.Time `json:"data_nascimento" db:"data_nascimento"`
	Telefone       []string  `json:"telefone" db:"telefone"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time `json:"updated_at" db:"updated_at"`
}
