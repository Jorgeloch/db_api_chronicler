package customerRouter

import (
	customerController "atividade_4/src/customer/controller"
	customerRepository "atividade_4/src/customer/repository"
	customerService "atividade_4/src/customer/service"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5"
)

func NewCustomerRouter(db *pgx.Conn) *fiber.App {
	router := fiber.New()

	repository := customerRepository.InitCustomerRepository(db)
	service := customerService.InitCustomerService(repository)
	controller := customerController.InitCustomerController(service)

	router.Get("/", controller.HandleFindAll)
	router.Get("/:cpf", controller.HandleFindByID)
	router.Post("/", controller.HandleCreateCustomer)
	router.Patch("/:cpf", controller.HandleUpdateCustomer)
	router.Delete("/:cpf", controller.HandleDeleteCustomer)

	return router
}
