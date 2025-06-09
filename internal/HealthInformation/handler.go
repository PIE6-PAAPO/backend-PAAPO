package healthinformation

import (
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
	healthInfoGroup := router.Group("/health-info")
	{
		healthInfoGroup.POST("", h.CreateHealthInfo)
		healthInfoGroup.GET("", h.GetHealthInfo)
		healthInfoGroup.PUT("", h.UpdateHealthInfo)
		healthInfoGroup.DELETE("", h.DeleteHealthInfo)
	}
}

func (h *Handler) CreateHealthInfo(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not authenticated"})
		return
	}

	var dto CreateHealthInformationDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	dto.UserID = userID.(string)
	healthInfo, err := h.service.CreateHealthInfo(dto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create health information"})
		return
	}

	c.JSON(http.StatusCreated, healthInfo)
}

func (h *Handler) GetHealthInfo(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not authenticated"})
		return
	}

	healthInfo, err := h.service.GetHealthInfo(userID.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get health information"})
		return
	}

	if healthInfo == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "health information not found"})
		return
	}

	c.JSON(http.StatusOK, healthInfo)
}

func (h *Handler) UpdateHealthInfo(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not authenticated"})
		return
	}

	var dto UpdateHealthInformationDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	healthInfo, err := h.service.UpdateHealthInfo(userID.(string), dto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update health information"})
		return
	}

	c.JSON(http.StatusOK, healthInfo)
}

func (h *Handler) DeleteHealthInfo(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not authenticated"})
		return
	}

	err := h.service.DeleteHealthInfo(userID.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete health information"})
		return
	}

	c.Status(http.StatusNoContent)
}
