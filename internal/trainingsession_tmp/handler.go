package trainingsession

import (
	"backend-PAAPO/internal/dto"
	"net/http"
	"strconv"
	"time"

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
	// Alias for frontend convenience
	router.GET("/session/active", h.GetActiveTrainingSession)
	// Reporting endpoints
	router.GET("/sessions", h.GetAllTrainingSessions)
	router.GET("/sessions/total-hours", h.GetTotalHoursTrained)
	router.GET("/sessions/average-duration", h.GetAverageSessionDuration)
	router.GET("/sessions/category-count", h.GetSessionCountPerCategory)
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
		if conflict, ok := err.(*ConflictError); ok {
			c.JSON(http.StatusConflict, gin.H{
				"error":   conflict.Error(),
				"session": conflict.Session,
			})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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

	// Filtering and pagination
	var filters SessionFilters
	if fromStr := c.Query("from"); fromStr != "" {
		if from, err := time.Parse(time.RFC3339, fromStr); err == nil {
			filters.From = &from
		}
	}
	if toStr := c.Query("to"); toStr != "" {
		if to, err := time.Parse(time.RFC3339, toStr); err == nil {
			filters.To = &to
		}
	}
	if cat := c.Query("category"); cat != "" {
		filters.Category = &cat
	}
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	sessions, count, err := h.service.GetAllTrainingSessions(userID.(string), filters, page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to retrieve training sessions"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"sessions": sessions,
		"count":    count,
		"page":     page,
		"limit":    limit,
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to retrieve active training session"})
		return
	}

	if session == nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "no active training session found"})
		return
	}

	c.JSON(http.StatusOK, session)
}

// Reporting endpoints
func (h *Handler) GetTotalHoursTrained(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not authenticated"})
		return
	}
	fromStr := c.Query("from")
	toStr := c.Query("to")
	if fromStr == "" || toStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "from and to are required (RFC3339)"})
		return
	}
	from, err1 := time.Parse(time.RFC3339, fromStr)
	to, err2 := time.Parse(time.RFC3339, toStr)
	if err1 != nil || err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid date format (RFC3339)"})
		return
	}
	hours, err := h.service.TotalHoursTrained(userID.(string), from, to)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"total_hours": hours})
}

func (h *Handler) GetAverageSessionDuration(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not authenticated"})
		return
	}
	fromStr := c.Query("from")
	toStr := c.Query("to")
	if fromStr == "" || toStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "from and to are required (RFC3339)"})
		return
	}
	from, err1 := time.Parse(time.RFC3339, fromStr)
	to, err2 := time.Parse(time.RFC3339, toStr)
	if err1 != nil || err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid date format (RFC3339)"})
		return
	}
	avg, err := h.service.AverageSessionDuration(userID.(string), from, to)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"average_duration_minutes": avg})
}

func (h *Handler) GetSessionCountPerCategory(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not authenticated"})
		return
	}
	counts, err := h.service.SessionCountPerCategory(userID.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"category_counts": counts})
}
