package repositories

import (
	"online-store/httpserver/models"

	"gorm.io/gorm"
)

type CartRepository interface {
	AddToCart(cartItem *models.CartItem) error
	GetCartByUserID(userID uint) (*models.Cart, error)
	RemoveFromCart(cartItemID uint) error
	ClearCart(userID uint) error
}

type cartRepository struct {
	db *gorm.DB
}

func NewCartRepository(db *gorm.DB) CartRepository {
	return &cartRepository{db: db}
}

func (r *cartRepository) AddToCart(cartItem *models.CartItem) error {
	return r.db.Create(cartItem).Error
}

func (r *cartRepository) GetCartByUserID(userID uint) (*models.Cart, error) {
	var cart models.Cart
	err := r.db.Preload("CartItems.Product").Where("user_id = ?", userID).First(&cart).Error
	return &cart, err
}

func (r *cartRepository) RemoveFromCart(cartItemID uint) error {
	return r.db.Delete(&models.CartItem{}, cartItemID).Error
}

func (r *cartRepository) ClearCart(userID uint) error {
	return r.db.Where("user_id = ?", userID).Delete(&models.CartItem{}).Error
}
