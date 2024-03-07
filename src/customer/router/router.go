package customerRouter

import (
	customerController "atividade_4/src/customer/controller"
	customerRepository "atividade_4/src/customer/repository"
	customerService "atividade_4/src/customer/service"
	"github.com/gofiber/fiber/v2"
)

func NewCustomerRouter() *fiber.App {
	router := fiber.New()

	repository := customerRepository.InitCustomerRepository()
	service := customerService.InitCustomerService(repository)
	controller := customerController.InitCustomerController(service)

	router.Get("/", controller.HandleFindAll)
	router.Get(":id", controller.HandleFindByID)
	router.Post("/", controller.HandleCreateCustomer)
	router.Patch("/", controller.HandleUpdateCustomer)
	router.Delete("/", controller.HandleDeleteCustomer)

	return router
}
