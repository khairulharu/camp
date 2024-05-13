package auth

import (
	"campsite/internal/domain"
	"campsite/internal/dto"
	"campsite/internal/util"
	"strconv"

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
	app.Delete("/booking/:id?", authMid, handler.DeleteBooking)
	app.Patch("/booking/:id?", authMid, handler.EditBooking)
}

func (b *bookingAuth) CreateBooking(ctx *fiber.Ctx) error {
	var bookingRequest dto.BookingRequest

	if err := ctx.BodyParser(&bookingRequest); err != nil {
		return ctx.Status(401).JSON("erorr in body parser cant parse get:", err.Error())
	}

	response := b.bokingService.AddBooking(ctx.Context(), bookingRequest)

	return ctx.Status(util.GetHttpStatus(response.Status)).JSON(response)
}

func (b *bookingAuth) GetAllBookings(ctx *fiber.Ctx) error {
	response := b.bokingService.GetAllBookings(ctx.Context())

	return ctx.Status(util.GetHttpStatus(response.Status)).JSON(response)
}

func (b *bookingAuth) DeleteBooking(ctx *fiber.Ctx) error {
	idS := ctx.Params("id")

	if idS == ("") {
		return ctx.Status(util.GetHttpStatus(dto.BADREQUEST)).JSON("message:id param null")
	}

	id, err := strconv.Atoi(idS)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	response := b.bokingService.DeleteBooking(ctx.Context(), dto.BookingRequest{ID: int64(id)})

	return ctx.Status(util.GetHttpStatus(response.Status)).JSON(response)
}

func (b *bookingAuth) EditBooking(ctx *fiber.Ctx) error {

	idS := ctx.Params("id")

	if idS == ("") {
		return ctx.Status(util.GetHttpStatus(dto.BADREQUEST)).JSON("message:id param null")
	}

	id, err := strconv.Atoi(idS)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	var bookingRequest dto.BookingRequest

	if err := ctx.BodyParser(&bookingRequest); err != nil {
		ctx.SendStatus(fiber.StatusBadRequest)
	}

	bookingRequest.CampsiteID = int64(id)

	response := b.bokingService.UpdateBooking(ctx.Context(), bookingRequest)

	return ctx.Status(util.GetHttpStatus(response.Status)).JSON(response)
}
