package customerService

import (
	"time"

	customerDTO "atividade_4/src/customer/dto"
	customerModel "atividade_4/src/customer/model"
	customerRepository "atividade_4/src/customer/repository"
	"github.com/google/uuid"
)

type CustomerService struct {
	repository *customerRepository.CustomerRepository
}

func InitCustomerService(r *customerRepository.CustomerRepository) *CustomerService {
	return &CustomerService{
		repository: r,
	}
}

func (service *CustomerService) FindByID(id string) (customerModel.Customer, error) {
	return service.repository.FindByID(id)
}

func (service *CustomerService) Create(dto customerDTO.CreateCustomerDTO) (string, error) {
	if !dto.ValidateCPF() {
		return "", nil
	}

	model := customerModel.Customer{
		CPF:            dto.CPF,
		Nome:           dto.Nome,
		Profissao:      dto.Profissao,
		DataNascimento: dto.DataNascimento,
		Telefone:       dto.Telefone,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	err := service.repository.Create(model)

	return model.CPF, err
}

func (service *CustomerService) Update(cpf string, dto customerDTO.UpdateCustomerDTO) (customerModel.Customer, error) {
	updatedCustomer, err := service.repository.FindByID(cpf)

	if err != nil {
		return updatedCustomer, err
	}

	if dto.Nome != "" {
		updatedCustomer.Nome = dto.Nome
	}
	if dto.Profissao != uuid.Nil {
		updatedCustomer.Profissao = dto.Profissao
	}
	if dto.DataNascimento.IsZero() {
		updatedCustomer.DataNascimento = dto.DataNascimento
	}
	if dto.Telefone != nil {
		updatedCustomer.Telefone = dto.Telefone
	}

	updatedCustomer.UpdatedAt = time.Now()

	err = service.repository.Update(updatedCustomer)

	return updatedCustomer, err
}

func (service *CustomerService) FindAll() ([]customerModel.Customer, error) {
	return service.repository.FindAll()
}

func (service *CustomerService) Delete(id string) error {
	return service.repository.Delete(id)
}
