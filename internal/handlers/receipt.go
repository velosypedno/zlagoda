package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/velosypedno/zlagoda/internal/models"
)

func isReceiptUpdateValid(receipt models.ReceiptUpdate) bool {
	if len(receipt.EmployeeId) != 10 {
		return false
	}
	if len(receipt.CardNumber) != 13 {
		return false
	}
	if receipt.PrintDate.After(time.Now()) {
		return false
	}
	if receipt.TotalSum < 0 {
		return false
	}
	if !isDecimalValid(receipt.TotalSum) {
		return false
	}
	return true
}

type receiptCreator interface {
	CreateReceipt(c models.ReceiptCreate) (string, error)
}

func NewReceiptCreatePOSTHandler(service receiptCreator) gin.HandlerFunc {
	return func(c *gin.Context) {
		type request struct {
			EmployeeId string  `json:"employee_id" binding:"required,len=10"`
			CardNumber string  `json:"card_number" binding:"len=13"`
			PrintDate  string  `json:"print_date" binding:"required"`
			TotalSum   float64 `json:"sum_total" binding:"required,gte=0"`
		}
		var req request
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: " + err.Error()})
			return
		}

		printDate, err := time.Parse("2006-01-02", req.PrintDate)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: invalid print date format"})
			return
		}
		if printDate.After(time.Now()) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: invalid print date"})
			return
		}

		model := models.ReceiptCreate{
			EmployeeId: req.EmployeeId,
			CardNumber: req.CardNumber,
			PrintDate:  printDate,
			TotalSum:   req.TotalSum,
			VAT:        0.2 * req.TotalSum,
		}

		id, err := service.CreateReceipt(model)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create employee: " + err.Error()})
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
			ReceiptNumber string  `json:"receipt_number"`
			EmployeeId    string  `json:"employee_id"`
			CardNumber    string  `json:"card_number"`
			PrintDate     string  `json:"print_date"`
			TotalSum      float64 `json:"sum_total"`
			VAT           float64 `json:"vat"`
		}
		var receiptNumber string = c.Param("receipt_number")
		if len(receiptNumber) != 10 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid receipt number"})
			return
		}

		receipt, err := service.GetReceiptByReceiptNumber(receiptNumber)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Receipt not found: " + err.Error()})
			return
		}

		resp := response{
			ReceiptNumber: receipt.ReceiptNumber,
			EmployeeId:    receipt.EmployeeId,
			CardNumber:    receipt.CardNumber,
			PrintDate:     receipt.PrintDate.Format("2006-01-02"),
			TotalSum:      receipt.TotalSum,
			VAT:           receipt.VAT,
		}

		c.JSON(http.StatusOK, resp)
	}
}

func NewReceiptsListGETHandler(service receiptReader) gin.HandlerFunc {
	type responseItem struct {
		ReceiptNumber string  `json:"receipt_number"`
		EmployeeId    string  `json:"employee_id"`
		CardNumber    string  `json:"card_number"`
		PrintDate     string  `json:"print_date"`
		TotalSum      float64 `json:"sum_total"`
		VAT           float64 `json:"vat"`
	}

	return func(c *gin.Context) {
		receipts, err := service.GetReceipts()
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Failed to retrieve receipts: " + err.Error()})
			return
		}

		var resp []responseItem
		for _, receipt := range receipts {
			resp = append(resp, responseItem{
				ReceiptNumber: receipt.ReceiptNumber,
				EmployeeId:    receipt.EmployeeId,
				CardNumber:    receipt.CardNumber,
				PrintDate:     receipt.PrintDate.Format("2006-01-02"),
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
		var receiptNumber string = c.Param("receipt_number")
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

func NewReceiptUpdatePATCHHandler(service receiptUpdater) gin.HandlerFunc {
	return func(c *gin.Context) {
		receiptNumber := c.Param("receipt_number")
		if len(receiptNumber) != 10 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid receipt number"})
			return
		}

		type request struct {
			EmployeeId *string  `json:"employee_id"`
			CardNumber *string  `json:"card_number"`
			PrintDate  *string  `json:"print_date"`
			TotalSum   *float64 `json:"sum_total"`
		}
		var req request
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: " + err.Error()})
			return
		}

		receiptCurrentState, err := service.GetReceiptByReceiptNumber(receiptNumber)
		if err != nil {
			receiptCurrentState = models.ReceiptRetrieve{}
		}
		currentPrintDateStr := receiptCurrentState.PrintDate.Format("2006-01-02")
		if req.EmployeeId == nil {
			req.EmployeeId = &receiptCurrentState.EmployeeId
		}
		if req.CardNumber == nil {
			req.CardNumber = &receiptCurrentState.CardNumber
		}
		if req.PrintDate == nil {
			req.PrintDate = &currentPrintDateStr
		}
		if req.TotalSum == nil {
			req.TotalSum = &receiptCurrentState.TotalSum
		}

		printDate, err := time.Parse("2006-01-02", *req.PrintDate)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: invalid print date format"})
			return
		}

		model := models.ReceiptUpdate{
			EmployeeId: *req.EmployeeId,
			CardNumber: *req.CardNumber,
			PrintDate:  printDate,
			TotalSum:   *req.TotalSum,
			VAT:        0.2 * *req.TotalSum,
		}
		if !isReceiptUpdateValid(model) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: input values are out of bounds"})
			return
		}

		err = service.UpdateReceipt(receiptNumber, model)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update receipt: " + err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Receipt updated successfully"})
	}
}
