package models

import (
	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	UserID    uint      `gorm:"uniqueIndex"` // un carrito por usuario

	Items []CartItem `gorm:"foreignKey:CartID"`
}