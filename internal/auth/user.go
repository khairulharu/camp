package auth

import (
	"campsite/internal/domain"
	"campsite/internal/dto"
	"campsite/internal/util"

	"github.com/gofiber/fiber/v2"
)

type userAuth struct {
	userService domain.UserService
}

func NewUser(app *fiber.App, userService domain.UserService, authMid fiber.Handler) {
	handler := userAuth{
		userService: userService,
	}

	app.Post("/login", handler.UserLogin)
	app.Post("/signup", handler.SignUp)

}

func (u *userAuth) UserLogin(ctx *fiber.Ctx) error {
	var userReq dto.UserRequest
	if err := ctx.BodyParser(&userReq); err != nil {
		return ctx.Status(401).JSON("Error: error parsed body to json")
	}

	response := u.userService.AuthUser(ctx.Context(), userReq)

	if response.Status != dto.OK {
		return ctx.Status(util.GetHttpStatus(response.Status)).JSON(response)
	}

	token, err := util.CreateToken(&domain.User{
		ID:   userReq.ID,
		Name: userReq.Name,
	})

	if err != nil {
		return ctx.SendStatus(util.GetHttpStatus(dto.INTERNALSERVERERROR))
	}

	return ctx.Status(util.GetHttpStatus(response.Status)).JSON(token)
}

func (u *userAuth) SignUp(ctx *fiber.Ctx) error {
	var userReq dto.UserRequest
	if err := ctx.BodyParser(&userReq); err != nil {
		return ctx.Status(401).JSON("Error: error parsed body to json")
	}

	response := u.userService.SignUp(ctx.Context(), userReq)

	return ctx.Status(util.GetHttpStatus(response.Status)).JSON(response)
}
