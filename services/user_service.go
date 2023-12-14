package services

import (
	"test-echo/database/models"
	"test-echo/database/repositories"
	"test-echo/dto"
)

type UserServiceInterface interface {
	FindMany(input *dto.FindManyUserFilter) ([]*models.APIUser, error)
}

type UserService struct {
	UserRepository repositories.UserRepositoryInterface
}

func NewUserService(userRepository repositories.UserRepositoryInterface) UserServiceInterface {
	return &UserService{
		UserRepository: userRepository,
	}
}

func (s *UserService) FindMany(input *dto.FindManyUserFilter) ([]*models.APIUser, error) {
	users, err := s.UserRepository.FindMany(input)
	if err != nil {
		return nil, err
	}

	return users, nil
}
