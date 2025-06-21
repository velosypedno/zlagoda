package services

import (
	"fmt"

	"github.com/velosypedno/zlagoda/internal/models"
)

type ReceiptRepo interface {
	CreateReceipt(c models.ReceiptCreate) (string, error)
	RetrieveReceiptByReceiptNumber(receiptNumber string) (models.ReceiptRetrieve, error)
	RetrieveReceipts() ([]models.ReceiptRetrieve, error)
	DeleteReceipt(receiptNumber string) error
	UpdateReceipt(receiptNumber string, c models.ReceiptUpdate) error
}

type SaleRepoInterface interface {
	CreateSale(s models.SaleCreate) error
}

type StoreProductRepoInterface interface {
	CheckStockAvailability(upc string, requiredQuantity int) (bool, error)
	UpdateProductQuantity(upc string, quantityChange int) error
}

type ReceiptService struct {
	receiptRepo      ReceiptRepo
	saleRepo         SaleRepoInterface
	storeProductRepo StoreProductRepoInterface
}

func NewReceiptService(receiptRepo ReceiptRepo, saleRepo SaleRepoInterface, storeProductRepo StoreProductRepoInterface) *ReceiptService {
	return &ReceiptService{
		receiptRepo:      receiptRepo,
		saleRepo:         saleRepo,
		storeProductRepo: storeProductRepo,
	}
}

func (s *ReceiptService) CreateReceipt(c models.ReceiptCreate) (string, error) {
	return s.receiptRepo.CreateReceipt(c)
}

func (s *ReceiptService) GetReceiptByReceiptNumber(receiptNumber string) (models.ReceiptRetrieve, error) {
	return s.receiptRepo.RetrieveReceiptByReceiptNumber(receiptNumber)
}

func (s *ReceiptService) GetReceipts() ([]models.ReceiptRetrieve, error) {
	return s.receiptRepo.RetrieveReceipts()
}

func (s *ReceiptService) DeleteReceipt(receiptNumber string) error {
	return s.receiptRepo.DeleteReceipt(receiptNumber)
}

func (s *ReceiptService) UpdateReceipt(receiptNumber string, c models.ReceiptUpdate) error {
	return s.receiptRepo.UpdateReceipt(receiptNumber, c)
}

func (s *ReceiptService) CreateReceiptComplete(c models.ReceiptCreateComplete, vatRate float64) (string, error) {
	// Validate stock availability first
	for _, item := range c.Items {
		available, err := s.storeProductRepo.CheckStockAvailability(*item.UPC, *item.ProductNumber)
		if err != nil {
			return "", fmt.Errorf("failed to check stock for UPC %s: %w", *item.UPC, err)
		}
		if !available {
			return "", fmt.Errorf("insufficient stock for UPC %s", *item.UPC)
		}
	}

	// Calculate totals
	var totalSum float64 = 0
	for _, item := range c.Items {
		totalSum += float64(*item.ProductNumber) * *item.SellingPrice
	}
	vat := vatRate * totalSum

	// Create receipt
	receipt := models.ReceiptCreate{
		EmployeeId: c.EmployeeId,
		CardNumber: c.CardNumber,
		PrintDate:  c.PrintDate,
		TotalSum:   &totalSum,
		VAT:        &vat,
	}

	receiptNumber, err := s.receiptRepo.CreateReceipt(receipt)
	if err != nil {
		return "", fmt.Errorf("failed to create receipt: %w", err)
	}

	// Create sales using existing SaleRepo
	for _, item := range c.Items {
		sale := models.SaleCreate{
			UPC:           *item.UPC,
			ReceiptNumber: receiptNumber,
			ProductNumber: *item.ProductNumber,
			SellingPrice:  *item.SellingPrice,
		}

		err = s.saleRepo.CreateSale(sale)
		if err != nil {
			return "", fmt.Errorf("failed to create sale for UPC %s: %w", *item.UPC, err)
		}

		// Update stock using existing StoreProductRepo
		err = s.storeProductRepo.UpdateProductQuantity(*item.UPC, -*item.ProductNumber)
		if err != nil {
			return "", fmt.Errorf("failed to update stock for UPC %s: %w", *item.UPC, err)
		}
	}

	return receiptNumber, nil
}
