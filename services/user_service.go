package services

import (
	"fmt"

	"github.com/yogapratama23/go-echo-boilerplate/database/models"
	"github.com/yogapratama23/go-echo-boilerplate/database/repositories"
	"github.com/yogapratama23/go-echo-boilerplate/dto"
)

type UserServiceInterface interface {
	GetUsers(input *dto.GetUsersPayload) (*models.APIUserPaginate, error)
}

type UserService struct {
	UserRepository repositories.UserRepositoryInterface
}

func NewUserService(userRepository repositories.UserRepositoryInterface) UserServiceInterface {
	return &UserService{
		UserRepository: userRepository,
	}
}

func (s *UserService) GetUsers(input *dto.GetUsersPayload) (*models.APIUserPaginate, error) {
	users, err := s.UserRepository.FindManyPagination(&input.FindManyUserFilter, input.PaginationInput)
	if err != nil {
		return nil, fmt.Errorf("[UserService][GetUsers] error : %s", err.Error())
	}

	return users, nil
}
