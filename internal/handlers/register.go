package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/velosypedno/zlagoda/internal/models"
	"github.com/velosypedno/zlagoda/internal/services"
)

type RegisterPayload struct {
	Login         string  `json:"login" binding:"required"`
	Password      string  `json:"password" binding:"required,min=6"`
	Surname       string  `json:"surname" binding:"required"`
	Name          string  `json:"name" binding:"required"`
	Patronymic    *string `json:"patronymic"`
	Role          string  `json:"role" binding:"required"`
	Salary        float64 `json:"salary" binding:"required,min=0"`
	DateOfBirth   string  `json:"date_of_birth" binding:"required"`
	DateOfStart   string  `json:"date_of_start" binding:"required"`
	PhoneNumber   string  `json:"phone_number" binding:"required"`
	City          string  `json:"city" binding:"required"`
	Street        string  `json:"street" binding:"required"`
	ZipCode       string  `json:"zip_code" binding:"required"`
}

func NewRegisterPOSTHandler(service services.RegisterService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var payload RegisterPayload
		if err := c.ShouldBindJSON(&payload); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Parse dates
		dateOfBirth, err := time.Parse("2006-01-02", payload.DateOfBirth)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date of birth format. Use YYYY-MM-DD"})
			return
		}

		dateOfStart, err := time.Parse("2006-01-02", payload.DateOfStart)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date of start format. Use YYYY-MM-DD"})
			return
		}

		// Convert payload to EmployeeCreate model
		employee := models.EmployeeCreate{
			Surname:     &payload.Surname,
			Name:        &payload.Name,
			Patronymic:  payload.Patronymic,
			Role:        &payload.Role,
			Salary:      &payload.Salary,
			DateOfBirth: &dateOfBirth,
			DateOfStart: &dateOfStart,
			PhoneNumber: &payload.PhoneNumber,
			City:        &payload.City,
			Street:      &payload.Street,
			ZipCode:     &payload.ZipCode,
		}

		token, err := service.Register(c, employee, payload.Login, payload.Password)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"token": token})
	}
} 