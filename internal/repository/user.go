package repository

import (
	"campsite/internal/domain"
	"context"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB //change this into *gorm.DB
}

func NewUserRepository(db *gorm.DB) domain.UserRepository {
	return &userRepository{
		db: db,
	}
}

func (u *userRepository) GetAll(ctx context.Context) ([]domain.User, error) {
	var users []domain.User
	err := u.db.WithContext(ctx).Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

// FindByID implements domain.UserRepository.
func (u *userRepository) FindByID(ctx context.Context, userID int64) (domain.User, error) {
	var user domain.User
	err := u.db.WithContext(ctx).Table("users").Where("id = ?", userID).First(user).Error
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}

// FindByUsername implements domain.UserRepository.
func (u *userRepository) FindByUsername(ctx context.Context, username string) (domain.User, error) {
	var user domain.User
	err := u.db.WithContext(ctx).Table("users").Where("name = ?", username).First(user).Error
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}

// Insert implements domain.UserRepository.
func (u *userRepository) Insert(ctx context.Context, user *domain.User) error {
	return u.db.WithContext(ctx).Table("users").Create(user).Error
}

// UpdateUser implements domain.UserRepository.
func (u *userRepository) UpdateUser(ctx context.Context, user *domain.User) error {
	return u.db.WithContext(ctx).Table("users").Model(&domain.User{}).Where("id = ?", user.ID).Updates(&user).Error
}
