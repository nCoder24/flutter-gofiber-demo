package handler

import (
	"demo/api/constant"
	"demo/core/models"
	"demo/core/operator"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	operator operator.UserOperator
}

func NewUserHandler(operator operator.UserOperator) UserHandler {
	return UserHandler{operator}
}

func (user UserHandler) GetAccount(ctx *fiber.Ctx) error {
	username := ctx.Params("username")

	account, err := user.operator.GetAccount(username)

	if err != nil {
		return err
	}

	return ctx.JSON(account)
}

func (user UserHandler) CreateAccount(ctx *fiber.Ctx) error {
	details := ctx.Locals(constant.Resource).(models.AccountDetails)

	if err := user.operator.CreateAccount(details); err != nil {
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

	success, err := user.operator.AuthenticateUser(credentials.Username, credentials.Password)

	if err != nil {
		return err
	}

	if success {
		return ctx.SendStatus(fiber.StatusAccepted)
	}

	return ctx.SendStatus(fiber.StatusUnauthorized)
}
