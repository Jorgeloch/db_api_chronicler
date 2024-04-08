package TagCustomerDTO

import "github.com/google/uuid"

type TagCustomerCreateDTO struct {
	Tag_id      uuid.UUID `json:"tag_id"`
	CustomerCPF string    `json:"cliente_cpf"`
}
