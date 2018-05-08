package models

import "github.com/jinzhu/gorm"

type Address struct {
	gorm.Model
	State string
	City string
	Zip string
	Country string
	UserID uint
}
