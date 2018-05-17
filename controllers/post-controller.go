package controllers

import (
	"github.com/gin-gonic/gin"
	"golang-gin-jwt-mysql-restful-api/models/post"
	"github.com/appleboy/gin-jwt"
	"fmt"
)

func CreatePost(c *gin.Context) {
	var post post.Post
	c.Bind(post)
	//var getDb = db.GetDb()
	//getDb.Create(&post)
	claims := jwt.ExtractClaims(c)
	userInfo:= claims["id"]
	fmt.Println(userInfo)
	c.JSON(200, post)
}
