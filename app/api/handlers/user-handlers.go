package handlers

import (
	"demo/app/services"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	usrService services.UserService
}

func NewUserHandler(usrService services.UserService) UserHandler {
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
