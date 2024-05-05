package domain

import (
	"campsite/internal/dto"
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

type BookingService interface {
	Camp(campID int64) dto.Response
}
