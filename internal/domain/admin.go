package domain

import (
	"campsite/internal/dto"
	"context"
)

type Admin struct {
	ID       int64 `gorm:"primaryKey"`
	Name     string
	Email    string
	Password string
}

type AdminRepository interface {
	FindByID(ctx context.Context, id int64) (Admin, error)
	GetAll(ctx context.Context) ([]Admin, error)
	Insert(ctx context.Context, admin *Admin) error
	Update(ctx context.Context, admin *Admin) error
}

type AdminService interface {
	GetAllCampsites(ctx context.Context) dto.Response
	GetAllBookings(ctx context.Context) dto.Response
	GetAllReviews(ctx context.Context) dto.Response
	GetAllUsers(ctx context.Context) dto.Response
}
