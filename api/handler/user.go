package handler

import (
	"demo/api/constant"
	apiErrs "demo/api/errors"
	internalErrs "demo/core/errors"
	"demo/core/models"
	"demo/core/operator"
	"log"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	operator operator.UserOperator
}

func NewUserHandler(operator operator.UserOperator) UserHandler {
	return UserHandler{operator}
}

func (usr UserHandler) GetUser(ctx *fiber.Ctx) error {
	username := ctx.Params("username")
	user, err := usr.operator.GetUser(username)

	if err == nil {
		return ctx.JSON(user)
	}

	log.Println(err)
	if err == internalErrs.ErrUserDoesNotExists {
		return apiErrs.UserNotFound
	}

	return apiErrs.CouldNotGetUser
}

func (usr UserHandler) AddNewUser(ctx *fiber.Ctx) error {
	usrData := ctx.Locals(constant.Resource).(models.UserData)

	if err := usr.operator.AddNewUser(usrData); err != nil {
		log.Println(err)
		return apiErrs.CouldNotAddUser
	}

	return ctx.SendStatus(fiber.StatusCreated)
}
