package customerController

import (
	customerDTO "atividade_4/src/customer/dto"
	customerService "atividade_4/src/customer/service"
	"log"

	"github.com/gofiber/fiber/v2"
)

type CustomerController struct {
	service *customerService.CustomerService
}

func InitCustomerController(s *customerService.CustomerService) *CustomerController {
	return &CustomerController{
		service: s,
	}
}

func (controller *CustomerController) HandleFindAll(c *fiber.Ctx) error {
	customers, err := controller.service.FindAll()
	if err != nil {
		log.Println(err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	return c.Status(fiber.StatusOK).JSON(customers)
}

func (controller *CustomerController) HandleFindByID(c *fiber.Ctx) error {
	cpf := c.Params("cpf")

	customer, err := controller.service.FindByID(cpf)
	if err != nil {
		log.Println(err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusOK).JSON(customer)
}

func (controller *CustomerController) HandleCreateCustomer(c *fiber.Ctx) error {
	var customerDTO customerDTO.CreateCustomerDTO

	err := c.BodyParser(&customerDTO)
	if err != nil {
		log.Println(err)
		return c.SendStatus(fiber.StatusBadRequest)
	}

	newCustomerCPF, err := controller.service.Create(customerDTO)
	if err != nil {
		log.Println(err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"location": newCustomerCPF,
	})
}

func (controller *CustomerController) HandleUpdateCustomer(c *fiber.Ctx) error {
	var customerDTO customerDTO.UpdateCustomerDTO

	err := c.BodyParser(&customerDTO)
	if err != nil {
		log.Println(err)
		return c.SendStatus(fiber.StatusBadRequest)
	}

	cpf := c.Params("cpf")

	customerUpdated, err := controller.service.Update(cpf, customerDTO)
	if err != nil {
		log.Println(err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusOK).JSON(customerUpdated)
}

func (controller *CustomerController) HandleDeleteCustomer(c *fiber.Ctx) error {
	cpf := c.Params("cpf")

	err := controller.service.Delete(cpf)
	if err != nil {
		log.Println(err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	return c.SendStatus(fiber.StatusOK)
}
