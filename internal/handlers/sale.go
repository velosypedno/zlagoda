package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/velosypedno/zlagoda/internal/models"
	"github.com/velosypedno/zlagoda/internal/utils"
)

type saleCreator interface {
	CreateSale(s models.SaleCreate) error
}

func NewSaleCreatePOSTHandler(service saleCreator) gin.HandlerFunc {
	return func(c *gin.Context) {
		type request struct {
			UPC           string  `json:"upc" binding:"required,len=12"`
			ReceiptNumber string  `json:"receipt_number" binding:"required,len=10"`
			ProductNumber int     `json:"product_number" binding:"required,gte=1"`
			SellingPrice  float64 `json:"selling_price" binding:"required,gte=0"`
		}
		var req request
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: " + err.Error()})
			return
		}

		if !utils.IsAmountValid(req.SellingPrice) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: invalid selling price"})
			return
		}

		model := models.SaleCreate{
			UPC:           req.UPC,
			ReceiptNumber: req.ReceiptNumber,
			ProductNumber: req.ProductNumber,
			SellingPrice:  req.SellingPrice,
		}

		err := service.CreateSale(model)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create sale: " + err.Error()})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"message": "Sale created successfully"})
	}
}

type saleReader interface {
	GetSaleByKey(upc, receiptNumber string) (models.SaleRetrieve, error)
	GetSalesByReceipt(receiptNumber string) ([]models.SaleRetrieve, error)
	GetSalesByUPC(upc string) ([]models.SaleRetrieve, error)
	GetAllSales() ([]models.SaleRetrieve, error)
	GetSalesWithDetails() ([]models.SaleWithDetails, error)
	GetSalesWithDetailsByReceipt(receiptNumber string) ([]models.SaleWithDetails, error)
}

func NewSaleRetrieveGETHandler(service saleReader) gin.HandlerFunc {
	return func(c *gin.Context) {
		type response struct {
			UPC           string  `json:"upc"`
			ReceiptNumber string  `json:"receipt_number"`
			ProductNumber int     `json:"product_number"`
			SellingPrice  float64 `json:"selling_price"`
		}

		upc := c.Param("upc")
		receiptNumber := c.Param("receipt_number")

		if len(upc) != 12 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UPC format"})
			return
		}

		if len(receiptNumber) != 10 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid receipt number format"})
			return
		}

		sale, err := service.GetSaleByKey(upc, receiptNumber)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Sale not found: " + err.Error()})
			return
		}

		resp := response{
			UPC:           sale.UPC,
			ReceiptNumber: sale.ReceiptNumber,
			ProductNumber: sale.ProductNumber,
			SellingPrice:  sale.SellingPrice,
		}

		c.JSON(http.StatusOK, resp)
	}
}

func NewSalesByReceiptGETHandler(service saleReader) gin.HandlerFunc {
	type responseItem struct {
		UPC           string  `json:"upc"`
		ReceiptNumber string  `json:"receipt_number"`
		ProductNumber int     `json:"product_number"`
		SellingPrice  float64 `json:"selling_price"`
	}

	return func(c *gin.Context) {
		receiptNumber := c.Param("receipt_number")
		if len(receiptNumber) != 10 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid receipt number format"})
			return
		}

		sales, err := service.GetSalesByReceipt(receiptNumber)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve sales: " + err.Error()})
			return
		}

		var resp []responseItem
		for _, sale := range sales {
			resp = append(resp, responseItem{
				UPC:           sale.UPC,
				ReceiptNumber: sale.ReceiptNumber,
				ProductNumber: sale.ProductNumber,
				SellingPrice:  sale.SellingPrice,
			})
		}

		c.JSON(http.StatusOK, resp)
	}
}

func NewSalesByUPCGETHandler(service saleReader) gin.HandlerFunc {
	type responseItem struct {
		UPC           string  `json:"upc"`
		ReceiptNumber string  `json:"receipt_number"`
		ProductNumber int     `json:"product_number"`
		SellingPrice  float64 `json:"selling_price"`
	}

	return func(c *gin.Context) {
		upc := c.Param("upc")
		if len(upc) != 12 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UPC format"})
			return
		}

		sales, err := service.GetSalesByUPC(upc)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve sales: " + err.Error()})
			return
		}

		var resp []responseItem
		for _, sale := range sales {
			resp = append(resp, responseItem{
				UPC:           sale.UPC,
				ReceiptNumber: sale.ReceiptNumber,
				ProductNumber: sale.ProductNumber,
				SellingPrice:  sale.SellingPrice,
			})
		}

		c.JSON(http.StatusOK, resp)
	}
}

func NewSalesListGETHandler(service saleReader) gin.HandlerFunc {
	type responseItem struct {
		UPC           string  `json:"upc"`
		ReceiptNumber string  `json:"receipt_number"`
		ProductNumber int     `json:"product_number"`
		SellingPrice  float64 `json:"selling_price"`
	}

	return func(c *gin.Context) {
		sales, err := service.GetAllSales()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve sales: " + err.Error()})
			return
		}

		var resp []responseItem
		for _, sale := range sales {
			resp = append(resp, responseItem{
				UPC:           sale.UPC,
				ReceiptNumber: sale.ReceiptNumber,
				ProductNumber: sale.ProductNumber,
				SellingPrice:  sale.SellingPrice,
			})
		}

		c.JSON(http.StatusOK, resp)
	}
}

