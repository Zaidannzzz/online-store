package models

import (
	"gorm.io/gorm"
)

// User represents a user of the online store
type User struct {
	gorm.Model
	Name     string  `json:"name"`
	Email    string  `json:"email" gorm:"unique"`
	Password string  `json:"password"`
	Orders   []Order `json:"orders" gorm:"foreignKey:UserID"`
	Carts    []Cart  `json:"carts" gorm:"foreignKey:UserID"`
}
