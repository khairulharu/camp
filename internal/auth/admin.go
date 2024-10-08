package auth

import (
	"campsite/internal/domain"
	"campsite/internal/util"

	"github.com/gofiber/fiber/v2"
)

type adminAuth struct {
	adminService domain.AdminService
}

func NewAdmin(app *fiber.App, adminService domain.AdminService, authMid fiber.Handler) {
	handler := adminAuth{
		adminService: adminService,
	}

	app.Get("/allbookings", authMid, handler.GetAllBookings)
	app.Get("/allcampsites", authMid, handler.GetAllCampsites)
	app.Get("/allreviews", authMid, handler.GetAllReviews)
}

func (a *adminAuth) GetAllBookings(ctx *fiber.Ctx) error {
	response := a.adminService.GetAllBookings(ctx.Context())

	return ctx.Status(util.GetHttpStatus(response.Status)).JSON(response)
}

func (a *adminAuth) GetAllCampsites(ctx *fiber.Ctx) error {
	response := a.adminService.GetAllCampsites(ctx.Context())

	return ctx.Status(util.GetHttpStatus(response.Status)).JSON(response)
}

func (a *adminAuth) GetAllReviews(ctx *fiber.Ctx) error {
	response := a.adminService.GetAllReviews(ctx.Context())

	return ctx.Status(util.GetHttpStatus(response.Status)).JSON(response)
}
