package main

import (
	"os"
	"github.com/gin-gonic/gin"
	"net/http"
	"golang-gin-jwt-mysql-restful-api/controllers"
)

func main() {
	port := os.Getenv("PORT")
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	if port == "" {
		port = "8000"
	}

	r.GET("/", controllers.HelloWorld)

	http.ListenAndServe(":"+port, r)
}