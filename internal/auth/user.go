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

	user := u.userService.AuthUser(ctx.Context(), userReq)

	if user == (dto.UserResponse{}) {
		return ctx.Status(401).JSON("user_notfound: user not found, signup")
	}

	token, err := util.CreateToken(&domain.User{
		ID:   user.ID,
		Name: user.Name,
	})

	if err != nil {
		return ctx.SendStatus(401)
	}

	return ctx.Status(200).JSON(token)
}

func (u *userAuth) SignUp(ctx *fiber.Ctx) error {
	var userReq dto.UserRequest
	if err := ctx.BodyParser(&userReq); err != nil {
		return ctx.Status(401).JSON("Error: error parsed body to json")
	}

	response := u.userService.SignUp(ctx.Context(), userReq)

	return ctx.Status(util.GetHttpStatus(response.Status)).JSON(response)
}
