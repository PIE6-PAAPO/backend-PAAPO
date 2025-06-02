package main

import (
	"log"

	healthinfo "backend-PAAPO/internal/HealthInformation"
	medicaldata "backend-PAAPO/internal/MedicalData"
	"backend-PAAPO/internal/users"
	"backend-PAAPO/pkg/database"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println(".env file not found, using defaults")
	}

	db, err := database.ConnectDB()
	if err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}

	userRepository := users.NewUserRepository(db)
	userService := users.NewService(userRepository)

	healthInfoRepo := healthinfo.NewHealthInformationRepository(db)
	healthInfoService := healthinfo.NewService(healthInfoRepo)
	healthInfoHandler := healthinfo.NewHandler(healthInfoService)

	medicalDataRepo := medicaldata.NewMedicalDataRepository(db)
	medicalDataService := medicaldata.NewService(medicalDataRepo)
	medicalDataHandler := medicaldata.NewHandler(medicalDataService)

	router := gin.Default()

	api := router.Group("/api/v1")
	{
		// User authentication routes
		users.UserRoutes(api, userService)

		// Health Information routes
		healthInfoHandler.RegisterRoutes(api)

		// Medical Data routes
		medicalDataHandler.RegisterRoutes(api)
	}

	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
}
