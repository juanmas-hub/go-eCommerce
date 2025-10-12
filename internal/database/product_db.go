package database

import (
	"go-eCommerce/initializers"
	"go-eCommerce/internal/models"
)

func GetProducts() []string {
	var names []string

	result := initializers.DB.
		Model(&models.Product{}).
		Select("name").
		Scan(&names)

	if result.Error != nil {
		// loguear
		return nil
	}

	return names
}

func AddProduct(product models.Product) int {
    result := initializers.DB.Create(&product)

    if result.Error != nil {
        return 1
    }

    return 0
}