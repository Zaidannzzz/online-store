package services

import (
	"online-store/httpserver/dto"
	"online-store/httpserver/models"
	"online-store/httpserver/repositories"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Register(dto *dto.RegisterDto) (*models.User, error)
	Login(dto *dto.LoginDto) (*models.User, error)
}

type userService struct {
	userRepository repositories.UserRepository
}

func NewUserService(r repositories.UserRepository) *userService {
	return &userService{r}
}

func (s *userService) Register(dto *dto.RegisterDto) (*models.User, error) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(dto.Password), bcrypt.DefaultCost)

	if err != nil {
		return nil, err
	}

	dto.Password = string(hashedPassword)

	user := models.User{
		Name:     dto.Name,
		Email:    dto.Email,
		Password: dto.Password,
	}

	_, err = s.userRepository.Register(&user)
	if err != nil {
		return &user, err
	}
	return &user, nil
}

func (s *userService) Login(dto *dto.LoginDto) (*models.User, error) {
	user := models.User{
		Email:    dto.Email,
		Password: dto.Password,
	}

	result, err := s.userRepository.Login(&user)
	if err != nil {
		return &user, err
	}

	return result, nil
}
