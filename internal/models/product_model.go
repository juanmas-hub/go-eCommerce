package models
import (
	"time"
)

type Product struct {
	ID          uint      `gorm:"primaryKey"`
	Name        string    `gorm:"size:100;not null"`
	Description string    `gorm:"type:text"`
	Price       float64   `gorm:"not null"`
	Stock       int       `gorm:"not null"`
	ImageURL    string    `gorm:"size:255"`
	CreatedAt   time.Time
	UpdatedAt   time.Time

	CartItems  []CartItem  `gorm:"foreignKey:ProductID"`
	OrderItems []OrderItem `gorm:"foreignKey:ProductID"`
	CategoryID *uint
	Category   *Category
}