package config

import (
	"fmt"
	"os"
	"rest-example/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	connectionString := os.Getenv("MYSQL_CONNECTION_STRING")
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
