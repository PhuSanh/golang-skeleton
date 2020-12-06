package handler

import (
	"github.com/labstack/echo/v4"
	"golang-skeleton/config"
	"golang-skeleton/domain/user"
	"golang-skeleton/utils"
	"net/http"
)

type Handler struct {
	cfg        *config.Config
	userDomain user.IUserService
}

func NewHandler(cfg *config.Config, userDomain user.IUserService) *Handler {
	return &Handler{
		cfg:        cfg,
		userDomain: userDomain,
	}
}

func responseSuccess(c echo.Context, data interface{}) error {
	return c.JSON(http.StatusOK, data)
}

func responseError(c echo.Context, err error) error {
	return c.JSON(http.StatusBadRequest, utils.NewError(err))
}
