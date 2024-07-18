package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"project-managment/internal/app/config"
	"project-managment/internal/app/models"
)

var DB *gorm.DB

func InitDB(config *config.Config) (*gorm.DB, error) {
	dsn := config.DBConnectionString()

	var err error

	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	err = DB.AutoMigrate(&models.User{}, &models.Task{}, &models.Project{})

	if err != nil {
		log.Fatalf("failed to migrate database %v", err)
	}

	return DB, nil

}
