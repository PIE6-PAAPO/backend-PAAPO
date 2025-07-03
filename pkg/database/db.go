// db.go
package database

import (
	"fmt"
	"os"

	healthinfo "backend-PAAPO/internal/HealthInformation"
	medicaldata "backend-PAAPO/internal/MedicalData"
	"backend-PAAPO/internal/models"
	"backend-PAAPO/internal/users"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	// Set defaults if not provided
	if host == "" {
		host = "localhost"
	}
	if port == "" {
		port = "5432"
	}
	if user == "" {
		user = "postgres"
	}
	if password == "" {
		password = "postgres"
	}
	if dbname == "" {
		dbname = "paapo"
	}

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	// Migrate all models
	if err := db.AutoMigrate(
		&users.User{},
		&healthinfo.HealthInformation{},
		&medicaldata.MedicalData{},
		&models.TrainingSession{},
	); err != nil {
		return nil, fmt.Errorf("migration failed: %w", err)
	}

	fmt.Println("✅ Connected to database with GORM")
	return db, nil
}
