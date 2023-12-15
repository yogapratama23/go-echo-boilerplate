package routes

import (
	"github.com/yogapratama23/go-echo-boilerplate/handlers"

	"github.com/labstack/echo/v4"
)

func RegisterUserRoute(g *echo.Group, h handlers.UserHandlerInterface) {
	user := g.Group("/user")
	user.GET("", h.GetUsers)
}
