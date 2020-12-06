package handler

import (
	"github.com/jrallison/go-workers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func (h *Handler) RegisterRoutes(v1 *echo.Group) {

	v1.GET("/health-check", func(context echo.Context) error {
		return context.JSON(http.StatusOK, map[string]string{"success": "true", "version": "v1"})
	})

	v1.GET("/test-consumer", func(context echo.Context) error {
		res, err := workers.Enqueue("myqueue", "Add", "hello-my-worker")
		if err != nil {
			return context.JSON(http.StatusOK, err)
		}
		return context.JSON(http.StatusOK, res)
	})

	v1.Use(middleware.Static("/swagger"))

	users := v1.Group("/users")
	users.POST("/login", h.Login)
	users.POST("", h.CreateUser)
	users.GET("/:id", h.GetUserByID)

}
