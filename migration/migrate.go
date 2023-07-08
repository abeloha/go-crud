package main

import (
	"log"

	initializer "github.com/abeloha/go-crud/initializer"
	model "github.com/abeloha/go-crud/model"
)

func init() {
	initializer.LoadEnv()
	initializer.ConnectToDB()
}
func migration() {
	log.Println("Migrating...")
	err := initializer.DB.AutoMigrate(&model.Post{})

	if err != nil {
		log.Fatal("Error with migration ", err)
	}

	log.Println("Migration Done...")
}
