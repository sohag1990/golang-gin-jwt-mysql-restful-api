package main

import (
	"os"
	"github.com/gin-gonic/gin"
	"net/http"
	"golang-gin-jwt-mysql-restful-api/controllers"
	"golang-gin-jwt-mysql-restful-api/db"
)

func main() {
	db.Init()
	port := os.Getenv("PORT")
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	if port == "" {
		port = "4200"
	}

	r.GET("/", controllers.GetUsers)
	r.POST("/user/", controllers.CreateUser)

	http.ListenAndServe(":"+port, r)

	defer db.CloseDb() // close mysql connection
}