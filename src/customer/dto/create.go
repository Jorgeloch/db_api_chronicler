package customerDTO

import (
	"time"

	"github.com/google/uuid"
	"github.com/klassmann/cpfcnpj"
)

type CreateCustomerDTO struct {
	CPF            string    `validate:"required,max=11" json:"cpf" db:"cpf"`
	Nome           string    `validate:"required,max=20" json:"nome" db:"nome"`
	Profissao      uuid.UUID `validate:"required" json:"profissao" db:"profissao"`
	DataNascimento time.Time `validate:"required" json:"data_nascimento" db:"data_nascimento"`
	Telefone       []string  `validate:"required" json:"telefone" db:"telefone"`
}

func (dto *CreateCustomerDTO) ValidateCPF() bool {
	return cpfcnpj.ValidateCPF(dto.CPF)
}
