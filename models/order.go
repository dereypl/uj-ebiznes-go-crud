package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	UserID   uint
	Products []Product `gorm:"many2many:order_product;"`
}
