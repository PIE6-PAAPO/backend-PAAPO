package users

import (
	"backend-PAAPO/internal/dto"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine, service *Service) {
	r.POST("/v1/auth/login", func(c *gin.Context) {
		var loginDTO dto.LoginDTO
		if err := c.ShouldBindJSON(&loginDTO); err != nil {
			c.JSON(400, gin.H{"error": "Invalid input"})
			return
		}
		user, err := service.Login(loginDTO)
		if err != nil {
			c.JSON(401, gin.H{"error": "Invalid credentials"})
			return
		}

		sessionID, err := service.CreateSession(user.ID)
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to create session"})
			return
		}

		accessToken, err := GenerateJWT(user.ID, user.Role)
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to generate token"})
			return
		}

		c.JSON(200, gin.H{
			"access_token": accessToken,
			"session_id":   sessionID,
			"expires_in":   900,
		})
	})

	r.POST("/v1/auth/register", func(c *gin.Context) {
		var registerDTO dto.RegisterDTO
		if err := c.ShouldBindJSON(&registerDTO); err != nil {
			c.JSON(400, gin.H{"error": "Invalid input"})
			return
		}
		newUser, err := service.Register(registerDTO)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		sessionID, err := service.CreateSession(newUser.ID)
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to create session"})
			return
		}

		accessToken, err := GenerateJWT(newUser.ID, newUser.Role)
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to generate token"})
			return
		}

		c.JSON(200, gin.H{
			"access_token": accessToken,
			"session_id":   sessionID,
			"expires_in":   900,
		})
	})
}
