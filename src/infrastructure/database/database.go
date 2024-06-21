package database

import (
	"github.com/JP-Cano/sports-management-app/src/core/entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func New(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("Error connecting to database: %v", err.Error())
		return nil, err
	}
	log.Println("Database connection established")
	return db, nil
}

func Close(db *gorm.DB) {
	conn, err := db.DB()
	if err != nil {
		log.Printf("Error retrieving DB connection: %v", err.Error())
		return
	}

	err = conn.Close()
	if err != nil {
		log.Printf("Error closing DB connection: %v", err.Error())
		return
	}

	log.Println("Database connection closed successfully")
	return
}

func Migrate(db *gorm.DB) {
	createUUIDExtension(db)
	err := db.AutoMigrate(entities.User{})
	if err != nil {
		log.Printf("Error auto migrating entities: %v", err.Error())
		return
	}
}

func createUUIDExtension(db *gorm.DB) {
	if err := db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp"`).Error; err != nil {
		log.Fatalf("Failed to create uuid-ossp extension: %v", err.Error())
		return
	}
}
