package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"fmt"
	"golang-gin-jwt-mysql-restful-api/models"
	"golang-gin-jwt-mysql-restful-api/models/post"
)


var db *gorm.DB
var err error

func Init() {
	db, err = gorm.Open("mysql", "root:111111@tcp(172.17.0.2:3306)/gin_jwt?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
	}

	//db.DropTable(&models.User{})
	//db.DropTable(&models.Profile{})
	//db.DropTable(&models.Address{})


	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Profile{})
	db.AutoMigrate(&models.Address{})
	db.AutoMigrate(&post.Post{})

}
func GetDb() *gorm.DB {
	return db
}
func CloseDb() {
	db.Close()
}
