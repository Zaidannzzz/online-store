package main

import (
	"log"
	// "online-store/config"
	// "online-store/internal/auth"
	// "online-store/internal/cart"
	// "online-store/internal/category"
	// "online-store/internal/order"
	// "online-store/internal/product"
	// "online-store/internal/user"
	// "online-store/pkg/db"
	// "online-store/pkg/middleware"
	// "online-store/pkg/cache"

	"github.com/gin-gonic/gin"
)

func main() {
	// config.LoadConfig()
	// db.InitDB()
	// cache.InitRedis()
	router := gin.Default()

	// auth.RegisterRoutes(router)
	// user.RegisterRoutes(router)
	// category.RegisterRoutes(router)
	// product.RegisterRoutes(router)
	// cart.RegisterRoutes(router)
	// order.RegisterRoutes(router)

	// router.Use(middleware.AuthMiddleware())

	if err := router.Run(":8080"); err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}
