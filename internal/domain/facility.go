package domain

import (
	"campsite/internal/dto"
	"context"
	"time"

	"gorm.io/gorm"
)

type Facility struct {
	ID          int64 `gorm:"primaryKey"`
	CampsiteID  int64
	Name        string
	Description string         `gorm:"type:text"`
	CreatedAt   time.Time      `gorm:"type:datetime(3)"`
	UpdatedAt   time.Time      `gorm:"type:datetime(3)"`
	DeletedAt   gorm.DeletedAt `gorm:"type:datetime(3)"`
}

type FacilityRepository interface {
	FindByID(ctx context.Context, id int64) (Facility, error)
	GetAll(ctx context.Context) ([]Facility, error)
	Insert(ctx context.Context, user *Facility) error
	Update(ctx context.Context, user *Facility) error
}

type FacilityService interface {
	AddReview(request dto.FacilityRequest) dto.Response
	GetAllReviews() dto.Response
	GetReview() dto.Response
	UpdateReview(request dto.FacilityRequest) dto.Response
	DeleteReview(request dto.FacilityRequest) dto.Response
}
