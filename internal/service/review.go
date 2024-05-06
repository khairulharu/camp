package service

import (
	"campsite/internal/domain"
	"campsite/internal/dto"
	"context"
	"time"
)

type reviewService struct {
	reviewRepository domain.ReviewRepository
}

func NewReview(reviewRepository domain.ReviewRepository) domain.ReviewService {
	return &reviewService{
		reviewRepository: reviewRepository,
	}
}

// AddReview implements domain.ReviewService.
func (r *reviewService) AddReview(ctx context.Context, request dto.ReviewRequest) dto.Response {
	review := domain.Review{
		CampsiteID: request.CampsiteID,
		UserID:     request.UserID,
		Rating:     request.Rating,
		Comment:    request.Comment,
		ReviewDate: time.Now(),
		CreatedAt:  time.Now(),
	}

	if err := r.reviewRepository.Insert(ctx, &review); err != nil {
		return dto.Response{
			Status:  "401",
			Message: "fail deleted review",
			Error:   err.Error(),
		}
	}

	return dto.Response{
		Status:  "200",
		Message: "succes add review",
	}
}

// DeleteReview implements domain.ReviewService.
func (r *reviewService) DeleteReview(ctx context.Context, request dto.ReviewRequest) dto.Response {

	err := r.reviewRepository.Delete(ctx, &domain.Review{ID: request.ID})

	if err != nil {
		return dto.Response{
			Status:  "401",
			Message: "fail delete review",
			Error:   err.Error(),
		}
	}

	return dto.Response{
		Status:  "200",
		Message: "deleted",
	}
}

// GetAllReviews implements domain.ReviewService.
func (r *reviewService) GetAllReviews(ctx context.Context) dto.Response {
	reviews, err := r.reviewRepository.GetAll(ctx)

	if err != nil {
		return dto.Response{
			Status:  "401",
			Message: "cannot gett all Review",
			Error:   err.Error(),
		}
	}

	var reviewsResponse []dto.ReviewResponse

	for _, value := range reviews {
		reviewsResponse = append(reviewsResponse, dto.ReviewResponse{
			ID:         value.ID,
			CampsiteID: value.CampsiteID,
			UserID:     value.UserID,
			Rating:     value.Rating,
			Comment:    value.Comment,
			ReviewDate: value.ReviewDate,
		})
	}

	return dto.Response{
		Status:  "200",
		Message: "accpted",
		Data:    reviewsResponse,
	}
}

// GetReview implements domain.ReviewService.
func (r *reviewService) GetReview(ctx context.Context, id int64) dto.Response {
	review, err := r.reviewRepository.FindByID(ctx, id)

	if err != nil {
		return dto.Response{
			Status:  "401",
			Message: "fail get review",
			Error:   err.Error(),
		}
	}

	return dto.Response{
		Status:  "200",
		Message: "acceptd",
		Data: dto.ReviewResponse{
			ID:         review.ID,
			CampsiteID: review.CampsiteID,
			UserID:     review.UserID,
			Rating:     review.Rating,
			Comment:    review.Comment,
			ReviewDate: review.ReviewDate,
		},
	}
}

// UpdateReview implements domain.ReviewService.
func (r *reviewService) UpdateReview(ctx context.Context, request dto.ReviewRequest) dto.Response {
	review := domain.Review{
		ID:         request.ID,
		Rating:     request.Rating,
		Comment:    request.Comment,
		ReviewDate: time.Now(),
	}

	if err := r.reviewRepository.Update(ctx, &review); err != nil {
		return dto.Response{
			Status:  "401",
			Message: "failed update review",
		}
	}

	return dto.Response{
		Status:  "200",
		Message: "accepted",
	}
}
