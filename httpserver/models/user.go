package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string  `json:"name"`
	Email    string  `json:"email" gorm:"unique"`
	Password string  `json:"password"`
	Orders   []Order `json:"orders" gorm:"foreignKey:UserID"`
	Carts    []Cart  `json:"carts" gorm:"foreignKey:UserID"`
}

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
