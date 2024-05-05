package domain

import (
	"campsite/internal/dto"
	"context"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID          int64          `gorm:"primaryKey"`
	Name        string         `db:"name"`
	Email       string         `db:"email"`
	Password    string         `db:"password"`
	PhoneNumber string         `db:"phone_number"`
	Address     string         `db:"address"`
	CreatedAt   time.Time      `gorm:"type:datetime(3)"`
	UpdatedAt   time.Time      `gorm:"type:datetime(3)"`
	DeletedAt   gorm.DeletedAt `gorm:"type:datetime(3)"`
}

type UserRepository interface {
	GetAll(ctx context.Context) ([]User, error)
	Insert(ctx context.Context, user *User) error
	FindByID(ctx context.Context, userID int64) (User, error)
	FindByUsername(ctx context.Context, username string) (User, error)
	UpdateUser(ctx context.Context, user *User) error
}

type UserService interface {
	GetUser(ctx context.Context, id int64) dto.Response
	UpdateUser(ctx context.Context, request dto.UserRequest) dto.Response
}
