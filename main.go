package main

import (
	customerRouter "atividade_4/src/customer/router"
	TagCustomerRouter "atividade_4/src/tag-cliente/router"
	tagRouter "atividade_4/src/tag/router"
	"context"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	conn, err := InitConnection()
	if err != nil {
		panic(err)
	}

	app := fiber.New()

	app.Use(logger.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Mount("/tag", tagRouter.NewTagRouter(conn))
	app.Mount("/customer", customerRouter.NewCustomerRouter(conn))
	app.Mount("/tag_customer", TagCustomerRouter.NewManagerRouter(conn))

	app.Listen(":8080")
}

func InitConnection() (*pgx.Conn, error) {
	URL := os.Getenv("DATABASE_URL")
	db, err := pgx.Connect(context.Background(), URL)
	if err != nil {
		panic(err)
	}
	return db, err
}
