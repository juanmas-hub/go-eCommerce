package database

import (
	"go-eCommerce/initializers"
	"go-eCommerce/internal/models"
)

func GetCategoryByName(name string) *models.Category{
	var category *models.Category

	if initializers.DB.Where("name = ?", name).First(category).Error != nil {
		//loguear
		return nil
	}

	return category
}