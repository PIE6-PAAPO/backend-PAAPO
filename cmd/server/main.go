package main

import (
	"log"
	"os"

	"time"

	healthinfo "backend-PAAPO/internal/HealthInformation"
	medicaldata "backend-PAAPO/internal/MedicalData"
	trainingsession "backend-PAAPO/internal/TrainingSession"
	"backend-PAAPO/internal/users"
	"backend-PAAPO/pkg/database"
	"backend-PAAPO/routes/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found, using system environment variables")
	}

	// Check for required environment variables
	requiredVars := []string{"JWT_SECRET", "DB_HOST", "DB_USER", "DB_PASSWORD", "DB_NAME"}
	for _, v := range requiredVars {
		if os.Getenv(v) == "" {
			log.Fatalf("Required environment variable %s is not set", v)
		}
	}

	// Log JWT secret info (don't log the actual secret in production)
	jwtSecret := os.Getenv("JWT_SECRET")
	log.Printf("JWT secret length: %d characters", len(jwtSecret))
	if len(jwtSecret) < 32 {
		log.Println("WARNING: JWT secret is shorter than recommended 32 characters")
	}

	db, err := database.ConnectDB()
	if err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}

	// Inicializa repositórios e serviços
	userRepository := users.NewUserRepository(db)
	userService := users.NewService(userRepository)

	healthInfoRepo := healthinfo.NewHealthInformationRepository(db)
	healthInfoService := healthinfo.NewService(healthInfoRepo)
	healthInfoHandler := healthinfo.NewHandler(healthInfoService)

	medicalDataRepo := medicaldata.NewMedicalDataRepository(db)
	medicalDataService := medicaldata.NewService(medicalDataRepo)
	medicalDataHandler := medicaldata.NewHandler(medicalDataService)

	// Training Session module
	trainingSessionRepo := trainingsession.NewTrainingSessionRepository(db)
	trainingSessionService := trainingsession.NewService(trainingSessionRepo)
	trainingSessionHandler := trainingsession.NewHandler(trainingSessionService)

	// Cria o router Gin com configurações padrão
	router := gin.Default()

	// Configure CORS middleware
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// API v1 routes
	api := router.Group("/api/v1")

	// Public routes (no auth required)
	authGroup := api.Group("/auth")
	users.UserRoutes(authGroup, userService)

	// Protected routes (require JWT auth)
	protected := api.Group("")
	protected.Use(middleware.AuthMiddleware())
	{
		// Health Information routes
		healthInfoHandler.RegisterRoutes(protected)

		// Medical Data routes
		medicalDataHandler.RegisterRoutes(protected)

		// Training Session routes
		trainingSessionHandler.RegisterRoutes(protected)
	}

	// Inicia o servidor na porta 8080
	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
}
