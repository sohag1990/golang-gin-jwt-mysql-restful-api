package main

import (
	"os"
	"github.com/gin-gonic/gin"
	"net/http"
	"golang-gin-jwt-mysql-restful-api/controllers"
	"golang-gin-jwt-mysql-restful-api/db"
	"golang-gin-jwt-mysql-restful-api/middlewares"
)

func main() {
	db.Init()
	port := os.Getenv("PORT")
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middlewares.CORSMiddlewareHandler())

	if port == "" {
		port = "4200"
	}

	//jwt Middleware
	authMiddleware := middlewares.GinJwtMiddlewareHandler()
	// public api calls
	r.POST("/api/v1/user/login", authMiddleware.LoginHandler)
	r.POST("/api/v1/user/", controllers.CreateUser)

	//restricted api calls
	auth := r.Group("/api/v1/")
	auth.Use(authMiddleware.MiddlewareFunc())
	{
		auth.GET("/users/", controllers.GetUsers)
		auth.GET("/user/:id", controllers.GetUser)

		auth.PUT("/user/", controllers.UpdateUser)
		auth.DELETE("/user/", controllers.DeleteUser)
	}

	http.ListenAndServe(":"+port, r)

	defer db.CloseDb() // close mysql connection
}