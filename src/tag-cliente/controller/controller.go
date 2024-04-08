package TagCustomerController

import (
	TagCustomerDTO "atividade_4/src/tag-cliente/dto"
	TagCustomerService "atividade_4/src/tag-cliente/service"
	"log"

	"github.com/gofiber/fiber/v2"
)

type TagCustomerController struct {
	service *TagCustomerService.TagCustomerService
}

func InitTagCustomerController(s *TagCustomerService.TagCustomerService) *TagCustomerController {
	return &TagCustomerController{
		service: s,
	}
}

func (controller *TagCustomerController) HandleFindAll(c *fiber.Ctx) error {
	TagCustomers, err := controller.service.FindAll()
	if err != nil {
		log.Println(err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	return c.Status(fiber.StatusOK).JSON(TagCustomers)
}

func (controller *TagCustomerController) HandleFindByCustomer(c *fiber.Ctx) error {
	CustomerCPF := c.Params("cliente_cpf")

	TagCustomer, err := controller.service.FindByCustomer(CustomerCPF)
	if err != nil {
		log.Println()
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusOK).JSON(TagCustomer)
}

func (controller *TagCustomerController) HandleFindByTag(c *fiber.Ctx) error {
	TagID := c.Params("tag_id")
	TagCustomer, err := controller.service.FindByTag(TagID)
	if err != nil {
		log.Println()
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusOK).JSON(TagCustomer)
}

func (controller *TagCustomerController) HandleCreateTagCustomer(c *fiber.Ctx) error {
	var TagCustomerDTO TagCustomerDTO.TagCustomerCreateDTO

	err := c.BodyParser(&TagCustomerDTO)
	if err != nil {
		log.Println(err)
		return c.SendStatus(fiber.StatusBadRequest)
	}

	err = controller.service.Create(TagCustomerDTO)
	if err != nil {
		log.Println(err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":      "OK",
		"cliente_cpf": TagCustomerDTO.CustomerCPF,
		"tag_id":      TagCustomerDTO.Tag_id.String(),
	})
}

func (controller *TagCustomerController) HandleDeleteTagCustomer(c *fiber.Ctx) error {
	CustomereCPF := c.Params("cliente_cpf")
	TagID := c.Params("tag_id")

	err := controller.service.Delete(CustomereCPF, TagID)
	if err != nil {
		log.Println(err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	return c.SendStatus(fiber.StatusOK)
}
