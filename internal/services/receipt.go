package services

import "github.com/velosypedno/zlagoda/internal/models"

type ReceiptRepo interface {
	CreateReceipt(c models.ReceiptCreate) (string, error)
	RetrieveReceiptByReceiptNumber(receiptNumber string) (models.ReceiptRetrieve, error)
	RetrieveReceipts() ([]models.ReceiptRetrieve, error)
	DeleteReceipt(receiptNumber string) error
	UpdateReceipt(receiptNumber string, c models.ReceiptUpdate) error
}

type ReceiptService struct {
	repo ReceiptRepo
}

func NewReceiptService(repo ReceiptRepo) *ReceiptService {
	return &ReceiptService{repo: repo}
}

func (s *ReceiptService) CreateReceipt(c models.ReceiptCreate) (string, error) {
	return s.repo.CreateReceipt(c)
}

func (s *ReceiptService) GetReceiptByReceiptNumber(receiptNumber string) (models.ReceiptRetrieve, error) {
	return s.repo.RetrieveReceiptByReceiptNumber(receiptNumber)
}

func (s *ReceiptService) GetReceipts() ([]models.ReceiptRetrieve, error) {
	return s.repo.RetrieveReceipts()
}

func (s *ReceiptService) DeleteReceipt(receiptNumber string) error {
	return s.repo.DeleteReceipt(receiptNumber)
}

func (s *ReceiptService) UpdateReceipt(receiptNumber string, c models.ReceiptUpdate) error {
	return s.repo.UpdateReceipt(receiptNumber, c)
}
