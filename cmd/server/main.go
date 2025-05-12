package main

import (
	"backend-PAAPO/pkg/database"
	"log"

	"backend-PAAPO/internal/users"

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

	router := gin.Default()

	users.UserRoutes(router, userService)

	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
}
