package service

import (
	"go-eCommerce/internal/database"
	"go-eCommerce/internal/dto"
	"go-eCommerce/internal/models"
)

func ListProducts(products *[]string) int{
	list := database.GetProducts()

	if list == nil{
		return 1
	}
	*products = list
	return 0
}

func AddProduct(productDTO dto.ProductDTO) int {

	var product models.Product

	product.Name = productDTO.Name
	product.Price = productDTO.Price
	product.Stock = productDTO.Stock
	product.ImageURL = productDTO.ImageURL

	category := database.GetCategoryByName(productDTO.CategoryName)
	if  category != nil {
		product.CategoryID = &category.ID
		product.Category = category
	}

	if database.AddProduct(product) != 0 {
		// loguear
		return 1
	}
	return 0
}