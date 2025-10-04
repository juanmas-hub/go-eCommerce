package models

import (
	"time"
)

type Order struct {
	ID          uint      `gorm:"primaryKey"`
	UserID      uint      `gorm:"index"`
	Status      string    `gorm:"size:20;default:pending"`
	TotalAmount float64   `gorm:"not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time

	Items []OrderItem `gorm:"foreignKey:OrderID"`
}