package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/velosypedno/zlagoda/internal/services"
)

type LoginPayload struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func NewLoginPOSTHandler(service services.LoginService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var payload LoginPayload
		if err := c.ShouldBindJSON(&payload); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		token, err := service.Login(c, payload.Login, payload.Password)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"token": token})
	}
} 