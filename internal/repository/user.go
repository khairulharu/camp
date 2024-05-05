package repository

import (
	"campsite/internal/domain"
	"database/sql"
)

type userRepository struct {
	db *sql.DB //change this into *gorm.DB
}

func NewUserRepository(db *sql.DB) domain.UserRepository {
	return &userRepository{
		db: db,
	}
}

func (user *userRepository) FindByID(userID int64) *domain.User {
	panic("dosnt implement")
}

func (user *userRepository) FindByUsername(username string) *domain.User {
	panic("dosnt implement")
}
