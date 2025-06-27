package main

import (
	"log"
	"time"

	healthinfo "backend-PAAPO/internal/HealthInformation"
	medicaldata "backend-PAAPO/internal/MedicalData"
	"backend-PAAPO/internal/users"
	"backend-PAAPO/pkg/database"

	"github.com/gin-contrib/cors"
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

	// Inicializa repositórios e serviços
	userRepository := users.NewUserRepository(db)
	userService := users.NewService(userRepository)

	healthInfoRepo := healthinfo.NewHealthInformationRepository(db)
	healthInfoService := healthinfo.NewService(healthInfoRepo)
	healthInfoHandler := healthinfo.NewHandler(healthInfoService)

	medicalDataRepo := medicaldata.NewMedicalDataRepository(db)
	medicalDataService := medicaldata.NewService(medicalDataRepo)
	medicalDataHandler := medicaldata.NewHandler(medicalDataService)

	// Cria o router Gin com configurações padrão
	router := gin.Default()

	// Middleware de CORS (dev: liberar todas origens)
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Agrupa as rotas na versão 1 da API
	api := router.Group("/api/v1")
	{
		// Rotas de autenticação e cadastro
		users.UserRoutes(api, userService)

		// Rotas de Health Information
	healthInfoHandler.RegisterRoutes(api)

		// Rotas de Medical Data
		medicalDataHandler.RegisterRoutes(api)
	}

	// Inicia o servidor na porta 8080
	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
}
