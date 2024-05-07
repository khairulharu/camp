package domain

import (
	"campsite/internal/dto"
	"context"
	"time"

	"gorm.io/gorm"
)

type Campsite struct {
	ID            int64   `gorm:"primaryKey"`
	Name          string  `db:"name"`
	Location      string  `db:"location"`
	Latitude      float64 `db:"latitude"`
	Longitude     float64 `db:"longitude"`
	Area          float64 `db:"area"`
	PricePerNight float64 `db:"price_per_night"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt
}

type CampsiteRepository interface {
	FindByID(ctx context.Context, id int64) (Campsite, error)
	GetAll(ctx context.Context) ([]Campsite, error)
	Insert(ctx context.Context, campsite *Campsite) error
	Update(ctx context.Context, campsite *Campsite) error
	Delete(ctx context.Context, campsite *Campsite) error
}

type CampsiteService interface {
	AddCampsite(ctx context.Context, request dto.CampsiteRequest) dto.Response
	GetAllCampsites() dto.Response
	GetCampsite(ctx context.Context, id int64) dto.Response
	UpdateCampsite(ctx context.Context, request dto.CampsiteRequest) dto.Response
	DeleteCampsite(ctx context.Context, request dto.CampsiteRequest) dto.Response
}
