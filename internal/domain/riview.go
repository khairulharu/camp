package domain

import (
	"campsite/internal/dto"
	"context"
	"time"

	"gorm.io/gorm"
)

type Review struct {
	ID         int64
	CampsiteID int64
	UserID     int64
	Rating     int8
	Comment    string
	RiviewDate time.Time      `gorm:"type:datetime(3)"`
	CreatedAt  time.Time      `gorm:"type:datetime(3)"`
	UpdatedAt  time.Time      `gorm:"type:datetime(3)"`
	DeletedAt  gorm.DeletedAt `gorm:"type:datetime(3)"`
}

type ReviewRepository interface {
	FindByID(ctx context.Context, id int64) (Review, error)
	GetAll(ctx context.Context) ([]Review, error)
	Insert(ctx context.Context, user *Review) error
	Update(ctx context.Context, user *Review) error
}

type ReviewService interface {
	AddReview(request dto.ReviewRequest) dto.Response
	GetAllReviews() dto.Response
	GetReview() dto.Response
	UpdateReview(request dto.ReviewRequest) dto.Response
	DeleteReview(request dto.ReviewRequest) dto.Response
}
