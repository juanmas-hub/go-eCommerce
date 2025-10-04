package models

import (
	"time"
)

type OrderItem struct {
	ID             uint      `gorm:"primaryKey"`
	OrderID        uint      `gorm:"index"`
	ProductID      uint      `gorm:"index"`
	Quantity       int       `gorm:"not null"`
	PriceAtPurchase float64  `gorm:"not null"`
	CreatedAt      time.Time
	UpdatedAt      time.Time

	Product Product `gorm:"foreignKey:ProductID"`
}