package domain

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Facility struct {
	ID          int64 `gorm:"primaryKey"`
	CampsiteID  int64
	Name        string
	Description string `gorm:"type:text"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt
}

type FacilityRepository interface {
	FindByID(ctx context.Context, id int64) (Facility, error)
	GetAll(ctx context.Context) ([]Facility, error)
	Insert(ctx context.Context, facility *Facility) error
	Update(ctx context.Context, facility *Facility) error
}

// type FacilityService interface {
// 	AddReview(request dto.FacilityRequest) dto.Response
// 	GetAllReviews() dto.Response
// 	GetReview() dto.Response
// 	UpdateReview(request dto.FacilityRequest) dto.Response
// 	DeleteReview(request dto.FacilityRequest) dto.Response
// }
