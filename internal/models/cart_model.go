package models

import (
	"time"
)

type Cart struct {
	ID        uint      `gorm:"primaryKey"`
	UserID    uint      `gorm:"uniqueIndex"` // un carrito por usuario
	CreatedAt time.Time
	UpdatedAt time.Time

	Items []CartItem `gorm:"foreignKey:CartID"`
}