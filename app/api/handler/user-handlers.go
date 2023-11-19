package handler

import (
	"demo/app/service"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	usrService service.UserService
}

func NewUserHandler(usrService service.UserService) UserHandler {
	return UserHandler{usrService}
}

func (h UserHandler) GetUsers(c *fiber.Ctx) error {
	users, err := h.usrService.GetAllAccounts()
	if err != nil {
		return err
	}

	return c.JSON(users)
}

func (h UserHandler) AddUser(c *fiber.Ctx) error {
	return c.SendString("User Added")
}
