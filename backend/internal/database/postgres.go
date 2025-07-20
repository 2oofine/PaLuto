package database

import (
	"log"
	"os"

	"github.com/2oofine/PaLuto/backend/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	dsn := os.Getenv("DATABASE_URL")
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to the database: ", err)
	}

	// ✅ Auto-migrate tables
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("Failed to run migrations:", err)
	}

	DB = db
	log.Println("Database connection successful")
}
