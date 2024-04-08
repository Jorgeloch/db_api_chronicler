package tagRouter

import (
	tagController "atividade_4/src/tag/controller"
	tagRepository "atividade_4/src/tag/repository"
	tagService "atividade_4/src/tag/service"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5"
)

func NewTagRouter(db *pgx.Conn) *fiber.App {
	router := fiber.New()

	repository := tagRepository.InitTagRepository(db)
	service := tagService.InitTagService(repository)
	controller := tagController.InitTagController(service)

	router.Get("/", controller.HandleFindAll)
	router.Get("/:id", controller.HandleFindByID)
	router.Post("/", controller.HandleCreateTag)
	router.Patch("/:id", controller.HandleUpdateTag)
	router.Delete("/:id", controller.HandleDeleteTag)

	return router
}
