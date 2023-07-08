package main

import (
	"fmt"

	"github.com/abeloha/go-crud/initializer"
	"github.com/gin-gonic/gin"
)

func init() {
	initializer.LoadEnv()
	initializer.ConnectToDB()
}
func main() {
	fmt.Println("Hello, World!")

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
