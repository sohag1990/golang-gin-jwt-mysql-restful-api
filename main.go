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

	auth := r.Group("/api/v1/")
	// public api calls
	auth.POST("/user/login", authMiddleware.LoginHandler)
	auth.POST("/user/", controllers.CreateUser)
	auth.GET("/posts", controllers.GetPosts)
	//restricted api calls
	auth.Use(authMiddleware.MiddlewareFunc())
	{
		auth.GET("/users/", controllers.GetUsers)
		auth.GET("/user/:id", controllers.GetUser)

		auth.PUT("/user/", controllers.UpdateUser)
		auth.DELETE("/user/", controllers.DeleteUser)

		auth.POST("/post/", controllers.CreatePost)
		auth.GET("/posts/my-posts", controllers.GetMyPosts)
		auth.GET("/post/:id", controllers.GetPost)
		auth.PUT("/post/", controllers.UpdatePost)
		auth.DELETE("/post/", controllers.DeletePost)
	}

	http.ListenAndServe(":"+port, r)

	defer db.CloseDb() // close mysql connection
}