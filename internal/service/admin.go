package service

import (
	"campsite/internal/domain"
	"campsite/internal/dto"
	"context"
)

type adminService struct {
	bookingRepository  domain.BookingRepository
	campsiteRepository domain.CampsiteRepository
	reviewRepository   domain.ReviewRepository
}

func NewAdmin(userRepository domain.UserRepository,
	bookingRepository domain.BookingRepository,
	campsiteRepository domain.CampsiteRepository,
	reviewRepository domain.ReviewRepository) domain.AdminService {
	return &adminService{
		bookingRepository:  bookingRepository,
		reviewRepository:   reviewRepository,
		campsiteRepository: campsiteRepository,
	}
}

// GetAllBookings implements domain.AdminService.
func (a *adminService) GetAllBookings(ctx context.Context) dto.Response {
	bookings, err := a.bookingRepository.GetAll(ctx)

	if err != nil {
		return dto.Response{
			Status:  dto.INTERNALSERVERERROR,
			Message: "Failed get bookings",
		}
	}

	var bookingsResponse []dto.BookingResponse

	for _, value := range bookings {
		bookingsResponse = append(bookingsResponse, dto.BookingResponse{
			ID:             value.ID,
			CampsiteID:     value.CampsiteID,
			UserID:         value.UserID,
			CheckInDate:    value.CheckInDate,
			CheckOutDate:   value.CheckOutDate,
			NumberOfPeople: value.NumberOfPeople,
			TotalCost:      value.TotalCost,
			Status:         value.Status,
		})
	}

	return dto.Response{
		Status:  dto.OK,
		Message: "accpeted",
		Data:    bookingsResponse,
	}
}

// GetAllCampsites implements domain.AdminService.
func (a *adminService) GetAllCampsites(ctx context.Context) dto.Response {
	campsites, err := a.campsiteRepository.GetAll(ctx)

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

// GetAllReviews implements domain.AdminService.
func (a *adminService) GetAllReviews(ctx context.Context) dto.Response {
	reviews, err := a.reviewRepository.GetAll(ctx)

	if err != nil {
		return dto.Response{
			Status:  dto.FORBIDDEN,
			Message: "cannot gett all Review",
			Error:   err.Error(),
		}
	}

	var reviewsResponse []dto.ReviewResponse

	for _, value := range reviews {
		reviewsResponse = append(reviewsResponse, dto.ReviewResponse{
			ID:         value.ID,
			CampsiteID: value.CampsiteID,
			UserID:     value.UserID,
			Rating:     value.Rating,
			Comment:    value.Comment,
			ReviewDate: value.ReviewDate,
		})
	}

	return dto.Response{
		Status:  dto.OK,
		Message: "accpted",
		Data:    reviewsResponse,
	}
}
