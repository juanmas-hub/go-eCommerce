package initializers

import (
	"go-eCommerce/internal/models"
)

func SyncDatabase(){
	DB.AutoMigrate(&models.User{})
}