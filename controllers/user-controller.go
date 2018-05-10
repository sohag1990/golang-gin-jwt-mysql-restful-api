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

	var getDb = db.GetDb()
	getDb.Create(&user)
	c.JSON(200, user)
}
func UpdateUser(c *gin.Context) {
	var user models.User

	c.BindJSON(&user)

	var getDb = db.GetDb()
	getDb.Model(&user).Updates(&user)
	c.JSON(200, user)
}

func DeleteUser(c *gin.Context) {
	var user models.User

	c.BindJSON(&user)

	var getDb = db.GetDb()
	getDb.Delete(&user)
	getDb.Delete(&user.Profile)
	getDb.Delete(&user.Profile.Address)
	c.JSON(200, user)
}
func GetUser(c *gin.Context) {
	userID :=c.Params.ByName("id")
	var user models.User
	var getDb = db.GetDb()

	if err := getDb.Model(&user).Where(userID).Preload("Profile").Preload("Profile.Address").First(&user).Error; err != nil {
		fmt.Println(err)
		c.AbortWithStatus(404)
	}
	c.JSON(200, user)
}
