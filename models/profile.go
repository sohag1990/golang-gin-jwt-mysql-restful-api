package models

import "github.com/jinzhu/gorm"

type Profile struct {
	gorm.Model
	FirstName string
	LastName string
	Address Address
	UserID uint
}
