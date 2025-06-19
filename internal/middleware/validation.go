package middleware

import (
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// ValidateID middleware validates that the ID parameter is a valid integer
func ValidateID(paramName string) gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param(paramName)
		if idStr == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Missing " + paramName + " parameter"})
			c.Abort()
			return
		}

		if _, err := strconv.Atoi(idStr); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid " + paramName + " format"})
			c.Abort()
			return
		}

		c.Next()
	}
}

// ValidateEmployeeID middleware validates employee ID format (10 characters)
func ValidateEmployeeID() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		if len(id) != 10 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid employee ID format"})
			c.Abort()
			return
		}

		// Check if ID contains only alphanumeric characters
		if !isAlphanumeric(id) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Employee ID must contain only alphanumeric characters"})
			c.Abort()
			return
		}

		c.Next()
	}
}

// ValidateCardNumber middleware validates customer card number format (13 characters)
func ValidateCardNumber() gin.HandlerFunc {
	return func(c *gin.Context) {
		cardNumber := c.Param("card_number")
		if len(cardNumber) != 13 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid card number format"})
			c.Abort()
			return
		}

		// Check if card number contains only alphanumeric characters
		if !isAlphanumeric(cardNumber) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Card number must contain only alphanumeric characters"})
			c.Abort()
			return
		}

		c.Next()
	}
}

// ValidateReceiptNumber middleware validates receipt number format (10 characters)
func ValidateReceiptNumber() gin.HandlerFunc {
	return func(c *gin.Context) {
		receiptNumber := c.Param("receipt_number")
		if len(receiptNumber) != 10 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid receipt number format"})
			c.Abort()
			return
		}

		// Check if receipt number contains only alphanumeric characters
		if !isAlphanumeric(receiptNumber) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Receipt number must contain only alphanumeric characters"})
			c.Abort()
			return
		}

		c.Next()
	}
}

// ValidatePhoneNumber middleware validates Ukrainian phone number format
func ValidatePhoneNumber() gin.HandlerFunc {
	phoneRegex := regexp.MustCompile(`^\+380\d{9}$`)

	return func(c *gin.Context) {
		var requestBody map[string]interface{}
		if err := c.ShouldBindJSON(&requestBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
			c.Abort()
			return
		}

		if phoneNumber, exists := requestBody["phone_number"]; exists && phoneNumber != nil {
			phoneStr, ok := phoneNumber.(string)
			if !ok {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Phone number must be a string"})
				c.Abort()
				return
			}

			if !phoneRegex.MatchString(phoneStr) {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid phone number format. Must be +380XXXXXXXXX"})
				c.Abort()
				return
			}
		}

		// Re-bind the JSON for the next handler
		c.Set("validated_body", requestBody)
		c.Next()
	}
}

// ValidateJSONContentType middleware ensures the request has JSON content type
func ValidateJSONContentType() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method == "POST" || c.Request.Method == "PUT" || c.Request.Method == "PATCH" {
			contentType := c.GetHeader("Content-Type")
			if !strings.Contains(contentType, "application/json") {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Content-Type must be application/json"})
				c.Abort()
				return
			}
		}
		c.Next()
	}
}

// ValidateQueryParams middleware validates common query parameters
func ValidateQueryParams() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Validate limit parameter if present
		if limitStr := c.Query("limit"); limitStr != "" {
			limit, err := strconv.Atoi(limitStr)
			if err != nil || limit < 1 || limit > 1000 {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit parameter. Must be between 1 and 1000"})
				c.Abort()
				return
			}
		}

		// Validate offset parameter if present
		if offsetStr := c.Query("offset"); offsetStr != "" {
			offset, err := strconv.Atoi(offsetStr)
			if err != nil || offset < 0 {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid offset parameter. Must be >= 0"})
				c.Abort()
				return
			}
		}

		// Validate sort parameter if present
		if sort := c.Query("sort"); sort != "" {
			validSortFields := []string{"id", "name", "date", "created_at", "updated_at"}
			if !contains(validSortFields, sort) {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid sort parameter"})
				c.Abort()
				return
			}
		}

		// Validate order parameter if present
		if order := c.Query("order"); order != "" {
			if order != "asc" && order != "desc" {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order parameter. Must be 'asc' or 'desc'"})
				c.Abort()
				return
			}
		}

		c.Next()
	}
}

// Helper functions

func isAlphanumeric(s string) bool {
	alphanumericRegex := regexp.MustCompile(`^[a-zA-Z0-9]+$`)
	return alphanumericRegex.MatchString(s)
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

// ValidateRole middleware validates employee role
func ValidateRole() gin.HandlerFunc {
	validRoles := []string{"manager", "cashier", "admin"}

	return func(c *gin.Context) {
		var requestBody map[string]interface{}
		if err := c.ShouldBindJSON(&requestBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
			c.Abort()
			return
		}

		if role, exists := requestBody["empl_role"]; exists && role != nil {
			roleStr, ok := role.(string)
			if !ok {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Role must be a string"})
				c.Abort()
				return
			}

			if !contains(validRoles, roleStr) {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid role. Must be one of: manager, cashier, admin"})
				c.Abort()
				return
			}
		}

		// Re-bind the JSON for the next handler
		c.Set("validated_body", requestBody)
		c.Next()
	}
}

// SanitizeInput middleware removes potential harmful characters from string inputs
func SanitizeInput() gin.HandlerFunc {
	return func(c *gin.Context) {
		var requestBody map[string]interface{}
		if err := c.ShouldBindJSON(&requestBody); err != nil {
			// If it's not JSON, skip sanitization
			c.Next()
			return
		}

		sanitizedBody := sanitizeMap(requestBody)
		c.Set("sanitized_body", sanitizedBody)
		c.Next()
	}
}

func sanitizeMap(m map[string]interface{}) map[string]interface{} {
	sanitized := make(map[string]interface{})

	for key, value := range m {
		switch v := value.(type) {
		case string:
			// Remove potential SQL injection characters and trim whitespace
			sanitized[key] = strings.TrimSpace(strings.ReplaceAll(v, ";", ""))
		case map[string]interface{}:
			sanitized[key] = sanitizeMap(v)
		default:
			sanitized[key] = value
		}
	}

	return sanitized
}
