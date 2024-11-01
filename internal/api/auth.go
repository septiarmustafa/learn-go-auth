package api

import (
	"belajar-auth/domain"
	"belajar-auth/dto"
	"belajar-auth/internal/util"

	"github.com/gofiber/fiber/v2"
)

type AuthApi struct {
	userService domain.UserService
}

func NewUser(app *fiber.App, userService domain.UserService) {
	h := AuthApi{
		userService: userService,
	}
	app.Post("token/generate", h.GenerateToken)
}

func (a AuthApi) GenerateToken(ctx *fiber.Ctx) error {
	var req dto.AuthReq
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.SendStatus(400)
	}

	token, err := a.userService.Authenticate(ctx.Context(), req)
	if err != nil {
		return ctx.SendStatus(util.GetHttpStatus(err))
	}
	return ctx.Status(200).JSON(token)
}
