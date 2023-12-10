package routes

import (
	"test-echo/handlers"

	"github.com/labstack/echo/v4"
)

func RegisterUserRoute(g *echo.Group, h handlers.UserHandlerInterface) {
	user := g.Group("/user")
	user.GET("", h.GetUsers)
}
