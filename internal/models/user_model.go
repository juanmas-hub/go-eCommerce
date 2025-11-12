package models

import "gorm.io/gorm"

type User struct {
    gorm.Model
    Username string `json:"username"`
    Email    string `json:"email"`
    Password string `json:"password" gorm:"unique"`
    Role     string `json:"role" gorm:"size:20;default:customer"`

    Orders   []Order `gorm:"foreignKey:UserID"`
	Cart     *Cart    `gorm:"foreignKey:UserID"`
}
