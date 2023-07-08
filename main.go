package main

import (
	"fmt"

	"github.com/abeloha/go-crud/controller"
	"github.com/abeloha/go-crud/initializer"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	initializer.LoadEnv()
	initializer.ConnectToDB()
}
func main() {
	fmt.Println("Hello, World!")

	r := gin.Default()

	r.GET("/", controller.PostIndex)
	r.POST("/", controller.PostCreate)
	r.Run() // listen and serve on 0.0.0.0:8080
}
