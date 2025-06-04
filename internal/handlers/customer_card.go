package handlers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/velosypedno/zlagoda/internal/models"
)

func validateCustomerCard(card *models.CustomerCardCreate) bool {
	if len(card.Surname) > 50 {
		return false
	}
	if len(card.Name) > 50 {
		return false
	}
	if len(card.Patronymic) > 50 {
		return false
	}
	if !strings.HasPrefix(card.PhoneNumber, "+380") || len(card.PhoneNumber) != 13 {
		return false
	}
	if len(card.City) > 50 {
		return false
	}
	if len(card.Street) > 50 {
		return false
	}
	if len(card.ZipCode) > 9 {
		return false
	}
	if card.Percent < 0 || card.Percent > 100 {
		return false
	}
	return true
}

type customerCardCreator interface {
	CreateCustomerCard(c models.CustomerCardCreate) (string, error)
}

func NewCustomerCardCreatePOSTHandler(service customerCardCreator) gin.HandlerFunc {
	return func(c *gin.Context) {
		type request struct {
			Surname     string `json:"cust_surname" binding:"required"`
			Name        string `json:"cust_name" binding:"required"`
			Patronymic  string `json:"cust_patronymic"`
			PhoneNumber string `json:"phone_number" binding:"required"`
			City        string `json:"city"`
			Street      string `json:"street"`
			ZipCode     string `json:"zip_code"`
			Percent     int    `json:"percent" binding:"required"`
		}
		var req request
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: " + err.Error()})
			return
		}
		model := models.CustomerCardCreate{
			Surname:     req.Surname,
			Name:        req.Name,
			Patronymic:  req.Patronymic,
			PhoneNumber: req.PhoneNumber,
			City:        req.City,
			Street:      req.Street,
			ZipCode:     req.ZipCode,
			Percent:     req.Percent,
		}
		if !validateCustomerCard(&model) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Validation failed: input values are out of bounds"})
			return
		}

		id, err := service.CreateCustomerCard(model)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create customer card: " + err.Error()})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"id": id})
	}
}

type customerCardReader interface {
	GetCustomerCardByCardNumber(cardNumber string) (models.CustomerCardRetrieve, error)
	GetCustomerCards() ([]models.CustomerCardRetrieve, error)
}

func NewCustomerCardRetrieveGETHandler(service customerCardReader) gin.HandlerFunc {
	return func(c *gin.Context) {
		type response struct {
			CardNumber  string `json:"card_number"`
			Surname     string `json:"cust_surname"`
			Name        string `json:"cust_name"`
			Patronymic  string `json:"cust_patronymic"`
			PhoneNumber string `json:"phone_number"`
			City        string `json:"city"`
			Street      string `json:"street"`
			ZipCode     string `json:"zip_code"`
			Percent     int    `json:"percent"`
		}

		cardNumber := c.Param("card_number")
		if len(cardNumber) != 13 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid card number"})
			return
		}

		customerCard, err := service.GetCustomerCardByCardNumber(cardNumber)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Customer card not found: " + err.Error()})
			return
		}

		resp := response{
			CardNumber:  customerCard.CardNumber,
			Surname:     customerCard.Surname,
			Name:        customerCard.Name,
			Patronymic:  customerCard.Patronymic,
			PhoneNumber: customerCard.PhoneNumber,
			City:        customerCard.City,
			Street:      customerCard.Street,
			ZipCode:     customerCard.ZipCode,
			Percent:     customerCard.Percent,
		}

		c.JSON(http.StatusOK, resp)
	}
}

func NewCustomerCardListsGETHandler(service customerCardReader) gin.HandlerFunc {
	type responseItem struct {
		CardNumber  string `json:"card_number"`
		Surname     string `json:"cust_surname"`
		Name        string `json:"cust_name"`
		Patronymic  string `json:"cust_patronymic"`
		PhoneNumber string `json:"phone_number"`
		City        string `json:"city"`
		Street      string `json:"street"`
		ZipCode     string `json:"zip_code"`
		Percent     int    `json:"percent"`
	}

	return func(c *gin.Context) {
		customerCards, err := service.GetCustomerCards()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve customer cards: " + err.Error()})
			return
		}

		var resp []responseItem
		for _, card := range customerCards {
			resp = append(resp, responseItem{
				CardNumber:  card.CardNumber,
				Surname:     card.Surname,
				Name:        card.Name,
				Patronymic:  card.Patronymic,
				PhoneNumber: card.PhoneNumber,
				City:        card.City,
				Street:      card.Street,
				ZipCode:     card.ZipCode,
				Percent:     card.Percent,
			})
		}

		c.JSON(http.StatusOK, resp)
	}
}

type customerCardRemover interface {
	DeleteCustomerCard(cardNumber string) error
}

func NewCustomerCardDeleteDELETEHandler(service customerCardRemover) gin.HandlerFunc {
	return func(c *gin.Context) {
		cardNumber := c.Param("card_number")
		if len(cardNumber) != 13 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid card number"})
			return
		}

		err := service.DeleteCustomerCard(cardNumber)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete customer card: " + err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Customer card deleted successfully"})
	}
}

type customerCardUpdater interface {
	UpdateCustomerCard(cardNumber string, c models.CustomerCardUpdate) error
}

func NewCustomerCardUpdatePATCHHandler(service customerCardUpdater) gin.HandlerFunc {
	return func(c *gin.Context) {
		cardNumber := c.Param("card_number")
		if len(cardNumber) != 13 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid card number"})
			return
		}

		type request struct {
			Surname     string `json:"cust_surname" binding:"required"`
			Name        string `json:"cust_name" binding:"required"`
			Patronymic  string `json:"cust_patronymic"`
			PhoneNumber string `json:"phone_number" binding:"required"`
			City        string `json:"city"`
			Street      string `json:"street"`
			ZipCode     string `json:"zip_code"`
			Percent     int    `json:"percent" binding:"required"`
		}
		var req request
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: " + err.Error()})
			return
		}

		model := models.CustomerCardUpdate{
			Surname:     req.Surname,
			Name:        req.Name,
			Patronymic:  req.Patronymic,
			PhoneNumber: req.PhoneNumber,
			City:        req.City,
			Street:      req.Street,
			ZipCode:     req.ZipCode,
			Percent:     req.Percent,
		}

		err := service.UpdateCustomerCard(cardNumber, model)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update customer card: " + err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Customer card updated successfully"})
	}
}
