package middlewares

import (
	"github.com/gin-gonic/gin"
	"time"
	"github.com/appleboy/gin-jwt"
	"fmt"
	"golang-gin-jwt-mysql-restful-api/db"
	"golang-gin-jwt-mysql-restful-api/models"
)

func GinJwtMiddlewareHandler() *jwt.GinJWTMiddleware {
	return &jwt.GinJWTMiddleware{
		Realm:      "test zone",
		Key:        []byte("&*3kasdJKH#$sjksQq774kakakl*(&(8320_))92@#$^$%&^%^*&@#kljlkjLLLLAskaiw#S"),
		Timeout:    time.Hour * 100,
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
		Authorizator: func(userId string, c *gin.Context) bool {
			//if userId == "sohag" {
			//	return true
			//}
			//return false

			return true
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		TokenLookup: "header:Authorization",
		TokenHeadName: "Bearer",
		TimeFunc: time.Now,
	}
}