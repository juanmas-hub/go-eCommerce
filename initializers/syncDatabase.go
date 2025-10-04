package initializers

import (
	"go-eCommerce/internal/models"
)

func SyncDatabase(){
	DB.AutoMigrate(
		&models.User{},
		&models.Product{},
		&models.Category{},
		&models.Cart{},
		&models.CartItem{},
		&models.Order{},
		&models.OrderItem{},
	)
}