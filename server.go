package main

import (
	config "mini-pos/configs"
	controller "mini-pos/controllers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	config.ConnectDB()
	// Configure middleware with the custom claims type
	config := middleware.JWTConfig{
		Claims:     &controller.JwtCustomClaims{},
		SigningKey: []byte("secret"),
	}

	route := echo.New()

	// Middleware
	route.Use(middleware.Logger())
	route.Use(middleware.Recover())

	// Login route
	route.POST("/login", controller.Login)

	// Restricted group
	r := route.Group("/user")
	r.Use(middleware.JWTWithConfig(config))

	r.GET("", controller.List)
	r.POST("/store", controller.Store)
	r.GET("/show/:id", controller.Show)
	r.PUT("/update/:id", controller.Update)
	r.DELETE("/delete/:id", controller.Delete)

	route.Start(":9000")
}
