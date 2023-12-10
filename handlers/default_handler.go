package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type DefaultHandlerInterface interface {
	Home(c echo.Context) error
}

type DefaultHandler struct{}

func NewDefaultHandler() DefaultHandlerInterface {
	return &DefaultHandler{}
}

func (h *DefaultHandler) Home(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Hello world",
		"data":    nil,
	})
}
