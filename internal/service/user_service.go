package service

import (
	"go-eCommerce/internal/models"
	"go-eCommerce/internal/database"
	"github.com/gin-gonic/gin"
)

func SignUp(username string, email string, password string, role string) error {
	var hash string

	hashResult := HashPassword(&hash, password)
	if hashResult != nil {
		return hashResult
	}

	var user models.User = models.User{
		Username: username, 
		Email: email, 
		Password: string(hash), 
		Role: role,
		Cart: &models.Cart{},
	}

	result := database.CreateUser(&user) 
	if result != nil {
		return result
	}
	return nil
}

func LogIn(c *gin.Context, email string, password string) int{

	var user models.User
	if database.CheckExistencebyEmail(&user, email) != 0{
		return 1
	}

	if CheckPassword(user, password) != 0{
		return 1
	}
	
	var tokenString string
	if GenerateToken(user, &tokenString) != 0{
		return 1
	}
	
	SetCookie(c, tokenString)

	return 0
}

func CheckRole(userID uint, role string) int {
	var userRole string = database.GetRole(userID)

	if userRole != role{
		//loguearlo
		return 1
	}
	return 0
}