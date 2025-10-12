package service

import (
	"go-eCommerce/internal/models"
	"go-eCommerce/internal/database"
	"github.com/gin-gonic/gin"
)

func SignUp(username string, email string, password string, role string) int {
	var hash string
	if HashPassword(&hash, password) != 0 {
		return 1
	}

	var user models.User = models.User{Username: username, Email: email, Password: string(hash), Role: role}
	if database.CreateUser(&user) != 0 {
		return 1
	}
	return 0
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