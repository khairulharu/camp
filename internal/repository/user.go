package repository

import (
	"campsite/internal/domain"
	"context"
	"database/sql"

	"gorm.io/gorm"
)

func UseUserRepository(isUseOrm bool, dbGorm *gorm.DB, mySql *sql.DB) domain.UserRepository {
	if isUseOrm {
		return NewUserRepository(dbGorm)
	}

	return NewUserRepositoryRare(mySql)
}

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
}

func NewUserRepositoryRare(dbRare *sql.DB) domain.UserRepository {
	return &userRepositoryRare{
		dbRare: dbRare,
	}
}

// FindByID implements domain.UserRepository.
func (u *userRepositoryRare) FindByID(ctx context.Context, userID int64) (domain.User, error) {

	conn, err := u.dbRare.Conn(ctx)

	if err != nil {
		return domain.User{}, err
	}

	defer conn.Close()

	rowsRecord := conn.QueryRowContext(ctx, `SELECT * FROM users WHERE id = ?`, userID)

	userReq := domain.User{}

	if err := rowsRecord.Scan(&userReq.ID, &userReq.Name, &userReq.Email, &userReq.Password, &userReq.PhoneNumber, &userReq.Address, &userReq.CreatedAt, &userReq.UpdatedAt, &userReq.DeletedAt); err != nil {
		return domain.User{}, err
	}

	return domain.User{}, nil
}

// FindByUsername implements domain.UserRepository.
func (u *userRepositoryRare) FindByUsername(ctx context.Context, username string) (domain.User, error) {
	conn, err := u.dbRare.Conn(ctx)

	if err != nil {
		return domain.User{}, err
	}

	defer conn.Close()

	recordRows := conn.QueryRowContext(ctx, `SELECT * FROM users WHERE name = ?`, username)

	userReq := domain.User{}

	if err := recordRows.Scan(&userReq.ID, &userReq.Name, &userReq.Email, &userReq.Password, &userReq.PhoneNumber, &userReq.Address, &userReq.CreatedAt, &userReq.UpdatedAt, &userReq.DeletedAt); err != nil {
		return domain.User{}, err
	}

	return domain.User{}, nil
}

// GetAll implements domain.UserRepository.
func (u *userRepositoryRare) GetAll(ctx context.Context) ([]domain.User, error) {
	panic("unimplemented")
}

// Insert implements domain.UserRepository.
func (u *userRepositoryRare) Insert(ctx context.Context, user *domain.User) error {

	conn, err := u.dbRare.Conn(ctx)
	if err != nil {
		return err
	}

	defer conn.Close()

	_, err = conn.ExecContext(ctx, `INSERT INTO users (id, name, email, password, phone_number, address, created_at, updated_at, deleted_at) 
	VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`, user.ID, user.Name, user.Email, user.Password, user.PhoneNumber, user.Address, user.CreatedAt, user.UpdatedAt, user.DeletedAt)

	if err != nil {
		return err
	}

	return nil
}

// Update implements domain.UserRepository.
func (u *userRepositoryRare) Update(ctx context.Context, user *domain.User) error {
	panic("unimplemented")
}
