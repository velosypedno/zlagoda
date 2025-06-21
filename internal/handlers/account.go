package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/velosypedno/zlagoda/internal/services"
)

func NewAccountGETHandler(service services.AccountService) gin.HandlerFunc {
	return func(c *gin.Context) {
		employeeID, exists := c.Get("employee_id")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Employee ID not found in context"})
			return
		}

		employeeIDStr, ok := employeeID.(string)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid employee ID type"})
			return
		}

		account, err := service.GetAccount(c, employeeIDStr)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, account)
	}
} 