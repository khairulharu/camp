package repository

import (
	"campsite/internal/domain"
	"context"

	"gorm.io/gorm"
)

type availabilityRepository struct {
	db *gorm.DB
}

func NewAvailabilityRepository(db *gorm.DB) domain.AvailabilityRepository {
	return &availabilityRepository{
		db: db,
	}
}

// FindByID implements domain.ReviewRepository.
func (a *availabilityRepository) FindByID(ctx context.Context, id int64) (domain.Availability, error) {
	var availability domain.Availability
	err := a.db.WithContext(ctx).Table("availabilities").Where("id = ?", &id).First(&availability).Error
	if err != nil {
		return domain.Availability{}, err
	}
	return availability, nil
}

// GetAll implements domain.AvailabilityRepository.
func (a *availabilityRepository) GetAll(ctx context.Context) ([]domain.Availability, error) {
	var availabilities []domain.Availability
	err := a.db.WithContext(ctx).Table("availabilities").Find(&availabilities).Error
	if err != nil {
		return nil, err
	}
	return availabilities, nil
}

// Insert implements domain.AvailabilityRepository.
func (a *availabilityRepository) Insert(ctx context.Context, availability *domain.Availability) error {
	return a.db.WithContext(ctx).Table("availabilities").Create(&availability).Error
}

// Update implements domain.AvailabilityRepository.
func (a *availabilityRepository) Update(ctx context.Context, availability *domain.Availability) error {
	return a.db.WithContext(ctx).Table("availabilities").Model(&domain.Availability{}).Where("id = ?", &availability.ID).Updates(&availability).Error
}
