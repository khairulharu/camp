package service

import (
	"campsite/internal/domain"
	"campsite/internal/dto"
	"context"
	"time"
)

type userService struct {
	userRepository domain.UserRepository
}

func NewUser(userRepository domain.UserRepository) domain.UserService {
	return &userService{
		userRepository: userRepository,
	}
}

// GetUser implements domain.UserService.
func (u *userService) GetUser(ctx context.Context, id int64) dto.Response {
	user, err := u.userRepository.FindByID(ctx, id)

	if err != nil {
		return dto.Response{
			Status:  "401",
			Message: "not found",
			Error:   err.Error(),
		}
	}

	return dto.Response{
		Status:  "200",
		Message: "accepted",
		Data: dto.UserResponse{
			ID:          user.ID,
			Name:        user.Name,
			Email:       user.Email,
			Password:    user.Password,
			PhoneNumber: user.PhoneNumber,
			Address:     user.Address,
		},
	}
}

// UpdateUser implements domain.UserService.
func (u *userService) UpdateUser(ctx context.Context, request dto.UserRequest) dto.Response {
	user := domain.User{
		ID:          request.ID,
		Name:        request.Name,
		Email:       request.Email,
		Password:    request.Password,
		PhoneNumber: request.PhoneNumber,
		Address:     request.Address,
		UpdatedAt:   time.Now(),
	}

	if err := u.userRepository.Update(ctx, &user); err != nil {
		return dto.Response{
			Status:  "401",
			Message: "failed update",
			Error:   err.Error(),
		}
	}

	return dto.Response{
		Status:  "200",
		Message: "user updated",
	}
}

// AuthUser implements domain.UserService.
func (u *userService) AuthUser(ctx context.Context, request dto.UserRequest) dto.UserResponse {
	user, err := u.userRepository.FindByUsername(ctx, request.Name)

	if request.Password != user.Password {
		return dto.UserResponse{}
	}

	if err != nil {
		return dto.UserResponse{}
	}

	return dto.UserResponse{
		ID:          user.ID,
		Name:        user.Name,
		Email:       user.Email,
		Password:    user.Password,
		PhoneNumber: request.PhoneNumber,
		Address:     request.Address,
	}
}

// SignUp implements domain.UserService.
func (u *userService) SignUp(ctx context.Context, request dto.UserRequest) dto.Response {

	userCheckName, _ := u.userRepository.FindByUsername(ctx, request.Name)

	if userCheckName != (domain.User{}) {
		return dto.Response{
			Status:  "401",
			Message: "user exist cannot create",
			Error:   "user Exist",
		}
	}

	err := u.userRepository.Insert(ctx, &domain.User{
		Name:        request.Name,
		Email:       request.Email,
		Password:    request.Password,
		PhoneNumber: request.PhoneNumber,
		Address:     request.Address,
	})

	if err != nil {
		return dto.Response{
			Status:  "500",
			Message: "something wrong",
			Error:   err.Error(),
		}
	}

	return dto.Response{
		Status:  "200",
		Message: "accepted",
	}

}
