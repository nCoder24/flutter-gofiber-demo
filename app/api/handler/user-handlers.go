package handler

import (
	"demo/app/service"
	"demo/core/models"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	service service.UserService
}

func NewUserHandler(service service.UserService) UserHandler {
	return UserHandler{service}
}

func (handler UserHandler) GetAccount(c *fiber.Ctx) error {
	username := c.Params("username")
	account, err := handler.service.GetAccount(username)

	if err != nil {
		return err
	}

	return c.JSON(account)
}

func (handler UserHandler) CreateAccount(c *fiber.Ctx) error {
	acDetails := new(models.AccountDetails)

	if err := c.BodyParser(acDetails); err != nil {
		return err
	}

	if err := handler.service.CreateAccount(*acDetails); err != nil {
		return err
	}

	return c.SendStatus(201)
}
