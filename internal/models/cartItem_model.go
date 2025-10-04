package models

import (
	"time"
)

type CartItem struct {
	ID        uint      `gorm:"primaryKey"`
	CartID    uint      `gorm:"index"`
	ProductID uint      `gorm:"index"`
	Quantity  int       `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time

	Product Product `gorm:"foreignKey:ProductID"`
}