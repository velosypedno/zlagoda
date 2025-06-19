package services

import "github.com/velosypedno/zlagoda/internal/models"

type SaleRepo interface {
	CreateSale(s models.SaleCreate) error
	RetrieveSaleByKey(upc, receiptNumber string) (models.SaleRetrieve, error)
	RetrieveSalesByReceipt(receiptNumber string) ([]models.SaleRetrieve, error)
	RetrieveSalesByUPC(upc string) ([]models.SaleRetrieve, error)
	RetrieveAllSales() ([]models.SaleRetrieve, error)
	RetrieveSalesWithDetails() ([]models.SaleWithDetails, error)
	RetrieveSalesWithDetailsByReceipt(receiptNumber string) ([]models.SaleWithDetails, error)
	UpdateSale(upc, receiptNumber string, s models.SaleUpdate) error
	DeleteSale(upc, receiptNumber string) error
	DeleteSalesByReceipt(receiptNumber string) error
	GetReceiptTotal(receiptNumber string) (float64, error)
	GetSalesStatsByProduct(productID int, startDate, endDate string) (int, float64, error)
	GetTopSellingProducts(limit int) ([]struct {
		ProductID    int     `json:"product_id"`
		ProductName  string  `json:"product_name"`
		TotalSold    int     `json:"total_sold"`
		TotalRevenue float64 `json:"total_revenue"`
	}, error)
}

type SaleService struct {
	repo SaleRepo
}

func NewSaleService(repo SaleRepo) *SaleService {
	return &SaleService{repo: repo}
}

func (s *SaleService) CreateSale(sale models.SaleCreate) error {
	return s.repo.CreateSale(sale)
}

func (s *SaleService) GetSaleByKey(upc, receiptNumber string) (models.SaleRetrieve, error) {
	return s.repo.RetrieveSaleByKey(upc, receiptNumber)
}

func (s *SaleService) GetSalesByReceipt(receiptNumber string) ([]models.SaleRetrieve, error) {
	return s.repo.RetrieveSalesByReceipt(receiptNumber)
}

func (s *SaleService) GetSalesByUPC(upc string) ([]models.SaleRetrieve, error) {
	return s.repo.RetrieveSalesByUPC(upc)
}

func (s *SaleService) GetAllSales() ([]models.SaleRetrieve, error) {
	return s.repo.RetrieveAllSales()
}

func (s *SaleService) GetSalesWithDetails() ([]models.SaleWithDetails, error) {
	return s.repo.RetrieveSalesWithDetails()
}

func (s *SaleService) GetSalesWithDetailsByReceipt(receiptNumber string) ([]models.SaleWithDetails, error) {
	return s.repo.RetrieveSalesWithDetailsByReceipt(receiptNumber)
}

func (s *SaleService) UpdateSale(upc, receiptNumber string, sale models.SaleUpdate) error {
	return s.repo.UpdateSale(upc, receiptNumber, sale)
}

func (s *SaleService) DeleteSale(upc, receiptNumber string) error {
	return s.repo.DeleteSale(upc, receiptNumber)
}

func (s *SaleService) DeleteSalesByReceipt(receiptNumber string) error {
	return s.repo.DeleteSalesByReceipt(receiptNumber)
}

func (s *SaleService) GetReceiptTotal(receiptNumber string) (float64, error) {
	return s.repo.GetReceiptTotal(receiptNumber)
}

func (s *SaleService) GetSalesStatsByProduct(productID int, startDate, endDate string) (int, float64, error) {
	return s.repo.GetSalesStatsByProduct(productID, startDate, endDate)
}

func (s *SaleService) GetTopSellingProducts(limit int) ([]struct {
	ProductID    int     `json:"product_id"`
	ProductName  string  `json:"product_name"`
	TotalSold    int     `json:"total_sold"`
	TotalRevenue float64 `json:"total_revenue"`
}, error) {
	return s.repo.GetTopSellingProducts(limit)
}
