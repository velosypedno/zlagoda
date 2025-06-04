package services

import "github.com/velosypedno/zlagoda/internal/models"

type CustomerCardRepo interface {
	CreateCustomerCard(c models.CustomerCardCreate) (string, error)
	RetrieveCustomerCardByCardNumber(cardNumber string) (models.CustomerCardRetrieve, error)
	RetrieveCustomerCards() ([]models.CustomerCardRetrieve, error)
	DeleteCustomerCard(cardNumber string) error
	UpdateCustomerCard(cardNumber string, c models.CustomerCardUpdate) error
}

type CustomerCardService struct {
	repo CustomerCardRepo
}

func NewCustomerCardService(repo CustomerCardRepo) *CustomerCardService {
	return &CustomerCardService{repo: repo}
}

func (s *CustomerCardService) CreateCustomerCard(c models.CustomerCardCreate) (string, error) {
	return s.repo.CreateCustomerCard(c)
}

func (s *CustomerCardService) GetCustomerCardByCardNumber(cardNumber string) (models.CustomerCardRetrieve, error) {
	return s.repo.RetrieveCustomerCardByCardNumber(cardNumber)
}

func (s *CustomerCardService) GetCustomerCards() ([]models.CustomerCardRetrieve, error) {
	return s.repo.RetrieveCustomerCards()
}

func (s *CustomerCardService) DeleteCustomerCard(cardNumber string) error {
	return s.repo.DeleteCustomerCard(cardNumber)
}

func (s *CustomerCardService) UpdateCustomerCard(cardNumber string, c models.CustomerCardUpdate) error {
	return s.repo.UpdateCustomerCard(cardNumber, c)
}
