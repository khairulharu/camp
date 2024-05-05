package domain

import (
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
	FindByID(userID int64) *User
	FindByUsername(username string) *User
}
