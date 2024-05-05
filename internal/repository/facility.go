package repository

import (
	"campsite/internal/domain"
	"context"

	"gorm.io/gorm"
)

type facilityRepository struct {
	db *gorm.DB
}

func NewFacilityRepository(db *gorm.DB) domain.FacilityRepository {
	return &facilityRepository{
		db: db,
	}
}

// FindByID implements domain.ReviewRepository.
func (f *facilityRepository) FindByID(ctx context.Context, id int64) (domain.Facility, error) {
	var facility domain.Facility
	err := f.db.WithContext(ctx).Table("facilities").Where("id = ?", id).First(facility).Error
	if err != nil {
		return domain.Facility{}, err
	}
	return facility, nil
}

// GetAll implements domain.facilityRepository.
func (f *facilityRepository) GetAll(ctx context.Context) ([]domain.Facility, error) {
	var facilitys []domain.Facility
	err := f.db.WithContext(ctx).Table("facilities").Find(&facilitys).Error
	if err != nil {
		return nil, err
	}
	return facilitys, nil
}

// Insert implements domain.facilityRepository.
func (f *facilityRepository) Insert(ctx context.Context, facility *domain.Facility) error {
	return f.db.WithContext(ctx).Table("facilities").Create(facility).Error
}

// Update implements domain.facilityRepository.
func (f *facilityRepository) Update(ctx context.Context, facility *domain.Facility) error {
	return f.db.WithContext(ctx).Table("facilities").Model(&domain.Facility{}).Where("id = ?", facility.ID).Updates(&facility).Error
}
