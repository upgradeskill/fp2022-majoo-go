package user

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func RegisterRouter(
	e *echo.Echo,
	handler *Handler,
) {
	e.GET("/health", func(c echo.Context) error {
		return c.NoContent(http.StatusOK)
	})

	userV1 := e.Group("api/v1/user")
	userV1.GET("", handler.UserList)
	userV1.POST("", handler.UserCreate)
	userV1.GET("/:id", handler.UserFindById)
	userV1.PUT("/:id", handler.UserUpdateById)
	userV1.DELETE("/:id", handler.UserDeleteById)
}
