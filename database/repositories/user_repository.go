package repositories

import (
	"database/sql"
	"log/slog"
	"test-echo/database/models"
	userqueries "test-echo/database/queries"
)

type UserRepositoryInterface interface {
	GetMany() ([]models.User, error)
}

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepositoryInterface {
	return &UserRepository{
		DB: db,
	}
}

func (r *UserRepository) GetMany() ([]models.User, error) {
	var users []models.User
	rows, err := r.DB.Query(userqueries.GET_MANY)
	if err != nil {
		slog.Error(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var user models.User
		if err = rows.Scan(
			&user.ID, &user.Username, &user.Fullname, &user.CreatedAt,
			&user.UpdatedAt, &user.DeletedAt,
		); err != nil {
			slog.Error(err.Error())
		}
		users = append(users, user)
	}

	return users, nil
}
