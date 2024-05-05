package repository

import (
	"campsite/internal/domain"
	"context"

	"gorm.io/gorm"
)

type activityRepository struct {
	db *gorm.DB
}

func NewActivityRepository(db *gorm.DB) domain.ActivityRepository {
	return &activityRepository{
		db: db,
	}
}

// FindByID implements domain.ReviewRepository.
func (a *activityRepository) FindByID(ctx context.Context, id int64) (domain.Activity, error) {
	var activity domain.Activity
	err := a.db.WithContext(ctx).Table("activities").Where("id = ?", id).First(activity).Error
	if err != nil {
		return domain.Activity{}, err
	}
	return activity, nil
}

// GetAll implements domain.activityRepository.
func (a *activityRepository) GetAll(ctx context.Context) ([]domain.Activity, error) {
	var activities []domain.Activity
	err := a.db.WithContext(ctx).Table("activities").Find(&activities).Error
	if err != nil {
		return nil, err
	}
	return activities, nil
}

// Insert implements domain.activityRepository.
func (a *activityRepository) Insert(ctx context.Context, activity *domain.Activity) error {
	return a.db.WithContext(ctx).Table("activities").Create(activity).Error
}

// Update implements domain.activityRepository.
func (a *activityRepository) Update(ctx context.Context, activity *domain.Activity) error {
	return a.db.WithContext(ctx).Table("activities").Model(&domain.Activity{}).Where("id = ?", activity.ID).Updates(&activity).Error
}
