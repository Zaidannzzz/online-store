package repositories

import (
	"fmt"
	"net/http"
	"online-store/httpserver/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	Register(user *models.User) (*models.User, error)
	Login(user *models.User) (*models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (r *userRepository) Register(user *models.User) (*models.User, error) {
	err := r.db.Create(user).Error
	if err != nil {
		fmt.Println(http.StatusExpectationFailed, "Need Request Body")
		return user, err
	}
	return user, nil
}

func (r *userRepository) Login(user *models.User) (*models.User, error) {
	err := r.db.Where("email = ?", user.Email).First(user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}
