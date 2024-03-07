package TagClientModel

type TagClient struct {
	Tag_id     int    `json:"tag_id" db:"tag_id"`
	ClienteCPF string `json:"cliente_cpf" db:"cliente_cpf"`
}
