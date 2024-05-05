package service

import (
	"campsite/internal/domain"
	"campsite/internal/dto"
	"context"
)

type campsiteService struct {
	campsiteRepository *domain.CampsiteRepository
}

func NewCampsite(campsiteRepository *domain.CampsiteRepository) domain.CampsiteService {
	return &campsiteService{
		campsiteRepository: campsiteRepository,
	}
}

// AddCampsite implements domain.CampsiteService.
func (c *campsiteService) AddCampsite(ctx context.Context, request dto.CampsiteRequest) dto.Response {
	panic("unimplemented")
}

// DeleteCampsite implements domain.CampsiteService.
func (c *campsiteService) DeleteCampsite(ctx context.Context, request dto.CampsiteRequest) dto.Response {
	panic("unimplemented")
}

// GetAllCampsites implements domain.CampsiteService.
func (c *campsiteService) GetAllCampsites() dto.Response {
	panic("unimplemented")
}

// GetCampsite implements domain.CampsiteService.
func (c *campsiteService) GetCampsite(ctx context.Context, id int64) dto.Response {
	panic("unimplemented")
}

// UpdateCampsite implements domain.CampsiteService.
func (c *campsiteService) UpdateCampsite(ctx context.Context, request dto.CampsiteRequest) dto.Response {
	panic("unimplemented")
}
