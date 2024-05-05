package service

import (
	"campsite/internal/domain"
	"campsite/internal/dto"
	"context"
)

type adminService struct {
	adminRepository *domain.AdminRepository
}

func NewAdmin(adminRepository *domain.AdminRepository) domain.AdminService {
	return &adminService{
		adminRepository: adminRepository,
	}
}

// GetAllBookings implements domain.AdminService.
func (a *adminService) GetAllBookings(ctx context.Context) dto.Response {
	panic("unimplemented")
}

// GetAllCampsites implements domain.AdminService.
func (a *adminService) GetAllCampsites(ctx context.Context) dto.Response {
	panic("unimplemented")
}

// GetAllReviews implements domain.AdminService.
func (a *adminService) GetAllReviews(ctx context.Context) dto.Response {
	panic("unimplemented")
}

// GetAllUsers implements domain.AdminService.
func (a *adminService) GetAllUsers(ctx context.Context) dto.Response {
	panic("unimplemented")
}
