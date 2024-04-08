package TagCustomerModel

import "github.com/google/uuid"

type TagCustomer struct {
	Tag_id      uuid.UUID `json:"tag_id" db:"tag_id"`
	CustomerCPF string    `json:"cliente_cpf" db:"cliente_cpf"`
}
