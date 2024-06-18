package main

import (
	"fmt"
	"log"
	"net/http"
	"online-store/config/db"
	controllers "online-store/httpserver/controller/auth"
	"online-store/httpserver/repositories"
	"online-store/httpserver/routers"
	"online-store/httpserver/services"
	"online-store/utils"

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
	appRoute := app.Group("/api")
	db, err := db.Connect()
	if err != nil {
		log.Fatal("Failed to connect to the database")
	}

	app.GET("/", func(c *gin.Context) {
		log.Println("Accessed / endpoint")
		c.String(http.StatusOK, "Hello, World!")
	})

	authService := utils.NewAuthHelper(utils.Constants.JWT_SECRET_KEY)

	userRepository := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepository)
	userController := controllers.NewUserController(userService, authService)

	routers.UserRouter(appRoute, userController, authService)

	err = app.Run(":8080")
	if err != nil {
		log.Fatal("Failed to start the server")
	}
	fmt.Println("Starting the server on port 8080")

}
