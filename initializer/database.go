package initializer

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() *gorm.DB {

	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	if dbUser == "" {
		dbUser = "''"
	}

	if dbPassword == "" {
		dbPassword = ""
	}

	if dbHost == "" {
		dbHost = "localhost"
	}

	if dbPort == "" {
		dbPort = "5432"
	}

	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", dbHost, dbUser, dbPassword, dbName, dbPort)

	log.Println("Connecting to database...")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	log.Println("Done...")

	if err != nil {
		log.Fatal("Connecting to db failed", err)
	}

	return DB
}
