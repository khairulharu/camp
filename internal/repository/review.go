package repository

import (
	"campsite/internal/domain"
	"context"

	"gorm.io/gorm"
)

type riviewRepository struct {
	db *gorm.DB
}

func NewReviewRepository(db *gorm.DB) domain.ReviewRepository {
	return &riviewRepository{
		db: db,
	}
}

// FindByID implements domain.ReviewRepository.
func (r *riviewRepository) FindByID(ctx context.Context, id int64) (domain.Review, error) {
	var review domain.Review
	err := r.db.WithContext(ctx).Table("reviews").Where("id = ?", id).First(review).Error
	if err != nil {
		return domain.Review{}, err
	}
	return review, nil
}

// GetAll implements domain.ReviewRepository.
func (r *riviewRepository) GetAll(ctx context.Context) ([]domain.Review, error) {
	var reviews []domain.Review
	err := r.db.WithContext(ctx).Table("reviews").Find(&reviews).Error
	if err != nil {
		return nil, err
	}
	return reviews, nil
}

// Insert implements domain.ReviewRepository.
func (r *riviewRepository) Insert(ctx context.Context, review *domain.Review) error {
	return r.db.WithContext(ctx).Table("reviews").Create(review).Error
}

// Update implements domain.ReviewRepository.
func (r *riviewRepository) Update(ctx context.Context, review *domain.Review) error {
	return r.db.WithContext(ctx).Table("reviews").Model(&domain.Review{}).Where("id = ?", review.ID).Updates(&review).Error
}

// Delete implements domain.ReviewRepository.
func (r *riviewRepository) Delete(ctx context.Context, review *domain.Review) error {
	return r.db.WithContext(ctx).Table("reviews").Delete(&review).Error
}
