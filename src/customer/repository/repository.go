package customerRepository

import (
	customerModel "atividade_4/src/customer/model"
	"context"
	"os"

	"github.com/jackc/pgx/v5"
)

type CustomerRepository struct {
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

func InitCustomerRepository() *CustomerRepository {
	return &CustomerRepository{
		db: InitConnection(),
	}
}

func (repository *CustomerRepository) FindAll() ([]customerModel.Customer, error) {
	rows, _ := repository.db.Query(context.Background(),
		`
    SELECT * FROM customer
    `)

	customers, err := pgx.CollectRows(rows, pgx.RowToStructByName[customerModel.Customer])
	if err != nil {
		return nil, err
	}
	return customers, nil
}

func (repository *CustomerRepository) FindByID(cpf string) (customerModel.Customer, error) {
	var customer customerModel.Customer
	err := repository.db.QueryRow(context.Background(),
		`
    SELECT * FROM customer
    WHERE cpf = $1 
    `, cpf).Scan(&customer)

	if err != nil {
		return customer, err
	}

	return customer, nil
}

func (repository *CustomerRepository) Create(customer customerModel.Customer) error {
	args := pgx.NamedArgs{
		"cpf":             customer.CPF,
		"nome":            customer.Nome,
		"profissao":       customer.Profissao,
		"data_nascimento": customer.DataNascimento,
		"telefone":        customer.Telefone,
		"created_at":      customer.CreatedAt,
		"updated_at":      customer.UpdatedAt,
	}

	_, err := repository.db.Exec(context.Background(),
		`
    INSERT INTO customer
    (cpf, nome, profissao, data_nascimento, telefone, updated_at, created_at) 
    VALUES 
    (@cpf @nome, @profissao, @data_nascimento, @telefone, @updated_at, @created_at)
    `, args)

	return err
}

func (repository *CustomerRepository) Update(customer customerModel.Customer) error {
	args := pgx.NamedArgs{
		"cpf":             customer.CPF,
		"nome":            customer.Nome,
		"profissao":       customer.Profissao,
		"data_nascimento": customer.DataNascimento,
		"telefone":        customer.Telefone,
		"created_at":      customer.CreatedAt,
		"updated_at":      customer.UpdatedAt,
	}

	_, err := repository.db.Exec(context.Background(),
		`
    UPDATE customer
    SET cpf = @cpf, 
        nome = @nome, 
        profissao = @profissao, 
        data_nascimento = @data_nascimento, 
        telefone = @telefone, 
        updated_at = @updated_at
    WHERE cpf = @cpf
    `, args)

	return err
}

func (repository *CustomerRepository) Delete(cpf string) error {
	_, err := repository.db.Exec(context.Background(),
		`
    DELETE FROM customer
    WHERE cpf = $1
    `, cpf)

	return err
}
