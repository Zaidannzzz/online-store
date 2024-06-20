package services

import (
	"online-store/httpserver/dto"
	"online-store/httpserver/models"
	"online-store/httpserver/repositories"
)

type CategoryService interface {
	GetAllCategories() ([]models.Category, error)
	GetCategoryByID(id uint) (*models.Category, error)
	CreateCategory(dto *dto.CategoryDTO) (*models.Category, error)
}

type categoryService struct {
	categoryRepository repositories.CategoryRepository
}

func NewCategoryService(categoryRepository repositories.CategoryRepository) CategoryService {
	return &categoryService{categoryRepository}
}

func (s *categoryService) GetAllCategories() ([]models.Category, error) {
	return s.categoryRepository.GetAllCategories()
}

func (s *categoryService) GetCategoryByID(id uint) (*models.Category, error) {
	return s.categoryRepository.GetCategoryByID(id)
}

func (s *categoryService) CreateCategory(dto *dto.CategoryDTO) (*models.Category, error) {
	category := &models.Category{
		Name: dto.Name,
	}
	return s.categoryRepository.CreateCategory(category)
}
