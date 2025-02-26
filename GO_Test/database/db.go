package database

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"main.go/models"
)

var DB *gorm.DB

func InitDB() {
	var err error
	// Change this DSN as per your MySQL setup
	dsn := "root:Kush@789#@tcp(dev.wikibedtimestories.com:31347)/testdb?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to the database", err)
	}

	// AutoMigrate creates the users table if it doesn't exist
	DB.AutoMigrate(&models.User{})
}
