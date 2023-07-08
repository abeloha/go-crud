package controller

import (
	"log"

	"github.com/abeloha/go-crud/initializer"
	"github.com/abeloha/go-crud/model"
	"github.com/gin-gonic/gin"
)

func PostIndex(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func PostCreate(c *gin.Context) {

	// get input

	// store input
	post := &model.Post{Title: "My Post", Body: "Post body details here"}

	result := initializer.DB.Create(&post) // pass pointer of data to Create

	log.Println("result.Error", result.Error)
	log.Println("result.RowsAffected", result.RowsAffected)
	// return it
	c.JSON(201, gin.H{
		"message": "Post created",
		"data":    post,
	})
}
