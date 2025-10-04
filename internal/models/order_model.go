package models

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	UserID      uint      `gorm:"index"`
	Status      string    `gorm:"size:20;default:pending"`
	TotalAmount float64   `gorm:"not null"`

	Items []OrderItem `gorm:"foreignKey:OrderID"`
}