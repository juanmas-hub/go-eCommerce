package database

import (
	"go-eCommerce/initializers"
	"go-eCommerce/internal/models"

)
func CreateUser[T any](user *T) int {

	result := initializers.DB.Create(user)

	if result.Error != nil{
		return 1
	}
	return 0
}

func CheckExistencebyEmail(user *models.User, email string) int {
	initializers.DB.First(user, "email = ?", email)

	if user.ID == 0{
		// log error
		return 1
	}
	return 0
}

func GetRole(id uint) string {
	var user models.User

	result := initializers.DB.First(&user, id)
	if result.Error != nil {
		// loguear
		return ""
	}

	return user.Role
}