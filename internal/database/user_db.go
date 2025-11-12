package database

import (
	"go-eCommerce/initializers"
	"go-eCommerce/internal/models"

)
func CreateUser(user *models.User) error {
    tx := initializers.DB.Begin()
    if tx.Error != nil {
        return tx.Error
    }

    if result := tx.Omit("Cart").Create(user); result.Error != nil {
        tx.Rollback() 
        return result.Error
    }

    if user.Cart != nil {
        user.Cart.UserID = user.ID

		user.Cart.ID = 0

        if result := tx.Create(user.Cart); result.Error != nil {
            tx.Rollback()
            return result.Error
        }
    }

    tx.Commit()
    return nil
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

func GetCartID(id uint) (uint, error) {
	var user models.User

	result := initializers.DB.First(&user, id)
	if result.Error != nil {
		// loguear
		return 0,result.Error
	}

	return user.Cart.ID, nil
}