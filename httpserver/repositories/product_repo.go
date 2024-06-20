package repositories

import (
	"online-store/httpserver/models"

	"gorm.io/gorm"
)

type ProductRepository interface {
	GetProductsByCategory(categoryID uint) ([]models.Product, error)
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db: db}
}

func (r *productRepository) GetProductsByCategory(categoryID uint) ([]models.Product, error) {
	var products []models.Product
	if err := r.db.Where("category_id = ?", categoryID).Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}
