package controllers

import (
	"go-eCommerce/internal/service"

	"github.com/gin-gonic/gin"
)

func AddProductToCart(c *gin.Context){
	var body struct{
		userID 	  uint
		productID uint
		quantity  int
	}

	if Bind(c, &body) != 0 {
		return
	}

	if service.AddProductToCart(body.userID, body.productID, body.quantity) != nil{
		
	}

}