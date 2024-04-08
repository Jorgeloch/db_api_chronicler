package customerModel

import (
	"time"
)

type Customer struct {
	CPF            string    `json:"cpf" db:"cpf"`
	Nome           string    `json:"nome" db:"nome"`
	DataNascimento time.Time `json:"data_nascimento" db:"data_nascimento"`
	Telefone       []string  `json:"telefone" db:"telefone"`
}
