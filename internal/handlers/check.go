package handlers

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/velosypedno/zlagoda/internal/config"
	"github.com/velosypedno/zlagoda/internal/models"
)

type CheckService interface {
	CreateCheck(req models.CheckCreate, vatRate float64) (*models.CheckCreateResponse, error)
}

func NewCheckCreatePOSTHandler(service CheckService, cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req models.CheckCreate
		if err := c.ShouldBindJSON(&req); err != nil {
			log.Printf("[CheckCreatePOST] Invalid input: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: " + err.Error()})
			return
		}
		// Validate print date
		printDate, err := time.Parse("2006-01-02", req.PrintDate)
		if err != nil || printDate.After(time.Now()) {
			log.Printf("[CheckCreatePOST] Invalid print date: %v", req.PrintDate)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid print date"})
			return
		}
		resp, err := service.CreateCheck(req, cfg.VAT_RATE)
		if err != nil {
			log.Printf("[CheckCreatePOST] Service error: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, resp)
	}
} 