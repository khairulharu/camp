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
			Status:  dto.BADREQUEST,
			Message: "failed add review",
			Error:   err.Error(),
		}
	}

	return dto.Response{
		Status:  dto.CREATED,
		Message: "accepted",
	}
}

// DeleteCampsite implements domain.CampsiteService.
func (c *campsiteService) DeleteCampsite(ctx context.Context, request dto.CampsiteRequest) dto.Response {
	err := c.campsiteRepository.Delete(ctx, &domain.Campsite{ID: request.ID})

	if err != nil {
		return dto.Response{
			Status:  dto.NOTFOUND,
			Message: "failed delete specific campsite",
		}
	}

	return dto.Response{
		Status:  dto.OK,
		Message: "accepted",
	}
}

// GetAllCampsites implements domain.CampsiteService.
func (c *campsiteService) GetAllCampsites(ctx context.Context) dto.Response {
	campsites, err := c.campsiteRepository.GetAll(ctx)

	if err != nil {
		return dto.Response{
			Status:  dto.NOTFOUND,
			Message: "failed get all",
		}
	}

	var campsitesResponse []dto.CampsiteResponse

	for _, value := range campsites {
		campsitesResponse = append(campsitesResponse, dto.CampsiteResponse{
			ID:            value.ID,
			Name:          value.Name,
			Location:      value.Location,
			Latitude:      value.Latitude,
			Longitude:     value.Longitude,
			Area:          value.Area,
			PricePerNight: value.PricePerNight,
		})
	}

	return dto.Response{
		Status:  dto.OK,
		Message: "accepted",
		Data:    campsitesResponse,
	}
}

// GetCampsite implements domain.CampsiteService.
func (c *campsiteService) GetCampsite(ctx context.Context, id int64) dto.Response {
	campsite, err := c.campsiteRepository.FindByID(ctx, id)

	if err != nil {
		return dto.Response{
			Status:  dto.NOTFOUND,
			Message: "campsite not found",
		}
	}

	return dto.Response{
		Status:  dto.OK,
		Message: "accepted",
		Data: dto.CampsiteResponse{
			ID:            campsite.ID,
			Name:          campsite.Name,
			Location:      campsite.Location,
			Latitude:      campsite.Latitude,
			Longitude:     campsite.Longitude,
			Area:          campsite.Area,
			PricePerNight: campsite.PricePerNight,
		},
	}
}

// UpdateCampsite implements domain.CampsiteService.
func (c *campsiteService) UpdateCampsite(ctx context.Context, request dto.CampsiteRequest) dto.Response {
	panic("unimplemented")
}
