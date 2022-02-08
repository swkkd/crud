package database

import (
	"fmt"
	"github.com/swkkd/crud/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func Connection(dsn string) {
	//Database connection
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	err = db.AutoMigrate(&models.Customer{})
	if err != nil {
		log.Fatalf("AutoMigrate failed with error: %v", err)
	}
	fmt.Println("Migration Successful!")
	DB = db
}
func GetDB() *gorm.DB {
	return DB
}
