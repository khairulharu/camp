package domain

import "campsite/internal/dto"

type Admin struct {
	ID       int64 `gorm:"primaryKey"`
	Name     string
	Email    string
	Password string
}

type AdminService interface {
	GetAllCampsites() dto.Response
	GetAllBookings() dto.Response
	GetAllReviews() dto.Response
	GetAllUsers() dto.Response
}
