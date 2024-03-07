package TagClientRepository

import (
	TagClientModel "atividade_4/src/tag-cliente/model"
	"context"
	"os"

	"github.com/jackc/pgx/v5"
)

type TagClientRepository struct {
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

func InitTagRepository() *TagClientRepository {
	return &TagClientRepository{
		db: InitConnection(),
	}
}

func (repository *TagClientRepository) FindAll() ([]TagClientModel.TagClient, error) {
	rows, _ := repository.db.Query(context.Background(),
		`
    SELECT * FROM tag_has_client
    `)

	tags, err := pgx.CollectRows(rows, pgx.RowToStructByName[TagClientModel.TagClient])
	if err != nil {
		return nil, err
	}
	return tags, nil
}

func (repository *TagClientRepository) FindByClient(cpf string) ([]TagClientModel.TagClient, error) {
	rows, _ := repository.db.Query(context.Background(),
		`
    SELECT * FROM tag_has_client
    WHERE cliente_cpf = $1 
    `, cpf)

	tags, err := pgx.CollectRows(rows, pgx.RowToStructByName[TagClientModel.TagClient])

	if err != nil {
		return tags, err
	}
	return tags, nil
}

func (repository *TagClientRepository) Create(tagClient TagClientModel.TagClient) error {
	args := pgx.NamedArgs{
		"cliente_cpf": tagClient.ClienteCPF,
		"tag_id":      tagClient.Tag_id,
	}
	_, err := repository.db.Exec(context.Background(),
		`
    INSERT INTO tag_has_client
    (cliente_cpf, tag_id) 
    VALUES 
    (@cliente_cpf, @tag_id)
    `, args)

	return err
}

func (repository *TagClientRepository) Delete(ClienteCPF string, TagID string) error {
	_, err := repository.db.Exec(context.Background(),
		`
    DELETE FROM tag_has_client
    WHERE cliente_cpf = $1 AND tag_id = $2
    `, ClienteCPF, TagID)
	if err != nil {
		return err
	}
	return nil
}
