package repositories

import (
	"test-echo/database/models"
	"test-echo/dto"

	"gorm.io/gorm"
)

type UserRepositoryInterface interface {
	FindMany(filters *dto.FindManyUserFilter) ([]*models.APIUser, error)
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
