package TagCustomerModel

type TagCustomer struct {
	Tag_id      int    `json:"tag_id" db:"tag_id"`
	CustomerCPF string `json:"cliente_cpf" db:"cliente_cpf"`
}
