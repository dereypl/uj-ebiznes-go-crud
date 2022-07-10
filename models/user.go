package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email      string
	FirstName  string
	LastName   string
	CreditCard CreditCard
	Orders     []Order
}
