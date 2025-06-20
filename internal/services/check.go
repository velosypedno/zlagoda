package services

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/velosypedno/zlagoda/internal/models"
)

type CheckRepo interface {
	BeginTx() (*sql.Tx, error)
	CreateReceiptTx(tx *sql.Tx, c models.ReceiptCreate) (string, error)
	CreateSaleTx(tx *sql.Tx, s models.SaleCreate) error
	UpdateProductQuantityTx(tx *sql.Tx, upc string, quantityChange int) error
	GetStoreProductStockTx(tx *sql.Tx, upc string) (int, error)
	RetrieveChecks() ([]models.ReceiptRetrieve, error)
}

type CheckServiceImpl struct {
	repo CheckRepo
}

func NewCheckService(repo CheckRepo) *CheckServiceImpl {
	return &CheckServiceImpl{repo: repo}
}

func (s *CheckServiceImpl) CreateCheck(req models.CheckCreate, vatRate float64) (*models.CheckCreateResponse, error) {
	tx, err := s.repo.BeginTx()
	if err != nil {
		log.Println("[CreateCheck] Failed to begin transaction:", err)
		return nil, err
	}
	defer func() {
		if err != nil {
			log.Println("[CreateCheck] Rolling back transaction due to error:", err)
			tx.Rollback()
		}
	}()
	// Calculate total sum
	totalSum := 0.0
	for _, item := range req.Items {
		if item.ProductNumber < 1 {
			log.Printf("[CreateCheck] Invalid product number for UPC %s: %d", item.UPC, item.ProductNumber)
			return nil, errors.New("Product number must be >= 1")
		}
		stock, err := s.repo.GetStoreProductStockTx(tx, item.UPC)
		if err != nil {
			log.Printf("[CreateCheck] Failed to get stock for UPC %s: %v", item.UPC, err)
			return nil, fmt.Errorf("Failed to get stock for UPC %s: %w", item.UPC, err)
		}
		if stock < item.ProductNumber {
			log.Printf("[CreateCheck] Insufficient stock for UPC %s: have %d, need %d", item.UPC, stock, item.ProductNumber)
			return nil, fmt.Errorf("Insufficient stock for UPC %s", item.UPC)
		}
		totalSum += float64(item.ProductNumber) * item.SellingPrice
	}
	vat := vatRate * totalSum
	printDate, _ := time.Parse("2006-01-02", req.PrintDate)
	receipt := models.ReceiptCreate{
		EmployeeId: &req.EmployeeId,
		CardNumber: req.CardNumber,
		PrintDate:  &printDate,
		TotalSum:   &totalSum,
		VAT:        &vat,
	}
	log.Printf("[CreateCheck] Creating receipt: %+v", receipt)
	receiptNumber, err := s.repo.CreateReceiptTx(tx, receipt)
	if err != nil {
		log.Printf("[CreateCheck] Failed to create receipt: %v", err)
		return nil, fmt.Errorf("Failed to create receipt: %w", err)
	}
	for _, item := range req.Items {
		sale := models.SaleCreate{
			UPC:           item.UPC,
			ReceiptNumber: receiptNumber,
			ProductNumber: item.ProductNumber,
			SellingPrice:  item.SellingPrice,
		}
		log.Printf("[CreateCheck] Creating sale: %+v", sale)
		err = s.repo.CreateSaleTx(tx, sale)
		if err != nil {
			log.Printf("[CreateCheck] Failed to create sale for UPC %s: %v", item.UPC, err)
			return nil, fmt.Errorf("Failed to create sale for UPC %s: %w", item.UPC, err)
		}
		err = s.repo.UpdateProductQuantityTx(tx, item.UPC, -item.ProductNumber)
		if err != nil {
			log.Printf("[CreateCheck] Failed to update stock for UPC %s: %v", item.UPC, err)
			return nil, fmt.Errorf("Failed to update stock for UPC %s: %w", item.UPC, err)
		}
	}
	if err = tx.Commit(); err != nil {
		log.Println("[CreateCheck] Failed to commit transaction:", err)
		return nil, err
	}
	log.Printf("[CreateCheck] Successfully created check: receipt_number=%s, total=%.2f, vat=%.2f", receiptNumber, totalSum, vat)
	return &models.CheckCreateResponse{
		ReceiptNumber: receiptNumber,
		PrintDate:     printDate,
		TotalSum:      totalSum,
		VAT:           vat,
	}, nil
}

func (s *CheckServiceImpl) GetChecks() ([]models.ReceiptRetrieve, error) {
	return s.repo.RetrieveChecks()
}
