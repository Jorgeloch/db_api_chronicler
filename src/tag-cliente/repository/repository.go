package TagCustomerRepository

import (
	TagCustomerModel "atividade_4/src/tag-cliente/model"
	"context"
	"os"

	"github.com/jackc/pgx/v5"
)

type TagCustomerRepository struct {
	db *pgx.Conn
}

func InitConnection() *pgx.Conn {
	URL := os.Getenv("DATABASE_URL")
	db, err := pgx.Connect(context.Background(), URL)
	if err != nil {
		panic(err)
	}
	return db
}

func InitTagRepository() *TagCustomerRepository {
	return &TagCustomerRepository{
		db: InitConnection(),
	}
}

func (repository *TagCustomerRepository) FindAll() ([]TagCustomerModel.TagCustomer, error) {
	rows, _ := repository.db.Query(context.Background(),
		`
    SELECT * FROM cliente_has_tag
    `)

	tags, err := pgx.CollectRows(rows, pgx.RowToStructByName[TagCustomerModel.TagCustomer])
	if err != nil {
		return nil, err
	}
	return tags, nil
}

func (repository *TagCustomerRepository) FindByCustomer(cpf string) ([]TagCustomerModel.TagCustomer, error) {
	rows, _ := repository.db.Query(context.Background(),
		`
    SELECT * FROM cliente_has_tag
    WHERE cliente_cpf = $1 
    `, cpf)

	tags, err := pgx.CollectRows(rows, pgx.RowToStructByName[TagCustomerModel.TagCustomer])

	if err != nil {
		return tags, err
	}
	return tags, nil
}

func (repository *TagCustomerRepository) Create(tagCustomer TagCustomerModel.TagCustomer) error {
	args := pgx.NamedArgs{
		"cliente_cpf": tagCustomer.CustomerCPF,
		"tag_id":      tagCustomer.Tag_id,
	}
	_, err := repository.db.Exec(context.Background(),
		`
    INSERT INTO cliente_has_tag
    (cliente_cpf, tag_id) 
    VALUES 
    (@cliente_cpf, @tag_id)
    `, args)

	return err
}

func (repository *TagCustomerRepository) Delete(CustomerCPF string, TagID string) error {
	_, err := repository.db.Exec(context.Background(),
		`
    DELETE FROM cliente_has_tag
    WHERE cliente_cpf = $1 AND tag_id = $2
    `, CustomerCPF, TagID)
	if err != nil {
		return err
	}
	return nil
}
