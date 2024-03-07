package TagClientRouter

import (
	TagClientController "atividade_4/src/tag-cliente/controller"
	TagClientRepository "atividade_4/src/tag-cliente/repository"
	TagClientService "atividade_4/src/tag-cliente/service"

	"github.com/gofiber/fiber/v2"
)

func NewManagerRouter() *fiber.App {
	router := fiber.New()

	repository := TagClientRepository.InitTagRepository()
	service := TagClientService.InitTagClientService(repository)
	controller := TagClientController.InitTagClientController(service)

	router.Get("/", controller.HandleFindAll)
	router.Get("/:cliente_cpf", controller.HandleFindByClient)
	router.Post("/", controller.HandleCreateTagClient)
	router.Delete("/cliente/:cliente_cpf/tag/:tag_id", controller.HandleDeleteTagClient)

	return router
}
