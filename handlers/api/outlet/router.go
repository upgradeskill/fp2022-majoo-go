package outlet

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

	userV1 := e.Group("api/v1/outlet")
	userV1.GET("", handler.OutletList)
	userV1.POST("", handler.OutletCreate)
	userV1.GET("/:id", handler.OutletFindById)
	userV1.PUT("/:id", handler.OutletUpdateById)
	userV1.DELETE("/:id", handler.OutletDeleteById)
}
