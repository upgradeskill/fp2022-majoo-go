package main

import (
	"log"
	config "mini-pos/configs"
	controller "mini-pos/controllers"
	"mini-pos/structs"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config.ConnectDB()
	// Configure middleware with the custom claims type
	config := middleware.JWTConfig{
		Claims:     &structs.JwtCustomClaims{},
		SigningKey: []byte(os.Getenv("JWT_SECRET")),
	}

	route := echo.New()

	// Middleware
	route.Use(middleware.Logger())
	route.Use(middleware.Recover())

	// Login route
	route.POST("/login", controller.Login)
	route.GET("/redis-ping", controller.RedisPing)

	// Restricted group
	userRoute := route.Group("/user")
	userRoute.Use(middleware.JWTWithConfig(config))
	userRoute.GET("", controller.UserList)
	userRoute.GET("/profile", controller.Profile)
	userRoute.POST("/store", controller.UserStore)
	userRoute.GET("/show/:id", controller.UserShow)
	userRoute.PUT("/update/:id", controller.UserUpdate)
	userRoute.DELETE("/delete/:id", controller.UserDelete)

	categoryRoute := route.Group("/category")
	categoryRoute.Use(middleware.JWTWithConfig(config))
	categoryRoute.GET("", controller.CategoryList)
	categoryRoute.POST("/store", controller.CategoryStore)
	categoryRoute.GET("/show/:id", controller.CategoryShow)
	categoryRoute.PUT("/update/:id", controller.CategoryUpdate)
	categoryRoute.DELETE("/delete/:id", controller.CategoryDelete)

	outletRoute := route.Group("/outlet")
	outletRoute.Use(middleware.JWTWithConfig(config))
	outletRoute.GET("", controller.OutletList)
	outletRoute.POST("/store", controller.OutletStore)
	// outletRoute.GET("/show/:id", controller.OutletShow)
	// outletRoute.PUT("/update/:id", controller.OutletUpdate)
	// outletRoute.DELETE("/delete/:id", controller.OutletDelete)

	route.Start(":9000")
}
