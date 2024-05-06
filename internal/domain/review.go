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
	ReviewDate time.Time      `gorm:"type:datetime(3)"`
	CreatedAt  time.Time      `gorm:"type:datetime(3)"`
	UpdatedAt  time.Time      `gorm:"type:datetime(3)"`
	DeletedAt  gorm.DeletedAt `gorm:"type:datetime(3)"`
}

type ReviewRepository interface {
	FindByID(ctx context.Context, id int64) (Review, error)
	GetAll(ctx context.Context) ([]Review, error)
	Insert(ctx context.Context, review *Review) error
	Update(ctx context.Context, review *Review) error
	Delete(ctx context.Context, review *Review) error
}

type ReviewService interface {
	AddReview(ctx context.Context, request dto.ReviewRequest) dto.Response
	GetAllReviews(ctx context.Context) dto.Response
	GetReview(ctx context.Context, id int64) dto.Response
	UpdateReview(ctx context.Context, request dto.ReviewRequest) dto.Response
	DeleteReview(ctx context.Context, request dto.ReviewRequest) dto.Response
}
