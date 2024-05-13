package auth

import (
	"campsite/internal/domain"
	"campsite/internal/dto"
	"campsite/internal/util"
	"strconv"

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
	app.Get("/review/:id?", authMid, handler.GetReview)
	app.Get("/reviews", authMid, handler.GetAllReviews)
	app.Delete("/review/:id?", authMid, handler.DeleteReview)

}

func (r *reviewAuth) UserAddReview(ctx *fiber.Ctx) error {
	var reviewReqBody dto.ReviewRequest

	if err := ctx.BodyParser(&reviewReqBody); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON("error parsing body to request get: ", err.Error())
	}

	userLogin := ctx.Locals("x-user").(*util.MyCustomClaims)

	reviewReqBody.UserID = userLogin.ID

	response := r.reviewService.AddReview(ctx.Context(), reviewReqBody)

	return ctx.Status(util.GetHttpStatus(response.Status)).JSON(response)
}

func (r *reviewAuth) GetReview(ctx *fiber.Ctx) error {
	idS := ctx.Params("id")

	if idS == ("") {
		return ctx.Status(util.GetHttpStatus(dto.BADREQUEST)).JSON("message:id param null")
	}

	id, err := strconv.Atoi(idS)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	response := r.reviewService.GetReview(ctx.Context(), int64(id))

	return ctx.Status(util.GetHttpStatus(response.Status)).JSON(response)
}

func (r *reviewAuth) GetAllReviews(ctx *fiber.Ctx) error {
	response := r.reviewService.GetAllReviews(ctx.Context())

	return ctx.Status(util.GetHttpStatus(response.Status)).JSON(response)
}

func (r *reviewAuth) DeleteReview(ctx *fiber.Ctx) error {
	idS := ctx.Params("id")

	if idS == ("") {
		return ctx.Status(fiber.StatusBadRequest).JSON("require param id for delete which review")
	}

	id, err := strconv.Atoi(idS)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON("error when parsing id to id int get: ", err.Error())
	}

	response := r.reviewService.DeleteReview(ctx.Context(), dto.ReviewRequest{
		ID: int64(id),
	})

	return ctx.Status(util.GetHttpStatus(response.Status)).JSON(response)
}
