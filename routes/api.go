package routes

import (
	"mini-pos/configs"
	controller "mini-pos/controllers"
	"mini-pos/structs"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Api() {
	configs.ConnectDB()

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
	route.POST("/register", controller.Register)
	route.GET("/redis-ping", controller.RedisPing)

	// Restricted group
	userRoute := route.Group("/user")
	userRoute.Use(middleware.JWTWithConfig(config))
	userRoute.GET("", controller.UserList)
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

	productRoute := route.Group("/product")
	productRoute.Use(middleware.JWTWithConfig(config))
	productRoute.GET("", controller.ProductList)
	productRoute.POST("/store", controller.ProductStore)
	productRoute.GET("/show/:id", controller.ProductShow)
	productRoute.PUT("/update/:id", controller.ProductUpdate)
	productRoute.DELETE("/delete/:id", controller.ProductDelete)

	transactionRoute := route.Group("/transaction")
	transactionRoute.Use(middleware.JWTWithConfig(config))
	transactionRoute.GET("", controller.TransactionList)
	transactionRoute.POST("/store", controller.TransactionStore)
	transactionRoute.GET("/show/:id", controller.TransactionShow)
	route.Start(":9000")
}
