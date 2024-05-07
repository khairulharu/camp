package service

import (
	"campsite/internal/domain"
	"campsite/internal/dto"
	"context"
	"time"
)

type bookingService struct {
	bookingRepository domain.BookingRepository
}

func NewBooking(bookingRepository domain.BookingRepository) domain.BookingService {
	return &bookingService{
		bookingRepository: bookingRepository,
	}
}

// AddBooking implements domain.BookingService.
func (b *bookingService) AddBooking(ctx context.Context, request dto.BookingRequest) dto.Response {
	booking := domain.Booking{
		CampsiteID:     request.CampsiteID,
		UserID:         request.UserID,
		CheckInDate:    time.Now(),
		NumberOfPeople: request.NumberOfPeople,
		TotalCost:      request.TotalCost,
		Status:         request.Status,
	}

	if err := b.bookingRepository.Insert(ctx, &booking); err != nil {
		return dto.Response{
			Status:  "400",
			Message: "Cannot add booking",
		}
	}

	return dto.Response{
		Status:  "200",
		Message: "booking add",
	}
}

// DeleteBooking implements domain.BookingService.
func (b *bookingService) DeleteBooking(ctx context.Context, request dto.BookingRequest) dto.Response {
	err := b.bookingRepository.Delete(ctx, &domain.Booking{
		ID: request.ID,
	})

	if err != nil {
		return dto.Response{
			Status:  "400",
			Message: "cannot delete",
			Error:   err.Error(),
		}
	}

	return dto.Response{
		Status:  "200",
		Message: "booking added",
	}
}

// GetAllBookings implements domain.BookingService.
func (b *bookingService) GetAllBookings(ctx context.Context) dto.Response {
	bookings, err := b.bookingRepository.GetAll(ctx)

	if err != nil {
		return dto.Response{
			Status:  "400",
			Message: "Failed get bookings",
		}
	}

	var bookingsResponse []dto.BookingResponse

	for _, value := range bookings {
		bookingsResponse = append(bookingsResponse, dto.BookingResponse{
			ID:             value.ID,
			CampsiteID:     value.CampsiteID,
			UserID:         value.UserID,
			CheckInDate:    value.CheckInDate,
			CheckOutDate:   value.CheckOutDate,
			NumberOfPeople: value.NumberOfPeople,
			TotalCost:      value.TotalCost,
			Status:         value.Status,
		})
	}

	return dto.Response{
		Status:  "200",
		Message: "accpeted",
		Data:    bookingsResponse,
	}
}

// GetBooking implements domain.BookingService.
func (b *bookingService) GetBooking(ctx context.Context, id int64) dto.Response {
	booking, err := b.bookingRepository.FindByID(ctx, id)

	if err != nil {
		return dto.Response{
			Status:  "400",
			Message: "cannot get booking",
			Error:   err.Error(),
		}
	}

	return dto.Response{
		Status:  "200",
		Message: "accepted",
		Data: dto.BookingResponse{
			ID:             booking.ID,
			CampsiteID:     booking.CampsiteID,
			UserID:         booking.UserID,
			CheckInDate:    booking.CheckInDate,
			CheckOutDate:   booking.CheckOutDate,
			NumberOfPeople: booking.NumberOfPeople,
			TotalCost:      booking.TotalCost,
			Status:         booking.Status,
		},
	}

}

// UpdateBooking implements domain.BookingService.
func (b *bookingService) UpdateBooking(ctx context.Context, request dto.BookingRequest) dto.Response {
	booking := domain.Booking{
		ID:             request.ID,
		CampsiteID:     request.CampsiteID,
		UserID:         request.UserID,
		CheckInDate:    request.CheckInDate,
		CheckOutDate:   request.CheckInDate,
		NumberOfPeople: request.NumberOfPeople,
		TotalCost:      request.TotalCost,
		Status:         request.Status,
	}

	if err := b.bookingRepository.Update(ctx, &booking); err != nil {
		return dto.Response{
			Status:  "400",
			Message: "error when update booking",
		}
	}

	return dto.Response{
		Status:  "200",
		Message: "booking update",
	}
}
