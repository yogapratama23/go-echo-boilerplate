package services

import (
	"test-echo/database/models"
	"test-echo/database/repositories"
)

type UserServiceInterface interface {
	GetUsers() ([]models.User, error)
}

type UserService struct {
	UserRepository repositories.UserRepositoryInterface
}

func NewUserService(userRepository repositories.UserRepositoryInterface) UserServiceInterface {
	return &UserService{
		UserRepository: userRepository,
	}
}

func (s *UserService) GetUsers() ([]models.User, error) {
	users, err := s.UserRepository.GetMany()
	if err != nil {
		return nil, err
	}

	return users, nil
}
