package medicaldata

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
	medicalDataGroup := router.Group("/medical-data")
	{
		medicalDataGroup.POST("", h.CreateMedicalData)
		medicalDataGroup.GET("", h.GetMedicalData)
		medicalDataGroup.PUT("", h.UpdateMedicalData)
		medicalDataGroup.DELETE("", h.DeleteMedicalData)
	}
}

func (h *Handler) CreateMedicalData(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not authenticated"})
		return
	}

	var dto CreateMedicalDataDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	medicalData, err := h.service.CreateMedicalData(userID.(string), dto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create medical data"})
		return
	}

	c.JSON(http.StatusCreated, medicalData)
}

func (h *Handler) GetMedicalData(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not authenticated"})
		return
	}

	medicalData, err := h.service.GetMedicalData(userID.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get medical data"})
		return
	}

	if medicalData == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "medical data not found"})
		return
	}

	c.JSON(http.StatusOK, medicalData)
}

func (h *Handler) UpdateMedicalData(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not authenticated"})
		return
	}

	var dto UpdateMedicalDataDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	medicalData, err := h.service.UpdateMedicalData(userID.(string), dto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update medical data"})
		return
	}

	c.JSON(http.StatusOK, medicalData)
}

func (h *Handler) DeleteMedicalData(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not authenticated"})
		return
	}

	err := h.service.DeleteMedicalData(userID.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete medical data"})
		return
	}

	c.Status(http.StatusNoContent)
}
