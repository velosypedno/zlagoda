package handlers

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/velosypedno/zlagoda/internal/config"
	"github.com/velosypedno/zlagoda/internal/models"
	"github.com/velosypedno/zlagoda/internal/utils"
)

type receiptCreator interface {
	CreateReceipt(c models.ReceiptCreate) (string, error)
}

func NewReceiptCreatePOSTHandler(service receiptCreator, cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		type request struct {
			EmployeeId *string  `json:"employee_id" binding:"required,len=10"`
			CardNumber *string  `json:"card_number" binding:"omitempty,len=13"`
			PrintDate  *string  `json:"print_date" binding:"required"`
			TotalSum   *float64 `json:"sum_total" binding:"required,gte=0"`
		}
		var req request
		if err := c.ShouldBindJSON(&req); err != nil {
			log.Printf("[ReceiptCreatePOST] BindJSON error: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: " + err.Error()})
			return
		}

		var printDate time.Time
		var parseError error = fmt.Errorf("no valid date format found")

		// Try multiple date formats
		formats := []string{
			"2006-01-02",           // Date only
			"2006-01-02T15:04:05",  // ISO format without timezone
			"2006-01-02T15:04:05Z", // ISO format with Z
			"2006-01-02 15:04:05",  // YYYY-MM-DD HH:MM:SS
			"2006-01-02 15:04",     // YYYY-MM-DD HH:MM
		}

		for _, format := range formats {
			if parsedTime, parseErr := time.Parse(format, *req.PrintDate); parseErr == nil {
				printDate = parsedTime
				parseError = nil
				break
			}
		}

		if parseError != nil {
			log.Printf("[ReceiptCreatePOST] Failed to parse print_date '%s'", *req.PrintDate)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: invalid print date format"})
			return
		}

		if !utils.IsDecimalValid(*req.TotalSum) {
			log.Printf("[ReceiptCreatePOST] Invalid total sum: %v", *req.TotalSum)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: invalid total sum"})
			return
		}

		var VAT float64 = cfg.VAT_RATE * *req.TotalSum

		model := models.ReceiptCreate{
			EmployeeId: req.EmployeeId,
			CardNumber: req.CardNumber,
			PrintDate:  &printDate,
			TotalSum:   req.TotalSum,
			VAT:        &VAT,
		}

		id, err := service.CreateReceipt(model)
		if err != nil {
			log.Printf("[ReceiptCreatePOST] Service error: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create receipt: " + err.Error()})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"id": id})
	}
}

type receiptCompleteCreator interface {
	CreateReceiptComplete(c models.ReceiptCreateComplete, vatRate float64) (string, error)
}

func NewReceiptCreateCompletePOSTHandler(service receiptCompleteCreator, cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		type request struct {
			EmployeeId *string `json:"employee_id" binding:"required,len=10"`
			CardNumber *string `json:"card_number" binding:"omitempty,len=13"`
			PrintDate  *string `json:"print_date" binding:"required"`
			Items      []struct {
				UPC           *string  `json:"upc" binding:"required,len=12"`
				ProductNumber *int     `json:"product_number" binding:"required,gte=1"`
				SellingPrice  *float64 `json:"selling_price" binding:"required,gte=0"`
			} `json:"items" binding:"required,dive"`
		}
		var req request
		if err := c.ShouldBindJSON(&req); err != nil {
			log.Printf("[ReceiptCreateCompletePOST] BindJSON error: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: " + err.Error()})
			return
		}

		var printDate time.Time
		var parseError error = fmt.Errorf("no valid date format found")

		// Try multiple date formats
		formats := []string{
			"2006-01-02",           // Date only
			"2006-01-02T15:04:05",  // ISO format without timezone
			"2006-01-02T15:04:05Z", // ISO format with Z
			"2006-01-02 15:04:05",  // YYYY-MM-DD HH:MM:SS
			"2006-01-02 15:04",     // YYYY-MM-DD HH:MM
		}

		for _, format := range formats {
			if parsedTime, parseErr := time.Parse(format, *req.PrintDate); parseErr == nil {
				printDate = parsedTime
				parseError = nil
				break
			}
		}

		if parseError != nil {
			log.Printf("[ReceiptCreateCompletePOST] Failed to parse print_date '%s'", *req.PrintDate)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: invalid print date format"})
			return
		}

		var items []models.ReceiptItem
		for _, item := range req.Items {
			if !utils.IsDecimalValid(*item.SellingPrice) {
				log.Printf("[ReceiptCreateCompletePOST] Invalid selling price: %v", *item.SellingPrice)
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: invalid selling price"})
				return
			}
			items = append(items, models.ReceiptItem{
				UPC:           item.UPC,
				ProductNumber: item.ProductNumber,
				SellingPrice:  item.SellingPrice,
			})
		}

		model := models.ReceiptCreateComplete{
			EmployeeId: req.EmployeeId,
			CardNumber: req.CardNumber,
			PrintDate:  &printDate,
			Items:      items,
		}

		id, err := service.CreateReceiptComplete(model, cfg.VAT_RATE)
		if err != nil {
			log.Printf("[ReceiptCreateCompletePOST] Service error: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create receipt: " + err.Error()})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"id": id})
	}
}

type receiptReader interface {
	GetReceiptByReceiptNumber(receiptNumber string) (models.ReceiptRetrieve, error)
	GetReceipts() ([]models.ReceiptRetrieve, error)
}

