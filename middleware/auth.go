package middleware

import (
	"campsite/util"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func Authenticate() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		token := strings.ReplaceAll(ctx.Get("Authorization"), "Bearer ", "")
		if token == "" {
			return ctx.SendStatus(401)
		}

		user, err := util.ValidateToken(token)
		if err != nil {
			return ctx.Status(400).JSON(err.Error())
		}
		ctx.Locals("x-user", user)
		return ctx.Next()
	}
}
