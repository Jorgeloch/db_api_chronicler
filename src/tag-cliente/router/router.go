package TagCustomerRouter

import (
	TagCustomerController "atividade_4/src/tag-cliente/controller"
	TagCustomerRepository "atividade_4/src/tag-cliente/repository"
	TagCustomerService "atividade_4/src/tag-cliente/service"

	"github.com/gofiber/fiber/v2"
)

func NewManagerRouter() *fiber.App {
	router := fiber.New()

	repository := TagCustomerRepository.InitTagRepository()
	service := TagCustomerService.InitTagCustomerService(repository)
	controller := TagCustomerController.InitTagCustomerController(service)

	router.Get("/", controller.HandleFindAll)
	router.Post("/", controller.HandleCreateTagCustomer)
	router.Get("/:cliente_cpf", controller.HandleFindByCustomer)
	router.Delete("/cliente/:cliente_cpf/tag/:tag_id", controller.HandleDeleteTagCustomer)

	return router
}
