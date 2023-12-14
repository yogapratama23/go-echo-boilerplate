package services

import (
	"test-echo/database/models"
	"test-echo/database/repositories"
	"test-echo/dto"
)

type UserServiceInterface interface {
	FindMany(input *dto.GetUsersPayload) (*models.APIUserPaginate, error)
}

type UserService struct {
	UserRepository repositories.UserRepositoryInterface
}

func NewUserService(userRepository repositories.UserRepositoryInterface) UserServiceInterface {
	return &UserService{
		UserRepository: userRepository,
	}
}

func (s *UserService) FindMany(input *dto.GetUsersPayload) (*models.APIUserPaginate, error) {
	users, err := s.UserRepository.FindManyPagination(&input.FindManyUserFilter, input.PaginationInput)
	if err != nil {
		return nil, err
	}

	return users, nil
}
