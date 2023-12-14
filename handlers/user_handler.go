package handlers

import (
	"net/http"
	"test-echo/dto"
	"test-echo/services"

	"github.com/labstack/echo/v4"
)

type UserHandlerInterface interface {
	GetUsers(c echo.Context) error
}

type UserHandler struct {
	UserService services.UserServiceInterface
}

func NewUserHandler(userService services.UserServiceInterface) UserHandlerInterface {
	return &UserHandler{
		UserService: userService,
	}
}

func (h *UserHandler) GetUsers(c echo.Context) error {
	var payload dto.GetUsersPayload
	err := (&echo.DefaultBinder{}).BindQueryParams(c, &payload)
	if err != nil {
		return err
	}

	data, err := h.UserService.FindMany(&payload)
	if err != nil {
		return err
	}

	resp := dto.ResponsePagination{
		Message:            "Get users",
		Data:               data.Users,
		PaginationResponse: data.PaginationResponse,
	}
	return c.JSON(http.StatusOK, resp)
}
