package domain

import (
	"campsite/internal/dto"
	"context"
	"time"

	"gorm.io/gorm"
)

type Booking struct {
	ID             int64 `gorm:"primaryKey"`
	CampsiteID     int64
	UserID         int64
	CheckInDate    time.Time
	CheckOutDate   time.Time
	NumberOfPeople int64
	TotalCost      float64
	Status         int8
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt
}

type BookingRepository interface {
	FindByID(ctx context.Context, id int64) (Booking, error)
	GetAll(ctx context.Context) ([]Booking, error)
	Insert(ctx context.Context, booking *Booking) error
	Update(ctx context.Context, booking *Booking) error
	Delete(ctx context.Context, booking *Booking) error
}

type BookingService interface {
	AddBooking(ctx context.Context, request dto.BookingRequest) dto.Response
	GetAllBookings(ctx context.Context) dto.Response
	GetBooking(ctx context.Context, id int64) dto.Response
	UpdateBooking(ctx context.Context, request dto.BookingRequest) dto.Response
	DeleteBooking(ctx context.Context, request dto.BookingRequest) dto.Response
}
