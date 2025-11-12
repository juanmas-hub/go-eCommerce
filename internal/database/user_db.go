package database

import (
	"go-eCommerce/initializers"
	"go-eCommerce/internal/models"
	"errors"
)
func CreateUser(user *models.User) error {
    /*tx := initializers.DB.Begin()
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
    return nil*/
	result := initializers.DB.Create(user)
	if result.Error != nil{
		return errors.New("error al crear el usuario: " + result.Error.Error())
	}

	if user.Cart.ID == 0 {
        return errors.New("el carrito no se creó correctamente, revise la configuración de la relación")
    }

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