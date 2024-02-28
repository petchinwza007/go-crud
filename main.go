package main

import (
	"example/go-crud/controller"
	"example/go-crud/initializers"
	"example/go-crud/middleware"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnv()
  	initializers.ConnectDB()
	initializers.MigrateDB()
}

func main() {
  r := gin.Default()
  callpost := r.Group("/callpost")
  {
	callpost.POST("/posts", controller.PostCreate)
	callpost.PATCH("/post/:id",controller.PostsUpdate)
	callpost.GET("/posts", controller.PostIndex)
	callpost.GET("/post/:id", controller.GetOnePost)
	callpost.DELETE("/post/:id", controller.PostDelete)
  }

  calluser := r.Group("/calluser")
  {
	calluser.POST("/signup", controller.Signup)
	calluser.POST("/login", controller.Login)
	calluser.GET("/validate", middleware.RequireAuth ,controller.Validate)

  }


  r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}