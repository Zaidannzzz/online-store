package main

import (
	"fmt"
	"log"
	"net/http"
	"online-store/config/db"
	"online-store/httpserver/controller/cart"
	"online-store/httpserver/controller/order"
	"online-store/httpserver/controller/product"
	"online-store/httpserver/middleware"
	"online-store/httpserver/repositories"
	"online-store/httpserver/routers"
	"online-store/httpserver/services"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Failed to load .env file: %v", err)
	}

	gin.SetMode(gin.ReleaseMode)

	app := gin.Default()
	db, err := db.Connect()
	if err != nil {
		log.Fatal("Failed to connect to the database")
	}

	app.GET("/", func(c *gin.Context) {
		log.Println("Accessed / endpoint")
		c.String(http.StatusOK, "Hello, World!")
	})

	// Repositories
	cartRepo := repositories.NewCartRepository(db)
	orderRepo := repositories.NewOrderRepository(db)
	productRepo := repositories.NewProductRepository(db)

	// Services
	cartService := services.NewCartService(cartRepo)
	orderService := services.NewOrderService(cartRepo, orderRepo)
	productService := services.NewProductService(productRepo)

	// Controllers
	cartHandler := cart.NewCartHandler(cartService)
	orderHandler := order.NewOrderHandler(orderService)
	productHandler := product.NewProductHandler(productService)

	// Middleware
	app.Use(middleware.AuthMiddleware())

	// Routes
	routers.SetupCartRoutes(app, cartHandler)
	routers.SetupOrderRoutes(app, orderHandler)
	routers.SetupProductRoutes(app, productHandler)

	err = app.Run(":8080")
	if err != nil {
		log.Fatal("Failed to start the server")
	}
	fmt.Println("Starting the server on port 8080")

}
