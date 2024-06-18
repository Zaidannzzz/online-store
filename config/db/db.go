package db

import (
	"fmt"
	"os"

	"online-store/httpserver/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))

	// Connect to the database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Perform database migrations or other setup tasks here
	err = runMigrations(db)
	if err != nil {
		return nil, err
	}

	fmt.Println("Connected to the database")

	return db, nil
}

func runMigrations(db *gorm.DB) error {
	err := db.AutoMigrate(&models.User{}, &models.Product{}, &models.Category{}, &models.Cart{}, &models.CartItem{}, &models.Order{})
	if err != nil {
		return err
	}

	fmt.Println("Database migrations completed")

	return nil
}
