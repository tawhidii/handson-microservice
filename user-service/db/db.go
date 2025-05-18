package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB

func ConnectDatabase() {
	dbDsn := os.Getenv("DATABASE_URL")
	fmt.Println("DB URL is", dbDsn)
	if dbDsn == "" {
		log.Fatal("DATABASE_URL is not set in the environment")
	}

	var err error
	DB, err = gorm.Open(postgres.Open(dbDsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	fmt.Println("✅ Successfully connected to the database.")
}
