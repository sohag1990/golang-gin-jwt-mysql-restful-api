package post

import (
	"github.com/jinzhu/gorm"
	"golang-gin-jwt-mysql-restful-api/models"
)

type Post struct {
	gorm.Model
	Type string
	Title string
	Content string
	Excerpt string
	Author models.User
	UserId uint
}


//Categories []int