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
	userRoute := route.Group("/user")
	userRoute.Use(middleware.JWTWithConfig(config))

	userRoute.GET("", controller.UserList)
	userRoute.POST("/store", controller.UserStore)
	userRoute.GET("/show/:id", controller.UserShow)
	userRoute.PUT("/update/:id", controller.UserUpdate)
	userRoute.DELETE("/delete/:id", controller.UserDelete)

	route.Start(":9000")
}
