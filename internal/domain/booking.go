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
	CheckInDate    time.Time `gorm:"type:datetime(3)"`
	CheckOutDate   time.Time `gorm:"type:datetime(3)"`
	NumberOfPeople int64
	TotalCost      float64
	Status         int8
	CreatedAt      time.Time      `gorm:"type:datetime(3)"`
	UpdatedAt      time.Time      `gorm:"type:datetime(3)"`
	DeletedAt      gorm.DeletedAt `gorm:"type:datetime(3)"`
}

type BookingRepository interface {
	FindByID(ctx context.Context, id int64) (Campsite, error)
	GetAll(ctx context.Context) ([]Campsite, error)
	Insert(ctx context.Context, user *Campsite) error
	Update(ctx context.Context, user *Campsite) error
}

type BookingService interface {
	AddReview(request dto.BookingRequest) dto.Response
	GetAllReviews() dto.Response
	GetReview() dto.Response
	UpdateReview(request dto.BookingRequest) dto.Response
	DeleteReview(request dto.BookingRequest) dto.Response
}
