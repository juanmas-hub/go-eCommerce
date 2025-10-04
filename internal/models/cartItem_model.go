package models

import (
	"gorm.io/gorm"
)

type CartItem struct {
	gorm.Model
	CartID    uint      `gorm:"index"`
	ProductID uint      `gorm:"index"`
	Quantity  int       `gorm:"not null"`


	Product Product `gorm:"foreignKey:ProductID"`
}