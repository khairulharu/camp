package service

import (
	"campsite/internal/domain"
	"campsite/internal/dto"
	"context"
	"time"
)

type campsiteService struct {
	campsiteRepository domain.CampsiteRepository
}

func NewCampsite(campsiteRepository domain.CampsiteRepository) domain.CampsiteService {
	return &campsiteService{
		campsiteRepository: campsiteRepository,
	}
}

// AddCampsite implements domain.CampsiteService.
func (c *campsiteService) AddCampsite(ctx context.Context, request dto.CampsiteRequest) dto.Response {
	campsite := domain.Campsite{
		Name:          request.Name,
		Location:      request.Location,
		Latitude:      request.Latitude,
		Longitude:     request.Longitude,
		Area:          request.Area,
		PricePerNight: request.PricePerNight,
		CreatedAt:     time.Now(),
	}

	if err := c.campsiteRepository.Insert(ctx, &campsite); err != nil {
		return dto.Response{
			Status:  "401",
			Message: "failed add review",
			Error:   err.Error(),
		}
	}

	return dto.Response{
		Status:  "200",
		Message: "accepted",
	}
}

// DeleteCampsite implements domain.CampsiteService.
func (c *campsiteService) DeleteCampsite(ctx context.Context, request dto.CampsiteRequest) dto.Response {
	err := c.campsiteRepository.Delete(ctx, &domain.Campsite{ID: request.ID})

	if err != nil {
		return dto.Response{
			Status: "",
		}
	}

	return dto.Response{
		Status:  "200",
		Message: "accepted",
	}
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
