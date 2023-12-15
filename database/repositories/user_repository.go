package repositories

import (
	"math"

	"github.com/yogapratama23/go-echo-boilerplate/database/models"
	"github.com/yogapratama23/go-echo-boilerplate/dto"
	"github.com/yogapratama23/go-echo-boilerplate/helpers"

	"gorm.io/gorm"
)

type UserRepositoryInterface interface {
	FindMany(filters *dto.FindManyUserFilter) ([]*models.APIUser, error)
	FindManyPagination(filters *dto.FindManyUserFilter, paginationInput dto.PaginationInput) (*models.APIUserPaginate, error)
}

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepositoryInterface {
	return &UserRepository{
		DB: db,
	}
}

func (r *UserRepository) FindMany(filters *dto.FindManyUserFilter) ([]*models.APIUser, error) {
	var users []*models.APIUser
	qb := r.DB.Model(&models.User{})

	if filters != nil {
		if len(filters.IDs) > 0 {
			qb.Where("id IN ?", filters.IDs)
		}
		if filters.Fullname != "" {
			fullname := "%" + filters.Fullname + "%"
			qb.Where("fullname LIKE ?", fullname)
		}
		if filters.Username != "" {
			username := "%" + filters.Username + "%"
			qb.Where("username LIKE ?", username)
		}
	}

	if err := qb.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (r *UserRepository) FindManyPagination(filters *dto.FindManyUserFilter, paginationInput dto.PaginationInput) (*models.APIUserPaginate, error) {
	var users models.APIUserPaginate
	var pagination helpers.PaginationHelper
	qb := r.DB.Model(&models.User{})

	if filters != nil {
		if len(filters.IDs) > 0 {
			qb.Where("id IN ?", filters.IDs)
		}
		if filters.Fullname != "" {
			fullname := "%" + filters.Fullname + "%"
			qb.Where("fullname LIKE ?", fullname)
		}
		if filters.Username != "" {
			username := "%" + filters.Username + "%"
			qb.Where("username LIKE ?", username)
		}
	}

	var total int64
	qb.Count(&total)

	paginationOptions := pagination.PaginationOptions(&paginationInput)
	if err := qb.Limit(paginationOptions.Limit).Offset(paginationOptions.Offset).Find(&users.Users).Error; err != nil {
		return nil, err
	}

	users.PaginationResponse.Total = int(total)
	users.PaginationResponse.Limit = paginationInput.Limit
	users.PaginationResponse.Page = paginationInput.Page
	users.PaginationResponse.TotalPage = int(math.Ceil(float64(total) / float64(paginationOptions.Limit)))

	return &users, nil
}
