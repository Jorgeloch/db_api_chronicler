package TagClientService

import (
	TagClientDTO "atividade_4/src/tag-cliente/dto"
	TagClientModel "atividade_4/src/tag-cliente/model"
	TagClientRepository "atividade_4/src/tag-cliente/repository"
)

type TagClientService struct {
	repository *TagClientRepository.TagClientRepository
}

func InitTagClientService(repository *TagClientRepository.TagClientRepository) *TagClientService {
	return &TagClientService{
		repository: repository,
	}
}

func (service *TagClientService) FindByClient(ClienteCPF string) ([]TagClientModel.TagClient, error) {
	return service.repository.FindByClient(ClienteCPF)
}

func (service *TagClientService) Create(dto TagClientDTO.TagClientCreateDTO) error {
	model := TagClientModel.TagClient{
		ClienteCPF: dto.ClienteCPF,
		Tag_id:     dto.Tag_id,
	}

	err := service.repository.Create(model)

	return err
}

func (service *TagClientService) FindAll() ([]TagClientModel.TagClient, error) {
	return service.repository.FindAll()
}

func (service *TagClientService) Delete(ClienteCPF string, TagID string) error {
	return service.repository.Delete(ClienteCPF, TagID)
}
