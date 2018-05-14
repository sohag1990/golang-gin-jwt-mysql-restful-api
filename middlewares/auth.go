package middlewares

import (
	"github.com/appleboy/gin-jwt"
	"time"
	"github.com/gin-gonic/gin"
)

func GinJWTMiddlewareHandler() *jwt.GinJWTMiddleware {
	return &jwt.GinJWTMiddleware {
		Realm: "Test Zone",
		Key: []byte(""),
		Timeout: time.Hour,
		MaxRefresh: time.Hour,

		Authenticator: func(userID string, password string, c *gin.Context) (string, bool) {
			return userID, true
		},
		Authorizator: func(userID string, c *gin.Context) bool {
			return true
		},
		Unauthorized: func(c *gin.Context, i int, s string) {

		},
		TokenLookup: "Header:Authorization",
		TokenHeadName: "Bearer",
		TimeFunc:time.Now,
	}
}
