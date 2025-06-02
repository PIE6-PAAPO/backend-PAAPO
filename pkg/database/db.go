// db.go
package database

import (
	"fmt"
	"os"

	healthinfo "backend-PAAPO/internal/HealthInformation"
	medicaldata "backend-PAAPO/internal/MedicalData"
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

	if host == "" {
		host = "localhost"
		port = "5432"
		user = "postgres"
		password = "senha"
		dbname = "meubanco"
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
	); err != nil {
		return nil, fmt.Errorf("migration failed: %w", err)
	}

	fmt.Println("✅ Connected to database with GORM")
	return db, nil
}
