package models

import "gorm.io/gorm"

type User struct {
    gorm.Model
    Username string `json:"username" gorm:"column:username"`
    Email    string `json:"email" gorm:"column:email"`
    Password string `json:"password" gorm:"column:password;unique"`
    Role     string `gorm:"size:20;default:customer"`

    Orders   []Order `gorm:"foreignKey:UserID"`
	Cart     *Cart    `gorm:"foreignKey:UserID"`
}
