package controllers

import (
	"github.com/gin-gonic/gin"
	"golang-gin-jwt-mysql-restful-api/models/post"
	"golang-gin-jwt-mysql-restful-api/db"
	"github.com/appleboy/gin-jwt"
	"strings"
	"strconv"
	"fmt"
)

func CreatePost(c *gin.Context) {
	var post post.Post
	c.BindJSON(&post)
	var getDb = db.GetDb()
	// jwt user info as string interface
	claims := jwt.ExtractClaims(c)
	// separate string by coma(,)
	userInfo:= strings.Split(claims["id"].(string), ",")
	// first index is username
	userName := userInfo[0]
	// second index is user ID
	userId := userInfo[1]
	// parse string to Uint8
	id ,_  := strconv.ParseUint(userId, 10, 8)
	// parse uint as user ID format and assign UserId as post data
	post.UserId = uint(id)
	// assign userName as post data
	post.UserName = userName
	if err := getDb.Model(&post).Related(&post.Author, "user_id").Error; err != nil {
		fmt.Println(err)
		c.AbortWithStatus(404)
	}
	post.Author.Password=""
	c.JSON(200, post)
}

func GetPosts(c *gin.Context) {
	var posts []post.Post

	var getDb = db.GetDb()
	//var author models.User
	if err := getDb.Find(&posts).Error; err != nil {
		fmt.Println(err)
		c.AbortWithStatus(404)
	}

	for i, _ := range posts {
		getDb.Model(posts[i]).Related(&posts[i].Author)
		// unset password from public api
		posts[i].Author.Password = ""
	}
	c.JSON(200, posts)
}

func GetMyPosts(c *gin.Context)  {
	claims := jwt.ExtractClaims(c)
	userInfo:= strings.Split(claims["id"].(string), ",")
	//userName := userInfo[0]
	userId := userInfo[1]
	id ,_  := strconv.ParseUint(userId, 10, 8)

	var getDb = db.GetDb()
	var posts []post.Post

	if err := getDb.Where("user_id=?",uint(id)).Find(&posts).Error; err != nil {
		fmt.Println(err)
		c.AbortWithStatus(404)
	}
	for i, _ := range posts {
		getDb.Model(posts[i]).Related(&posts[i].Author)
		// unset password from public api
		posts[i].Author.Password = ""
	}
	c.JSON(200, posts)
}

func GetPost(c *gin.Context) {
	postId := c.Param("id")
	var getDb = db.GetDb()
	var post post.Post

	if err :=getDb.Where(postId).Find(&post).Error; err != nil {
		fmt.Println(err)
		c.AbortWithStatus(404)
	}
	if err :=getDb.Model(&post).Related(&post.Author).Error; err != nil{
		fmt.Println(err)
		c.AbortWithStatus(404)
	}
	c.JSON(200, post)
}

func UpdatePost(c *gin.Context) {
	var post post.Post
	c.BindJSON(&post)

	var getDb = db.GetDb()
	if err := getDb.Model(&post).Update(&post).Error; err != nil {
		fmt.Println(err)
		c.AbortWithStatus(404)
	}

	c.JSON(200, post)
}

func DeletePost(c *gin.Context) {
	var post post.Post
	c.BindJSON(&post)
	var getDb = db.GetDb()
	if err :=getDb.Where(post.ID).Delete(&post).Error; err != nil{
		c.AbortWithStatus(404)
	}
	c.JSON(200, post)
}