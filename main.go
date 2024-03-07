package main

import (
	customerRouter "atividade_4/src/customer/router"
	TagCustomerRouter "atividade_4/src/tag-cliente/router"
	tagRouter "atividade_4/src/tag/router"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	app := fiber.New()

	app.Use(logger.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Mount("/tag", tagRouter.NewTagRouter())
	app.Mount("/customer", customerRouter.NewCustomerRouter())
	app.Mount("/tag_customer", TagCustomerRouter.NewManagerRouter())

	app.Listen(":8080")
}
