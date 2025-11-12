package service

import "go-eCommerce/internal/database"

func AddProductToCart(userID uint, productID uint, quantity int) error {

	cartID, err := database.GetCartID(userID)
	if err != nil{
		return err
	}
	print(cartID)
	
	return nil
}