package controllers

import (
	"net/http"
	//"go-eCommerce/internal/models"
	"go-eCommerce/internal/dto"
	"go-eCommerce/internal/service"

	"github.com/gin-gonic/gin"
)

func ListProducts(c *gin.Context) {
	var products []string

	if service.ListProducts(&products) != 0 {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to list products",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": products,
	})
}

func AddProduct(c *gin.Context) {
	var body struct{
		UserId 		uint
		Product		dto.ProductDTO
	}

	if Bind(c, &body) != 0 {
		return
	}

	var roleRequired string = "admin"

	if service.CheckRole(body.UserId, roleRequired) != 0{
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "the user is not an admin",
		})
		return
	}	

	if service.AddProduct(body.Product) != 0 {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to add a product",
		})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{})
}
