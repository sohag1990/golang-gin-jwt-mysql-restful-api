package middlewares

import (
	"github.com/appleboy/gin-jwt"
	"time"
	"github.com/gin-gonic/gin"
	"fmt"
	"golang-gin-jwt-mysql-restful-api/db"
	"golang-gin-jwt-mysql-restful-api/models"
)

func GinJWTMiddlewareHandler() *jwt.GinJWTMiddleware {
	return &jwt.GinJWTMiddleware{
		Realm: "Test Zone",
		Key: []byte("skjdlfjs@#@#!@!kjklj"),
		Timeout: time.Hour,
		MaxRefresh: time.Hour,
		Authenticator: func(userId string, password string, c *gin.Context) (string, bool) {
			var getDB = db.GetDb()
			var user models.User
			if err := getDB.Where("username=? AND password=?", userId, password).Find(&user).Limit(1).Error; err != nil {
				c.AbortWithStatus(404)
				fmt.Println(err)
			} else {
				fmt.Println(user)
				//c.JSON(200, user)
				return userId, true
			}

			return userId, false
		},
		Authorizator: func(userID string, c *gin.Context) bool {
			return true
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code": code,
				"message": message,
			})
		},
		TokenLookup: "Header:Authentication",
		TokenHeadName: "Bearer",
		TimeFunc: time.Now,
	}
}
