package repository

import (
	"campsite/internal/domain"
	"context"
	"database/sql"
	"errors"

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
	err := u.db.WithContext(ctx).Table("users").Where("id = ?", &userID).First(&user).Error
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}

// FindByUsername implements domain.UserRepository.
func (u *userRepository) FindByUsername(ctx context.Context, username string) (domain.User, error) {
	var user domain.User
	err := u.db.WithContext(ctx).Table("users").Where("name = ?", &username).First(&user).Error
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}

// Insert implements domain.UserRepository.
func (u *userRepository) Insert(ctx context.Context, user *domain.User) error {
	return u.db.WithContext(ctx).Table("users").Create(&user).Error
}

// UpdateUser implements domain.UserRepository.
func (u *userRepository) Update(ctx context.Context, user *domain.User) error {
	return u.db.WithContext(ctx).Table("users").Model(&domain.User{}).Where("id = ?", &user.ID).Updates(&user).Error
}

type userRepositoryRare struct {
	dbRare *sql.DB
	dbGorm *gorm.DB
}

func NewUserRepositoryRare(dbRare *sql.DB, dbGorm *gorm.DB) domain.UserRepository {
	return &userRepositoryRare{
		dbRare: dbRare,
		dbGorm: dbGorm,
	}
}

// FindByID implements domain.UserRepository.
func (u *userRepositoryRare) FindByID(ctx context.Context, userID int64) (domain.User, error) {
	dbRecord, err := u.dbGorm.ConnPool.QueryContext(ctx, `SELECT * FROM users WHERE id = ?`, userID)

	if err != nil {
		return domain.User{}, err
	}

	userReq := domain.User{}

	if err := dbRecord.Scan(&userReq.ID, &userReq.Name, &userReq.Email, &userReq.Password, &userReq.PhoneNumber, &userReq.Address, &userReq.CreatedAt, &userReq.UpdatedAt, &userReq.DeletedAt); err != nil {
		if !dbRecord.Next() {
			return domain.User{}, errors.New("scan method called but failed")
		}
		return domain.User{}, err
	}

	return userReq, nil
}

// FindByUsername implements domain.UserRepository.
func (u *userRepositoryRare) FindByUsername(ctx context.Context, username string) (domain.User, error) {
	panic("unimplemented")
}

// GetAll implements domain.UserRepository.
func (u *userRepositoryRare) GetAll(ctx context.Context) ([]domain.User, error) {
	panic("unimplemented")
}

// Insert implements domain.UserRepository.
func (u *userRepositoryRare) Insert(ctx context.Context, user *domain.User) error {

	_, err := u.dbGorm.ConnPool.ExecContext(ctx, `INSERT INTO users (id, name, email, password, phone_number, address, created_at) 
	VALUES (?, ?, ?, ?, ?, ?, ?)`, user.ID, user.Name, user.Email, user.Password, user.PhoneNumber, user.Address, user.CreatedAt)

	if err != nil {
		return err
	}

	return nil
}

// Update implements domain.UserRepository.
func (u *userRepositoryRare) Update(ctx context.Context, user *domain.User) error {
	panic("unimplemented")
}
