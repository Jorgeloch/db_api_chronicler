package tagRouter

import (
	tagController "atividade_4/src/tag/controller"
	tagRepository "atividade_4/src/tag/repository"
	tagService "atividade_4/src/tag/service"
	"github.com/gofiber/fiber/v2"
)

func NewTagRouter() *fiber.App {
	router := fiber.New()

	repository := tagRepository.InitTagRepository()
	service := tagService.InitTagService(repository)
	controller := tagController.InitTagController(service)

	router.Get("/", controller.HandleFindAll)
	router.Get(":id", controller.HandleFindByID)
	router.Post("/", controller.HandleCreateTag)
	router.Patch("/", controller.HandleUpdateTag)
	router.Delete("/", controller.HandleDeleteTag)

	return router
}
