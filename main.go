package main

import (
	"github.com/gin-gonic/gin"

	"github.com/rizqonmaulana/social-post-service/controllers"
	"github.com/rizqonmaulana/social-post-service/models"
)

func main() {
	r := gin.Default()

//   r.GET("/", func(c *gin.Context) {
//     c.JSON(http.StatusOK, gin.H{"data": "hello world"})    
//   })
	r.GET("/posts", controllers.FindPosts)

	models.ConnectDatabase()

	r.Run()
}