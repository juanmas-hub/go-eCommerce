package dto

type ProductDTO struct {
	Name         string  `json:"name"`
	Price        float64 `json:"price"`
	Stock        int 	 `json:"stock"`
	ImageURL     string	 `json:"image"`
	CategoryName string  `json:"category"`
}