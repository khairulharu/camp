package repository

import (
	"campsite/internal/domain"
	"context"

	"gorm.io/gorm"
)

type adminRepository struct {
	db *gorm.DB
}

func NewAdminRepository(db *gorm.DB) domain.AdminRepository {
	return &adminRepository{
		db: db,
	}
}

// FindByID implements domain.AdminRepository.
func (a *adminRepository) FindByID(ctx context.Context, id int64) (domain.Admin, error) {
	var admin domain.Admin
	err := a.db.WithContext(ctx).Table("admins").Where("id = ?", id).First(admin).Error
	if err != nil {
		return domain.Admin{}, err
	}
	return admin, nil
}

// GetAll implements domain.AdminRepository.
func (a *adminRepository) GetAll(ctx context.Context) ([]domain.Admin, error) {
	var admins []domain.Admin
	err := a.db.WithContext(ctx).Table("admins").Find(&admins).Error
	if err != nil {
		return nil, err
	}
	return admins, nil
}

// Insert implements domain.AdminRepository.
func (a *adminRepository) Insert(ctx context.Context, admin *domain.Admin) error {
	return a.db.WithContext(ctx).Table("admins").Create(admin).Error
}

// Update implements domain.AdminRepository.
func (a *adminRepository) Update(ctx context.Context, admin *domain.Admin) error {
	return a.db.WithContext(ctx).Table("admins").Model(&domain.Availability{}).Where("id = ?", admin.ID).Updates(&admin).Error
}
