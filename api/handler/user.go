package handler

import (
	"demo/api/constant"
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

func (user UserHandler) GetAccount(ctx *fiber.Ctx) error {
	username := ctx.Params("username")

	account, err := user.service.GetAccount(username)

	if err != nil {
		return err
	}

	return ctx.JSON(account)
}

func (user UserHandler) CreateAccount(ctx *fiber.Ctx) error {
	details := ctx.Locals(constant.Resource).(models.AccountDetails)

	if err := user.service.CreateAccount(details); err != nil {
		return err
	}

	return ctx.SendStatus(fiber.StatusCreated)
}

func (user UserHandler) Login(ctx *fiber.Ctx) error {
	credentials := new(struct {
		Username string `json:"username"`
		Password string `json:"password"`
	})

	if err := ctx.BodyParser(credentials); err != nil {
		return err
	}

	success, err := user.service.AuthenticateUser(credentials.Username, credentials.Password)

	if err != nil {
		return err
	}

	if success {
		return ctx.SendStatus(fiber.StatusAccepted)
	}

	return ctx.SendStatus(fiber.StatusUnauthorized)
}
