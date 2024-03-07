package TagCustomerService

import (
	TagCustomerDTO "atividade_4/src/tag-cliente/dto"
	TagCustomerModel "atividade_4/src/tag-cliente/model"
	TagCustomerRepository "atividade_4/src/tag-cliente/repository"
)

type TagCustomerService struct {
	repository *TagCustomerRepository.TagCustomerRepository
}

func InitTagCustomerService(repository *TagCustomerRepository.TagCustomerRepository) *TagCustomerService {
	return &TagCustomerService{
		repository: repository,
	}
}

func (service *TagCustomerService) FindByCustomer(CustomereCPF string) ([]TagCustomerModel.TagCustomer, error) {
	return service.repository.FindByCustomer(CustomereCPF)
}

func (service *TagCustomerService) Create(dto TagCustomerDTO.TagCustomerCreateDTO) error {
	model := TagCustomerModel.TagCustomer{
		CustomerCPF: dto.CustomerCPF,
		Tag_id:      dto.Tag_id,
	}

	err := service.repository.Create(model)

	return err
}

func (service *TagCustomerService) FindAll() ([]TagCustomerModel.TagCustomer, error) {
	return service.repository.FindAll()
}

func (service *TagCustomerService) Delete(CustomereCPF string, TagID string) error {
	return service.repository.Delete(CustomereCPF, TagID)
}
