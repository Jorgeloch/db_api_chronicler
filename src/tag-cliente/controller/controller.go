package TagClientController

import (
	TagClientDTO "atividade_4/src/tag-cliente/dto"
	TagClientService "atividade_4/src/tag-cliente/service"

	"github.com/gofiber/fiber/v2"
)

type TagClientController struct {
	service *TagClientService.TagClientService
}

func InitTagClientController(s *TagClientService.TagClientService) *TagClientController {
	return &TagClientController{
		service: s,
	}
}

func (controller *TagClientController) HandleFindAll(c *fiber.Ctx) error {
	TagClients, err := controller.service.FindAll()
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	return c.Status(fiber.StatusOK).JSON(TagClients)
}

func (controller *TagClientController) HandleFindByClient(c *fiber.Ctx) error {
	ClienteCPF := c.Params("cliente_cpf")

	TagClient, err := controller.service.FindByClient(ClienteCPF)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusOK).JSON(TagClient)
}

func (controller *TagClientController) HandleCreateTagClient(c *fiber.Ctx) error {
	var TagClientDTO TagClientDTO.TagClientCreateDTO

	err := c.BodyParser(&TagClientDTO)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	err = controller.service.Create(TagClientDTO)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.SendStatus(fiber.StatusOK)
}

func (controller *TagClientController) HandleDeleteTagClient(c *fiber.Ctx) error {
	ClienteCPF := c.Params("cliente_cpf")
	TagID := c.Params("tag_id")

	err := controller.service.Delete(ClienteCPF, TagID)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	return c.SendStatus(fiber.StatusOK)
}
