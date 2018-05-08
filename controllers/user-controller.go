package controllers

import (
	"github.com/gin-gonic/gin"
	"golang-gin-jwt-mysql-restful-api/models"
	"golang-gin-jwt-mysql-restful-api/db"
	"fmt"
)

func GetUsers(c *gin.Context)  {
	var users []models.User
	var getDb = db.GetDb()
	if err := getDb.Preload("Profile").Preload("Profile.Address").Find(&users).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}

	c.JSON(200, users)
}

func CreateUser(c *gin.Context) {
	var user models.User

	c.BindJSON(&user)

	fmt.Println(user)
	c.JSON(200, user)
}
func UpdateUser(c *gin.Context) {

}

func DeleteUser(c *gin.Context) {

}
func GetUser(c *gin.Context) {

}
