package auth

import (
	"campsite/internal/domain"
	"campsite/internal/dto"
	"campsite/internal/util"

	"github.com/gofiber/fiber/v2"
)

type reviewAuth struct {
	reviewService domain.ReviewService
}

func NewReview(app *fiber.App, reviewService domain.ReviewService, authMid fiber.Handler) {
	handler := reviewAuth{
		reviewService: reviewService,
	}

	app.Post("/review", authMid, handler.UserAddReview)

}

func (r *reviewAuth) UserAddReview(ctx *fiber.Ctx) error {
	var reviewReqBody dto.ReviewRequest

	if err := ctx.BodyParser(&reviewReqBody); err != nil {
		return ctx.Status(401).JSON("error parsing body to request")
	}

	userLogin := ctx.Locals("x-user").(*util.MyCustomClaims)

	reviewReqBody.UserID = userLogin.ID

	response := r.reviewService.AddReview(ctx.Context(), reviewReqBody)

	return ctx.Status(util.GetHttpStatus(response.Status)).JSON(response)
}
