package repositories

import (
	"online-store/httpserver/models"

	"gorm.io/gorm"
)

type OrderRepository interface {
	CreateOrder(userID uint, items []models.CartItem, total float64) error
}

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepository{db: db}
}

func (r *orderRepository) CreateOrder(userID uint, items []models.CartItem, total float64) error {
	order := models.Order{
		UserID: userID,
		Total:  total,
		Status: "Pending",
		Items:  items,
	}
	return r.db.Create(&order).Error
}
