package service

import (
	"campsite/internal/domain"
	"campsite/internal/dto"
	"context"
)

type userService struct {
	userRepository *domain.UserRepository
}

func NewUser(userRepository *domain.UserRepository) domain.UserService {
	return &userService{
		userRepository: userRepository,
	}
}

// GetUser implements domain.UserService.
func (u *userService) GetUser(ctx context.Context, id int64) dto.Response {
	panic("not implementde")
}

// UpdateUser implements domain.UserService.
func (u *userService) UpdateUser(ctx context.Context, request dto.UserRequest) dto.Response {
	panic("unimplemented")
}
