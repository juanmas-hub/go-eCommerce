package models

import (
	"time"
)

type Category struct {
	ID          uint      `gorm:"primaryKey"`
	Name        string    `gorm:"size:100;unique;not null"`
	Description string    `gorm:"type:text"`
	CreatedAt   time.Time
	UpdatedAt   time.Time

	Products []Product `gorm:"foreignKey:CategoryID"`
}