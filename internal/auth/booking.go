package auth

import (
	"campsite/internal/domain"
	"campsite/internal/dto"
	"campsite/util"

	"github.com/gofiber/fiber/v2"
)

type bookingAuth struct {
	bokingService domain.BookingService
}

func NewBooking(app *fiber.App, bookingService domain.BookingService, authMid fiber.Handler) {
	handler := bookingAuth{
		bokingService: bookingService,
	}

	app.Post("/booking", authMid, handler.CreateBooking)
	app.Get("/bookings", handler.GetAllBookings)
}

func (b *bookingAuth) CreateBooking(ctx *fiber.Ctx) error {
	var bookingRequest dto.BookingRequest

	if err := ctx.BodyParser(&bookingRequest); err != nil {
		return ctx.Status(401).JSON("erorr in body parser cant parse ")
	}

	response := b.bokingService.AddBooking(ctx.Context(), bookingRequest)

	return ctx.Status(util.GetHttpStatus(response.Status)).JSON(response)
}

func (b *bookingAuth) GetAllBookings(ctx *fiber.Ctx) error {
	response := b.bokingService.GetAllBookings(ctx.Context())

	return ctx.Status(util.GetHttpStatus(response.Status)).JSON(response)
}
