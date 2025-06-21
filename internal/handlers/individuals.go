package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/velosypedno/zlagoda/internal/models"
)

type individualsService interface {
	QueryVlad1(categoryID int, months int) ([]models.Vlad1Response, error)
	QueryVlad2() ([]models.Vlad2Response, error)
	QueryArthur1(startDate, endDate string) ([]models.Arthur1Response, error)
	QueryArthur2() ([]models.Arthur2Response, error)
	QueryOleksii1(discountThreshold int) ([]models.Oleksii1Response, error)
	QueryOleksii2() ([]models.Oleksii2Response, error)
}

// Vlad1 - Most sold product in a category within a time period
func NewVlad1GETHandler(service individualsService) gin.HandlerFunc {
	return func(c *gin.Context) {
		categoryIDStr := c.Query("category_id")
		monthsStr := c.DefaultQuery("months", "1")

		if categoryIDStr == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "category_id parameter is required"})
			return
		}

		categoryID, err := strconv.Atoi(categoryIDStr)
		if err != nil || categoryID <= 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category_id parameter"})
			return
		}

		months, err := strconv.Atoi(monthsStr)
		if err != nil || months <= 0 || months > 12 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid months parameter (1-12)"})
			return
		}

		results, err := service.QueryVlad1(categoryID, months)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to execute Vlad1 query: " + err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"description": "Most sold products in category within time period",
			"parameters": gin.H{
				"category_id": categoryID,
				"months":      months,
			},
			"results": results,
		})
	}
}

// Vlad2 - Employees who never sold promotional products
func NewVlad2GETHandler(service individualsService) gin.HandlerFunc {
	return func(c *gin.Context) {
		results, err := service.QueryVlad2()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to execute Vlad2 query: " + err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"description": "Employees who never sold promotional products",
			"results":     results,
		})
	}
}

// Arthur1 - Category sales statistics within date range
func NewArthur1GETHandler(service individualsService) gin.HandlerFunc {
	return func(c *gin.Context) {
		startDate := c.Query("start_date")
		endDate := c.Query("end_date")

		if startDate == "" || endDate == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "start_date and end_date parameters are required (YYYY-MM-DD)"})
			return
		}

		// Basic date format validation
		if len(startDate) != 10 || len(endDate) != 10 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format, use YYYY-MM-DD"})
			return
		}

		results, err := service.QueryArthur1(startDate, endDate)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to execute Arthur1 query: " + err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"description": "Category sales statistics within date range",
			"parameters": gin.H{
				"start_date": startDate,
				"end_date":   endDate,
			},
			"results": results,
		})
	}
}

// Arthur2 - Products in store that have never been sold and are not promotional
func NewArthur2GETHandler(service individualsService) gin.HandlerFunc {
	return func(c *gin.Context) {
		results, err := service.QueryArthur2()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to execute Arthur2 query: " + err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"description": "Products in store that have never been sold and are not promotional",
			"results":     results,
		})
	}
}

// Oleksii1 - Cashiers who served customers with high discount
func NewOleksii1GETHandler(service individualsService) gin.HandlerFunc {
	return func(c *gin.Context) {
		discountThresholdStr := c.DefaultQuery("discount_threshold", "10")

		discountThreshold, err := strconv.Atoi(discountThresholdStr)
		if err != nil || discountThreshold < 0 || discountThreshold > 100 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid discount_threshold parameter (0-100)"})
			return
		}

		results, err := service.QueryOleksii1(discountThreshold)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to execute Oleksii1 query: " + err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"description": "Cashiers who served customers with high discount",
			"parameters": gin.H{
				"discount_threshold": discountThreshold,
			},
			"results": results,
		})
	}
}

// Oleksii2 - Customers who bought from all categories in the last month
func NewOleksii2GETHandler(service individualsService) gin.HandlerFunc {
	return func(c *gin.Context) {
		results, err := service.QueryOleksii2()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to execute Oleksii2 query: " + err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"description": "Customers who bought from all categories in the last month",
			"results":     results,
		})
	}
}
