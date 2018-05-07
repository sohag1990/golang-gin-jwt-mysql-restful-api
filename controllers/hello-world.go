package controllers

import (
	"github.com/gin-gonic/gin"
	"golang-gin-jwt-mysql-restful-api/models"
)

func HelloWorld(c *gin.Context)  {
	var hello models.Hello
	hello.ID = 1
	hello.Hello = "Hello world"
	c.JSON(200, hello)
}