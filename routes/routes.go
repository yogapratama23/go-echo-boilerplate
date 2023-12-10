package routes

import (
	"database/sql"
	"test-echo/database/repositories"
	"test-echo/handlers"
	"test-echo/services"

	"github.com/labstack/echo/v4"
)

type RouteInterface interface {
	Register(v1 *echo.Group)
}

type Route struct {
	DB *sql.DB
}

func NewRoute(db *sql.DB) RouteInterface {
	return &Route{
		DB: db,
	}
}

func (r *Route) Register(v1 *echo.Group) {
	// register repositories
	userRepository := repositories.NewUserRepository(r.DB)

	// register services
	userService := services.NewUserService(userRepository)

	// register handlers
	defaultHandler := handlers.NewDefaultHandler()
	userHandler := handlers.NewUserHandler(userService)

	// register routes
	v1.GET("", defaultHandler.Home)

	user := v1.Group("/user")
	user.GET("", userHandler.GetUsers)
}
