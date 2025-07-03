package trainingsession

import (
	"backend-PAAPO/internal/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) RegisterRoutes(router *gin.RouterGroup) {
	trainingGroup := router.Group("/training-sessions")
	{
		trainingGroup.POST("", h.CreateTrainingSession)
		trainingGroup.PUT("/end", h.EndTrainingSession)
		trainingGroup.GET("", h.GetAllTrainingSessions)
		trainingGroup.GET("/active", h.GetActiveTrainingSession)
	}
}

// CreateTrainingSession creates a new training session
func (h *Handler) CreateTrainingSession(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not authenticated"})
		return
	}

	var createDTO dto.CreateTrainingSessionDTO
	if err := c.ShouldBindJSON(&createDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "invalid request body",
			"details": err.Error(),
		})
		return
	}

	session, err := h.service.CreateTrainingSession(userID.(string), createDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, session)
}

// EndTrainingSession ends the current active training session
func (h *Handler) EndTrainingSession(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not authenticated"})
		return
	}

	var endDTO dto.EndTrainingSessionDTO
	if err := c.ShouldBindJSON(&endDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "invalid request body",
			"details": err.Error(),
		})
		return
	}

	session, err := h.service.EndTrainingSession(userID.(string), endDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, session)
}

// GetAllTrainingSessions retrieves all training sessions for the user
func (h *Handler) GetAllTrainingSessions(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not authenticated"})
		return
	}

	sessions, err := h.service.GetAllTrainingSessions(userID.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to retrieve training sessions",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"sessions": sessions,
		"count":    len(sessions),
	})
}

// GetActiveTrainingSession retrieves the current active training session
func (h *Handler) GetActiveTrainingSession(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not authenticated"})
		return
	}

	session, err := h.service.GetActiveTrainingSession(userID.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to retrieve active training session",
		})
		return
	}

	if session == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "no active training session found",
		})
		return
	}

	c.JSON(http.StatusOK, session)
}
