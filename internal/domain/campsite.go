package domain

import (
	"campsite/internal/dto"
	"context"
	"time"

	"gorm.io/gorm"
)

type Campsite struct {
	ID            int64          `gorm:"primaryKey"`
	Name          string         `db:"name"`
	Location      string         `db:"location"`
	Latitude      float64        `db:"latitude"`
	Longitude     float64        `db:"longitude"`
	Area          float64        `db:"area"`
	PricePerNight float64        `db:"price_per_night"`
	CreatedAt     time.Time      `gorm:"type:datetime(3)"`
	UpdatedAt     time.Time      `gorm:"type:datetime(3)"`
	DeletedAt     gorm.DeletedAt `gorm:"type:datetime(3)"`
}

type CampsiteRepository interface {
	FindByID(ctx context.Context, id int64) (Campsite, error)
	GetAll(ctx context.Context) ([]Campsite, error)
	Insert(ctx context.Context, user *Campsite) error
	Update(ctx context.Context, user *Campsite) error
}

type CampsiteService interface {
	AddReview(request dto.CampsiteRequest) dto.Response
	GetAllReviews() dto.Response
	GetReview() dto.Response
	UpdateReview(request dto.CampsiteRequest) dto.Response
	DeleteReview(request dto.CampsiteRequest) dto.Response
}
