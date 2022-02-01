package config

import (
	"fmt"
	"log"
	"os"
	"rest-example/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	connectionString := os.Getenv("MYSQL_CONNECTION_STRING")
	fmt.Println("connectionString", connectionString)
	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	} else {
		fmt.Println("DATABASE is CONNECTED")
	}
	InitialMigration(db)
	return db
}

func InitialMigration(db *gorm.DB) {
	db.AutoMigrate(&models.User{})
}
