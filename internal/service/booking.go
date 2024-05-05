package service

import (
	"campsite/internal/domain"
	"campsite/internal/dto"
	"context"
)

type bookingService struct {
	bookingRepository *domain.BookingRepository
}

func NewBooking(bookingRepository *domain.BookingRepository) domain.BookingService {
	return &bookingService{
		bookingRepository: bookingRepository,
	}
}

// AddBooking implements domain.BookingService.
func (b *bookingService) AddBooking(ctx context.Context, request dto.BookingRequest) dto.Response {
	panic("unimplemented")
}

// DeleteBooking implements domain.BookingService.
func (b *bookingService) DeleteBooking(ctx context.Context, request dto.BookingRequest) dto.Response {
	panic("unimplemented")
}

// GetAllBookings implements domain.BookingService.
func (b *bookingService) GetAllBookings(ctx context.Context) dto.Response {
	panic("unimplemented")
}

// GetBooking implements domain.BookingService.
func (b *bookingService) GetBooking(ctx context.Context) dto.Response {
	panic("unimplemented")
}

// UpdateBooking implements domain.BookingService.
func (b *bookingService) UpdateBooking(ctx context.Context, request dto.BookingRequest) dto.Response {
	panic("unimplemented")
}
