package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"go-tasks-app-practice/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB

func InitDB() (*gorm.DB, error) {

	err := godotenv.Load()
	if err != nil {
		log.Printf("No .env file found, fallback to system environment variables")
	}

	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	DataSourceName := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, port)
	db, err := gorm.Open(postgres.Open(DataSourceName), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// auto migrate the models
	err = db.AutoMigrate(&models.User{}, &models.Task{})
	if err != nil {
		return nil, err
	}

	DB = db
	return DB, nil
}
