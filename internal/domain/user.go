package domain

import (
	"campsite/internal/dto"
	"context"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID          int64  `gorm:"primaryKey"`
	Name        string `db:"name"`
	Email       string `db:"email"`
	Password    string `db:"password"`
	PhoneNumber string `db:"phone_number"`
	Address     string `db:"address"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt
}

type UserRepository interface {
	GetAll(ctx context.Context) ([]User, error)
	Insert(ctx context.Context, user *User) error
	FindByID(ctx context.Context, userID int64) (User, error)
	FindByUsername(ctx context.Context, username string) (User, error)
	Update(ctx context.Context, user *User) error
}

type UserService interface {
	AuthUser(ctx context.Context, request dto.UserRequest) dto.UserResponse
	GetUser(ctx context.Context, id int64) dto.Response
	UpdateUser(ctx context.Context, request dto.UserRequest) dto.Response
	SignUp(ctx context.Context, request dto.UserRequest) dto.Response
}
