package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/velosypedno/zlagoda/internal/models"
)

type customerCardCreator interface {
	CreateCustomerCard(c models.CustomerCardCreate) (string, error)
}

func NewCustomerCardCreatePOSTHandler(service customerCardCreator) gin.HandlerFunc {
	return func(c *gin.Context) {
		type request struct {
			Surname     *string `json:"cust_surname" binding:"omitempty,required,max=50"`
			Name        *string `json:"cust_name" binding:"omitempty,required,max=50"`
			Patronymic  *string `json:"cust_patronymic" binding:"omitempty,max=50"`
			PhoneNumber *string `json:"phone_number" binding:"omitempty,required,len=13,startswith=+380"`
			City        *string `json:"city" binding:"omitempty,max=50"`
			Street      *string `json:"street" binding:"omitempty,max=50"`
			ZipCode     *string `json:"zip_code" binding:"omitempty,max=9"`
			Percent     *int    `json:"percent" binding:"omitempty,required,gte=0"`
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
			CardNumber  *string `json:"card_number"`
			Surname     *string `json:"cust_surname"`
			Name        *string `json:"cust_name"`
			Patronymic  *string `json:"cust_patronymic"`
			PhoneNumber *string `json:"phone_number"`
			City        *string `json:"city"`
			Street      *string `json:"street"`
			ZipCode     *string `json:"zip_code"`
			Percent     *int    `json:"percent"`
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

func NewCustomerCardsListGETHandler(service customerCardReader) gin.HandlerFunc {
	type responseItem struct {
		CardNumber  *string `json:"card_number"`
		Surname     *string `json:"cust_surname"`
		Name        *string `json:"cust_name"`
		Patronymic  *string `json:"cust_patronymic"`
		PhoneNumber *string `json:"phone_number"`
		City        *string `json:"city"`
		Street      *string `json:"street"`
		ZipCode     *string `json:"zip_code"`
		Percent     *int    `json:"percent"`
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
	GetCustomerCardByCardNumber(cardNumber string) (models.CustomerCardRetrieve, error)
}

func NewCustomerCardUpdatePATCHHandler(service customerCardUpdater) gin.HandlerFunc {
	return func(c *gin.Context) {
		cardNumber := c.Param("card_number")
		if len(cardNumber) != 13 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid card number"})
			return
		}

		type request struct {
			Surname     *string `json:"cust_surname" binding:"omitempty,max=50"`
			Name        *string `json:"cust_name" binding:"omitempty,max=50"`
			Patronymic  *string `json:"cust_patronymic" binding:"omitempty,max=50"`
			PhoneNumber *string `json:"phone_number" binding:"omitempty,len=13,startswith=+380"`
			City        *string `json:"city" binding:"omitempty,max=50"`
			Street      *string `json:"street" binding:"omitempty,max=50"`
			ZipCode     *string `json:"zip_code" binding:"omitempty,max=9"`
			Percent     *int    `json:"percent" binding:"omitempty,gte=0"`
		}
		var req request
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: " + err.Error()})
			return
		}

		custCardCurrentState, err := service.GetCustomerCardByCardNumber(cardNumber)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Customer card not found: " + err.Error()})
			return
		}
		if req.Surname == nil {
			req.Surname = custCardCurrentState.Surname
		}
		if req.Name == nil {
			req.Name = custCardCurrentState.Name
		}
		if req.Patronymic == nil {
			req.Patronymic = custCardCurrentState.Patronymic
		}
		if req.PhoneNumber == nil {
			req.PhoneNumber = custCardCurrentState.PhoneNumber
		}
		if req.City == nil {
			req.City = custCardCurrentState.City
		}
		if req.Street == nil {
			req.Street = custCardCurrentState.Street
		}
		if req.ZipCode == nil {
			req.ZipCode = custCardCurrentState.ZipCode
		}
		if req.Percent == nil {
			req.Percent = custCardCurrentState.Percent
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

		err = service.UpdateCustomerCard(cardNumber, model)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update customer card: " + err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Customer card updated successfully"})
	}
}
