package handler

import (
	"demo/app/service"

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

func (handler UserHandler) AddUser(c *fiber.Ctx) error {
	return c.SendString("User Added")
}
