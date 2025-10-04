package models
import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string    `gorm:"size:100;not null"`
	Price       float64   `gorm:"not null"`
	Stock       int       `gorm:"not null"`
	ImageURL    string    `gorm:"size:255"`

	CartItems  []CartItem  `gorm:"foreignKey:ProductID"`
	OrderItems []OrderItem `gorm:"foreignKey:ProductID"`
	CategoryID *uint
	Category   *Category
}