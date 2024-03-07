package tagController

import (
	tagDTO "atividade_4/src/tag/dto"
	tagService "atividade_4/src/tag/service"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type TagController struct {
	service *tagService.TagService
}

func InitTagController(s *tagService.TagService) *TagController {
	return &TagController{
		service: s,
	}
}

func (controller *TagController) HandleFindAll(c *fiber.Ctx) error {
	tags, err := controller.service.FindAll()
	if err != nil {
		log.Println(err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	return c.Status(fiber.StatusOK).JSON(tags)
}

func (controller *TagController) HandleFindByID(c *fiber.Ctx) error {
	id := c.Params("id")

	tag, err := controller.service.FindByID(id)
	if err != nil {
		log.Println(err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusOK).JSON(tag)
}

func (controller *TagController) HandleCreateTag(c *fiber.Ctx) error {
	var tagDTO tagDTO.CreateTagDTO

	err := c.BodyParser(&tagDTO)
	if err != nil {
		log.Println(err)
		return c.SendStatus(fiber.StatusBadRequest)
	}

	newTagID, err := controller.service.Create(tagDTO)
	if err != nil {
		log.Println(err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	if newTagID == uuid.Nil {
		log.Println("falha na verificacao da cor hexadecimal")
		return c.SendStatus(fiber.StatusBadRequest)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"location": newTagID,
	})
}

func (controller *TagController) HandleUpdateTag(c *fiber.Ctx) error {
	var tagDTO tagDTO.UpdateTagDTO

	err := c.BodyParser(&tagDTO)
	if err != nil {
		log.Println(err)
		return c.SendStatus(fiber.StatusBadRequest)
	}

	id := c.Params("id")

	tagUpdated, err := controller.service.Update(id, tagDTO)
	if err != nil {
		log.Println(err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusOK).JSON(tagUpdated)
}

func (controller *TagController) HandleDeleteTag(c *fiber.Ctx) error {
	id := c.Params("id")

	err := controller.service.Delete(id)
	if err != nil {
		log.Println(err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	return c.SendStatus(fiber.StatusOK)
}
