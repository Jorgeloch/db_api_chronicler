package customerRepository

import (
	customerModel "atividade_4/src/customer/model"
	"context"

	"github.com/jackc/pgx/v5"
)

type CustomerRepository struct {
	db *pgx.Conn
}

func InitCustomerRepository(db *pgx.Conn) *CustomerRepository {
	return &CustomerRepository{
		db: db,
	}
}

func (repository *CustomerRepository) FindAll() ([]customerModel.Customer, error) {
	rows, _ := repository.db.Query(context.Background(),
		`
    SELECT * FROM cliente
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
    SELECT cpf, nome, telefone, data_nascimento FROM cliente
    WHERE cpf = $1 
    `, cpf).Scan(&customer.CPF, &customer.Nome, &customer.Telefone, &customer.DataNascimento)

	if err != nil {
		return customer, err
	}

	return customer, nil
}

func (repository *CustomerRepository) Create(customer customerModel.Customer) error {
	args := pgx.NamedArgs{
		"cpf":             customer.CPF,
		"nome":            customer.Nome,
		"data_nascimento": customer.DataNascimento,
		"telefone":        customer.Telefone,
	}

	_, err := repository.db.Exec(context.Background(),
		`
    INSERT INTO cliente
    (cpf, nome, data_nascimento, telefone)
    VALUES 
    (@cpf, @nome, @data_nascimento, @telefone)
    `, args)

	return err
}

func (repository *CustomerRepository) Update(customer customerModel.Customer) error {
	args := pgx.NamedArgs{
		"cpf":             customer.CPF,
		"nome":            customer.Nome,
		"data_nascimento": customer.DataNascimento,
		"telefone":        customer.Telefone,
	}

	_, err := repository.db.Exec(context.Background(),
		`
    UPDATE cliente
    SET cpf = @cpf, 
        nome = @nome, 
        data_nascimento = @data_nascimento, 
        telefone = @telefone 
    WHERE cpf = @cpf
    `, args)

	return err
}

func (repository *CustomerRepository) Delete(cpf string) error {
	_, err := repository.db.Exec(context.Background(),
		`
    DELETE FROM cliente
    WHERE cpf = $1
    `, cpf)

	return err
}
