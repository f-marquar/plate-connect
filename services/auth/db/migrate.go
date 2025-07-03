package db

import (
	"log"

	"plate-connect/services/auth/models"
)

func Migrate() {

	err := DB.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`).Error
	if err != nil {
		log.Fatalf("Error creating uuid-ossp extension: %v", err)
	}

	err = DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatalf("Migration not successfull: %v", err)
	}
	log.Println("Migration successfull")
}