func NewSalesWithDetailsListGETHandler(service saleReader) gin.HandlerFunc {
	return func(c *gin.Context) {
		sales, err := service.GetSalesWithDetails()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve sales with details: " + err.Error()})
			return
		}

		c.JSON(http.StatusOK, sales)
	}
}

func NewSalesWithDetailsByReceiptGETHandler(service saleReader) gin.HandlerFunc {
	return func(c *gin.Context) {
		receiptNumber := c.Param("receipt_number")
		if len(receiptNumber) != 10 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid receipt number format"})
			return
		}

		sales, err := service.GetSalesWithDetailsByReceipt(receiptNumber)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve sales with details: " + err.Error()})
			return
		}

		c.JSON(http.StatusOK, sales)
	}
}

type saleUpdater interface {
	UpdateSale(upc, receiptNumber string, s models.SaleUpdate) error
	GetSaleByKey(upc, receiptNumber string) (models.SaleRetrieve, error)
}

func NewSaleUpdatePATCHHandler(service saleUpdater) gin.HandlerFunc {
	return func(c *gin.Context) {
		upc := c.Param("upc")
		receiptNumber := c.Param("receipt_number")

		if len(upc) != 12 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UPC format"})
			return
		}

		if len(receiptNumber) != 10 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid receipt number format"})
			return
		}

		type request struct {
			ProductNumber *int     `json:"product_number" binding:"omitempty,gte=1"`
			SellingPrice  *float64 `json:"selling_price" binding:"omitempty,gte=0"`
		}
		var req request
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: " + err.Error()})
			return
		}

		// Check if sale exists
		_, err := service.GetSaleByKey(upc, receiptNumber)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Sale not found: " + err.Error()})
			return
		}

		if req.SellingPrice != nil && !utils.IsAmountValid(*req.SellingPrice) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: invalid selling price"})
			return
		}

		model := models.SaleUpdate{
			ProductNumber: req.ProductNumber,
			SellingPrice:  req.SellingPrice,
		}

		err = service.UpdateSale(upc, receiptNumber, model)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update sale: " + err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Sale updated successfully"})
	}
}

type saleRemover interface {
	DeleteSale(upc, receiptNumber string) error
	DeleteSalesByReceipt(receiptNumber string) error
}

func NewSaleDeleteDELETEHandler(service saleRemover) gin.HandlerFunc {
	return func(c *gin.Context) {
		upc := c.Param("upc")
		receiptNumber := c.Param("receipt_number")

		if len(upc) != 12 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UPC format"})
			return
		}

		if len(receiptNumber) != 10 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid receipt number format"})
			return
		}

		err := service.DeleteSale(upc, receiptNumber)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete sale: " + err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Sale deleted successfully"})
	}
}

func NewSalesByReceiptDeleteDELETEHandler(service saleRemover) gin.HandlerFunc {
	return func(c *gin.Context) {
		receiptNumber := c.Param("receipt_number")
		if len(receiptNumber) != 10 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid receipt number format"})
			return
		}

		err := service.DeleteSalesByReceipt(receiptNumber)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete sales: " + err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "All sales for receipt deleted successfully"})
	}
}

type saleAnalytics interface {
	GetReceiptTotal(receiptNumber string) (float64, error)
	GetSalesStatsByProduct(productID int, startDate, endDate string) (int, float64, error)
	GetTopSellingProducts(limit int) ([]struct {
		ProductID    int     `json:"product_id"`
		ProductName  string  `json:"product_name"`
		TotalSold    int     `json:"total_sold"`
		TotalRevenue float64 `json:"total_revenue"`
	}, error)
}

func NewReceiptTotalGETHandler(service saleAnalytics) gin.HandlerFunc {
	return func(c *gin.Context) {
		receiptNumber := c.Param("receipt_number")
		if len(receiptNumber) != 10 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid receipt number format"})
			return
		}

		total, err := service.GetReceiptTotal(receiptNumber)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to calculate receipt total: " + err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"receipt_number": receiptNumber,
			"total":          total,
		})
	}
}

func NewSalesStatsByProductGETHandler(service saleAnalytics) gin.HandlerFunc {
	return func(c *gin.Context) {
		productIDStr := c.Param("product_id")
		productID, err := strconv.Atoi(productIDStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
			return
		}

		startDate := c.Query("start_date")
		endDate := c.Query("end_date")

		if startDate == "" || endDate == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Missing start_date or end_date query parameters"})
			return
		}

		totalQuantity, totalRevenue, err := service.GetSalesStatsByProduct(productID, startDate, endDate)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get sales stats: " + err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"product_id":     productID,
			"start_date":     startDate,
			"end_date":       endDate,
			"total_quantity": totalQuantity,
			"total_revenue":  totalRevenue,
		})
	}
}

func NewTopSellingProductsGETHandler(service saleAnalytics) gin.HandlerFunc {
	return func(c *gin.Context) {
		limitStr := c.DefaultQuery("limit", "10")
		limit, err := strconv.Atoi(limitStr)
		if err != nil || limit < 1 || limit > 100 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit parameter (must be 1-100)"})
			return
		}

		products, err := service.GetTopSellingProducts(limit)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get top selling products: " + err.Error()})
			return
		}

		c.JSON(http.StatusOK, products)
	}
}
