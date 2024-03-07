package tagRepository

import (
	tagModel "atividade_4/src/tag/model"
	"context"
	"os"

	"github.com/jackc/pgx/v5"
)

type TagRepository struct {
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

func InitTagRepository() *TagRepository {
	return &TagRepository{
		db: InitConnection(),
	}
}

func (repository *TagRepository) FindAll() ([]tagModel.Tag, error) {
	rows, _ := repository.db.Query(context.Background(),
		`
    SELECT * FROM tag
    `)

	tags, err := pgx.CollectRows(rows, pgx.RowToStructByName[tagModel.Tag])
	if err != nil {
		return nil, err
	}
	return tags, nil
}

func (repository *TagRepository) FindByID(id string) (tagModel.Tag, error) {
	var tag tagModel.Tag
	err := repository.db.QueryRow(context.Background(),
		`
    SELECT * FROM tag
    WHERE id = $1 
    `, id).Scan(&tag.ID, &tag.Nome, &tag.Cor)

	if err != nil {
		return tag, err
	}

	return tag, nil
}

func (repository *TagRepository) Create(tag tagModel.Tag) error {
	args := pgx.NamedArgs{
		"id":   tag.ID,
		"nome": tag.Nome,
		"cor":  tag.Cor,
	}

	_, err := repository.db.Exec(context.Background(),
		`
    INSERT INTO tag
    (id, nome, cor) 
    VALUES 
    (@id, @nome, @cor)
    `, args)

	return err
}

func (repository *TagRepository) Update(tag tagModel.Tag) error {
	args := pgx.NamedArgs{
		"id":   tag.ID,
		"nome": tag.Nome,
		"cor":  tag.Cor,
	}

	_, err := repository.db.Exec(context.Background(),
		`
    UPDATE tag
    SET nome = @nome, cor = @cor
    WHERE id = @id
    `, args)

	return err
}

func (repository *TagRepository) Delete(id string) error {
	_, err := repository.db.Exec(context.Background(),
		`
    DELETE FROM tag
    WHERE id = $1
    `, id)

	return err
}
