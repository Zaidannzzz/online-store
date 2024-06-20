package services

import (
	"online-store/httpserver/models"
	"online-store/httpserver/repositories"
)

type ProductService interface {
	GetProductsByCategory(categoryID uint) ([]models.Product, error)
}

type productService struct {
	repo repositories.ProductRepository
}

func NewProductService(repo repositories.ProductRepository) ProductService {
	return &productService{repo: repo}
}

func (s *productService) GetProductsByCategory(categoryID uint) ([]models.Product, error) {
	return s.repo.GetProductsByCategory(categoryID)
}
