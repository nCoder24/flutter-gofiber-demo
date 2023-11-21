package handler

import (
	"demo/core/models"
	"demo/core/service"

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
		return err // Check if there's a better way of handling server err
	}

	if err := handler.service.CreateAccount(*acDetails); err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusCreated)
}

func (handler UserHandler) Login(c *fiber.Ctx) error {
	credentials := new(struct {
		Username string `json:"username"`
		Password string `json:"password"`
	})

	if err := c.BodyParser(credentials); err != nil {
		return err
	}

	success, err := handler.service.AuthenticateUser(credentials.Username, credentials.Password)

	if err != nil {
		return err
	}

	if success {
		return c.SendStatus(fiber.StatusAccepted)
	}

	return c.SendStatus(fiber.StatusUnauthorized)
}
