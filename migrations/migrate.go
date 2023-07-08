package main

import (
	"log"

	initializer "github.com/abeloha/go-crud/initializers"
	model "github.com/abeloha/go-crud/models"
)

func init() {
	initializer.LoadEnv()
	initializer.ConnectToDB()
}
func main() {
	log.Println("Migrating...")
	err := initializer.DB.AutoMigrate(&model.Post{})

	if err != nil {
		log.Fatal("Error with migration ", err)
	}

	log.Println("Migration Done...")
}
