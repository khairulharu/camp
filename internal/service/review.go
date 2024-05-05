package service

import (
	"campsite/internal/domain"
	"campsite/internal/dto"
	"context"
)

type reviewService struct {
	reviewRepository *domain.ReviewRepository
}

func NewReview(reviewRepository *domain.ReviewRepository) domain.ReviewService {
	return &reviewService{
		reviewRepository: reviewRepository,
	}
}

// AddReview implements domain.ReviewService.
func (r *reviewService) AddReview(ctx context.Context, request dto.ReviewRequest) dto.Response {
	panic("unimplemented")
}

// DeleteReview implements domain.ReviewService.
func (r *reviewService) DeleteReview(ctx context.Context, request dto.ReviewRequest) dto.Response {
	panic("unimplemented")
}

// GetAllReviews implements domain.ReviewService.
func (r *reviewService) GetAllReviews(ctx context.Context) dto.Response {
	panic("unimplemented")
}

// GetReview implements domain.ReviewService.
func (r *reviewService) GetReview(ctx context.Context, id int64) dto.Response {
	panic("unimplemented")
}

// UpdateReview implements domain.ReviewService.
func (r *reviewService) UpdateReview(ctx context.Context, request dto.ReviewRequest) dto.Response {
	panic("unimplemented")
}
