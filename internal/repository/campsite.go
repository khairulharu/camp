package repository

import (
	"campsite/internal/domain"
	"context"

	"gorm.io/gorm"
)

type campsiteRepository struct {
	db *gorm.DB
}

func NewCampsiteRepository(db *gorm.DB) domain.CampsiteRepository {
	return &campsiteRepository{
		db: db,
	}
}

// FindByID implements domain.ReviewRepository.
func (c *campsiteRepository) FindByID(ctx context.Context, id int64) (domain.Campsite, error) {
	var campsite domain.Campsite
	err := c.db.WithContext(ctx).Table("campsites").Where("id = ?", id).First(campsite).Error
	if err != nil {
		return domain.Campsite{}, err
	}
	return campsite, nil
}

// GetAll implements domain.campsiteRepository.
func (c *campsiteRepository) GetAll(ctx context.Context) ([]domain.Campsite, error) {
	var campsites []domain.Campsite
	err := c.db.WithContext(ctx).Table("campsites").Find(&campsites).Error
	if err != nil {
		return nil, err
	}
	return campsites, nil
}

// Insert implements domain.campsiteRepository.
func (c *campsiteRepository) Insert(ctx context.Context, campsite *domain.Campsite) error {
	return c.db.WithContext(ctx).Table("campsites").Create(campsite).Error
}

// Update implements domain.campsiteRepository.
func (c *campsiteRepository) Update(ctx context.Context, campsite *domain.Campsite) error {
	return c.db.WithContext(ctx).Table("campsites").Model(&domain.Campsite{}).Where("id = ?", campsite.ID).Updates(&campsite).Error
}

// Delete implements domain.CampsiteRepository.
func (c *campsiteRepository) Delete(ctx context.Context, campsite *domain.Campsite) error {
	return c.db.WithContext(ctx).Table("campsites").Delete(&campsite).Error
}
