package customerService

import (
	customerDTO "atividade_4/src/customer/dto"
	customerModel "atividade_4/src/customer/model"
	customerRepository "atividade_4/src/customer/repository"
)

type CustomerService struct {
	repository *customerRepository.CustomerRepository
}

func InitCustomerService(r *customerRepository.CustomerRepository) *CustomerService {
	return &CustomerService{
		repository: r,
	}
}

func (service *CustomerService) FindByID(cpf string) (customerModel.Customer, error) {
	return service.repository.FindByID(cpf)
}

func (service *CustomerService) Create(dto customerDTO.CreateCustomerDTO) (string, error) {
	if !dto.ValidateCPF() {
		return "", nil
	}

	model := customerModel.Customer{
		CPF:            dto.CPF,
		Nome:           dto.Nome,
		DataNascimento: dto.DataNascimento,
		Telefone:       dto.Telefone,
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
	if dto.DataNascimento.IsZero() {
		updatedCustomer.DataNascimento = dto.DataNascimento
	}
	if dto.Telefone != nil {
		updatedCustomer.Telefone = dto.Telefone
	}

	err = service.repository.Update(updatedCustomer)

	return updatedCustomer, err
}

func (service *CustomerService) FindAll() ([]customerModel.Customer, error) {
	return service.repository.FindAll()
}

func (service *CustomerService) Delete(id string) error {
	return service.repository.Delete(id)
}