func NewReceiptRetrieveGETHandler(service receiptReader) gin.HandlerFunc {
	return func(c *gin.Context) {
		type response struct {
			ReceiptNumber *string  `json:"receipt_number"`
			EmployeeId    *string  `json:"employee_id"`
			CardNumber    *string  `json:"card_number"`
			PrintDate     *string  `json:"print_date"`
			TotalSum      *float64 `json:"sum_total"`
			VAT           *float64 `json:"vat"`
		}

		receiptNumber := c.Param("receipt_number")
		if len(receiptNumber) != 10 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid receipt number"})
			return
		}

		receipt, err := service.GetReceiptByReceiptNumber(receiptNumber)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Receipt not found: " + err.Error()})
			return
		}

		printDate := receipt.PrintDate.Format("2006-01-02 15:04:05")

		resp := response{
			ReceiptNumber: receipt.ReceiptNumber,
			EmployeeId:    receipt.EmployeeId,
			CardNumber:    receipt.CardNumber,
			PrintDate:     &printDate,
			TotalSum:      receipt.TotalSum,
			VAT:           receipt.VAT,
		}

		c.JSON(http.StatusOK, resp)
	}
}

func NewReceiptsListGETHandler(service receiptReader) gin.HandlerFunc {
	type responseItem struct {
		ReceiptNumber *string  `json:"receipt_number"`
		EmployeeId    *string  `json:"employee_id"`
		CardNumber    *string  `json:"card_number"`
		PrintDate     *string  `json:"print_date"`
		TotalSum      *float64 `json:"sum_total"`
		VAT           *float64 `json:"vat"`
	}

	return func(c *gin.Context) {
		receipts, err := service.GetReceipts()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve receipts: " + err.Error()})
			return
		}

		var resp []responseItem
		for _, receipt := range receipts {
			printDate := receipt.PrintDate.Format("2006-01-02 15:04:05")
			resp = append(resp, responseItem{
				ReceiptNumber: receipt.ReceiptNumber,
				EmployeeId:    receipt.EmployeeId,
				CardNumber:    receipt.CardNumber,
				PrintDate:     &printDate,
				TotalSum:      receipt.TotalSum,
				VAT:           receipt.VAT,
			})
		}

		c.JSON(http.StatusOK, resp)
	}
}

type receiptRemover interface {
	DeleteReceipt(receiptNumber string) error
}

func NewReceiptDeleteDELETEHandler(service receiptRemover) gin.HandlerFunc {
	return func(c *gin.Context) {
		receiptNumber := c.Param("receipt_number")
		if len(receiptNumber) != 10 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid receipt number"})
			return
		}

		err := service.DeleteReceipt(receiptNumber)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete receipt: " + err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Receipt deleted successfully"})
	}
}

type receiptUpdater interface {
	UpdateReceipt(receiptNumber string, c models.ReceiptUpdate) error
	GetReceiptByReceiptNumber(receiptNumber string) (models.ReceiptRetrieve, error)
}

func NewReceiptUpdatePATCHHandler(service receiptUpdater, cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		receiptNumber := c.Param("receipt_number")
		if len(receiptNumber) != 10 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid receipt number"})
			return
		}

		type request struct {
			EmployeeId *string  `json:"employee_id" binding:"omitempty,len=10"`
			CardNumber *string  `json:"card_number" binding:"omitempty,len=13"`
			PrintDate  *string  `json:"print_date"`
			TotalSum   *float64 `json:"sum_total" binding:"omitempty,gte=0"`
		}
		var req request
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: " + err.Error()})
			return
		}

		receiptCurrentState, err := service.GetReceiptByReceiptNumber(receiptNumber)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Receipt not found: " + err.Error()})
			return
		}

		currentPrintDateStr := receiptCurrentState.PrintDate.Format("2006-01-02 15:04:05")
		if req.EmployeeId == nil {
			req.EmployeeId = receiptCurrentState.EmployeeId
		}
		if req.CardNumber == nil {
			req.CardNumber = receiptCurrentState.CardNumber
		}
		if req.PrintDate == nil {
			req.PrintDate = &currentPrintDateStr
		}
		if req.TotalSum == nil {
			req.TotalSum = receiptCurrentState.TotalSum
		}

		var printDate time.Time
		var parseError error = fmt.Errorf("no valid date format found")

		// Try multiple date formats
		formats := []string{
			"2006-01-02",           // Date only
			"2006-01-02T15:04:05",  // ISO format without timezone
			"2006-01-02T15:04:05Z", // ISO format with Z
			"2006-01-02 15:04:05",  // YYYY-MM-DD HH:MM:SS
			"2006-01-02 15:04",     // YYYY-MM-DD HH:MM
		}

		for _, format := range formats {
			if parsedTime, parseErr := time.Parse(format, *req.PrintDate); parseErr == nil {
				printDate = parsedTime
				parseError = nil
				break
			}
		}

		if parseError != nil {
			log.Printf("[ReceiptUpdatePATCH] Failed to parse print_date '%s'", *req.PrintDate)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: invalid print date format"})
			return
		}

		if !utils.IsDecimalValid(*req.TotalSum) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: invalid total sum"})
			return
		}

		var VAT float64 = cfg.VAT_RATE * *req.TotalSum

		model := models.ReceiptUpdate{
			EmployeeId: req.EmployeeId,
			CardNumber: req.CardNumber,
			PrintDate:  &printDate,
			TotalSum:   req.TotalSum,
			VAT:        &VAT,
		}

		err = service.UpdateReceipt(receiptNumber, model)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update receipt: " + err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Receipt updated successfully"})
	}
}
