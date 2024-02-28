package controller

import (
	"example/go-crud/initializers"
	"example/go-crud/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostCreate(c *gin.Context) {
	var body struct {
		Body string
		Title string
	}

	c.Bind(&body)


	post := models.Post{Title: body.Title, Body: body.Body}

	result := initializers.DB.Create(&post) // pass pointer of data to Create

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"message": post,
	})
}

func PostIndex(c *gin.Context) {

	var posts []models.Post
	initializers.DB.Find(&posts)

	c.IndentedJSON(http.StatusOK, gin.H{
		"posts":posts,
	})
}

func GetOnePost(c *gin.Context) {
	var post models.Post
	id := c.Param("id")

	initializers.DB.Find(&post, id)
	c.IndentedJSON(http.StatusOK, gin.H{
		"post":post,
	})
}

func PostsUpdate(c *gin.Context) {
	id := c.Param("id")
	
	var body struct{
		Body string
		Title string
	}

	c.Bind(&body)
	var post models.Post

	//find post
	initializers.DB.First(&post, id)

	//update post
	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body: body.Body,
	})

	c.JSON(http.StatusOK, gin.H{"post":post})
}

func PostDelete(c *gin.Context) {
	id := c.Param("id")

	var post models.Post
	initializers.DB.Delete(&post, id)

	c.IndentedJSON(http.StatusOK, gin.H{
		"message" : "delete successful",
	})
}