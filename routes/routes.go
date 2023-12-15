package routes

import (
	"github.com/yogapratama23/go-echo-boilerplate/database/repositories"
	"github.com/yogapratama23/go-echo-boilerplate/handlers"
	"github.com/yogapratama23/go-echo-boilerplate/services"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type RouteInterface interface {
	RegisterV1(v1 *echo.Group)
}

type Route struct {
	DB *gorm.DB
}

func NewRoute(db *gorm.DB) RouteInterface {
	return &Route{
		DB: db,
	}
}

func (r *Route) RegisterV1(v1 *echo.Group) {
	// register repositories
	userRepository := repositories.NewUserRepository(r.DB)

	// register services
	userService := services.NewUserService(userRepository)

	// register handlers
	defaultHandler := handlers.NewDefaultHandler()
	userHandler := handlers.NewUserHandler(userService)

	// register routes
	v1.GET("", defaultHandler.Home)
	RegisterUserRoute(v1, userHandler)
}
