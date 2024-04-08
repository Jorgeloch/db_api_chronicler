package TagCustomerRouter

import (
	customerRepository "atividade_4/src/customer/repository"
	customerService "atividade_4/src/customer/service"
	TagCustomerController "atividade_4/src/tag-cliente/controller"
	TagCustomerRepository "atividade_4/src/tag-cliente/repository"
	TagCustomerService "atividade_4/src/tag-cliente/service"
	tagRepository "atividade_4/src/tag/repository"
	tagService "atividade_4/src/tag/service"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5"
)

func NewManagerRouter(db *pgx.Conn) *fiber.App {
	router := fiber.New()

	customerRepository := customerRepository.InitCustomerRepository(db)
	tagRepository := tagRepository.InitTagRepository(db)
	customerService := customerService.InitCustomerService(customerRepository)
	tagService := tagService.InitTagService(tagRepository)

	repository := TagCustomerRepository.InitTagRepository(db, customerService, tagService)
	service := TagCustomerService.InitTagCustomerService(repository)
	controller := TagCustomerController.InitTagCustomerController(service)

	router.Get("/", controller.HandleFindAll)
	router.Post("/", controller.HandleCreateTagCustomer)
	router.Get("/customer/:cliente_cpf", controller.HandleFindByCustomer)
	router.Get("/tag/:tag_id", controller.HandleFindByTag)
	router.Delete("/customer/:cliente_cpf/tag/:tag_id", controller.HandleDeleteTagCustomer)

	return router
}
