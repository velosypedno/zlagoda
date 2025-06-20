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

type checkReader interface {
	GetChecks() ([]models.ReceiptRetrieve, error)
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

func NewChecksListGETHandler(service checkReader) gin.HandlerFunc {
	type responseItem struct {
		ReceiptNumber *string  `json:"receipt_number"`
		EmployeeId    *string  `json:"employee_id"`
		CardNumber    *string  `json:"card_number"`
		PrintDate     *string  `json:"print_date"`
		TotalSum      *float64 `json:"sum_total"`
		VAT           *float64 `json:"vat"`
	}

	return func(c *gin.Context) {
		checks, err := service.GetChecks()
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to retrieve checks: " + err.Error()})
			return
		}
		var resp []responseItem
		for _, check := range checks {
			printDate := ""
			if check.PrintDate != nil {
				printDate = check.PrintDate.Format("2006-01-02")
			}
			resp = append(resp, responseItem{
				ReceiptNumber: check.ReceiptNumber,
				EmployeeId:    check.EmployeeId,
				CardNumber:    check.CardNumber,
				PrintDate:     &printDate,
				TotalSum:      check.TotalSum,
				VAT:           check.VAT,
			})
		}
		c.JSON(200, resp)
	}
}

func NewCheckRetrieveGETHandler(service checkReader) gin.HandlerFunc {
	return func(c *gin.Context) {
		receiptNumber := c.Param("receipt_number")
		if len(receiptNumber) != 10 {
			c.JSON(400, gin.H{"error": "Invalid receipt number"})
			return
		}
		checks, err := service.GetChecks()
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to retrieve checks: " + err.Error()})
			return
		}
		var found *models.ReceiptRetrieve
		for i := range checks {
			if checks[i].ReceiptNumber != nil && *checks[i].ReceiptNumber == receiptNumber {
				found = &checks[i]
				break
			}
		}
		if found == nil {
			c.JSON(404, gin.H{"error": "Check not found"})
			return
		}
		printDate := ""
		if found.PrintDate != nil {
			printDate = found.PrintDate.Format("2006-01-02")
		}
		c.JSON(200, gin.H{
			"receipt_number": found.ReceiptNumber,
			"employee_id":    found.EmployeeId,
			"card_number":    found.CardNumber,
			"print_date":     printDate,
			"sum_total":      found.TotalSum,
			"vat":            found.VAT,
		})
	}
}
