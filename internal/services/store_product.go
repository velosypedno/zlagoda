package services

import "github.com/velosypedno/zlagoda/internal/models"

type StoreProductRepo interface {
	CreateStoreProduct(sp models.StoreProductCreate) (string, error)
	RetrieveStoreProductByUPC(upc string) (models.StoreProductRetrieve, error)
	RetrieveStoreProducts() ([]models.StoreProductRetrieve, error)
	RetrieveStoreProductsWithDetails() ([]models.StoreProductWithDetails, error)
	RetrieveStoreProductsByProductID(productID int) ([]models.StoreProductRetrieve, error)
	RetrievePromotionalProducts() ([]models.StoreProductRetrieve, error)
	UpdateStoreProduct(upc string, sp models.StoreProductUpdate) error
	DeleteStoreProduct(upc string) error
	UpdateProductQuantity(upc string, quantityChange int) error
	CheckStockAvailability(upc string, requiredQuantity int) (bool, error)
	RetrieveStoreProductsByCategory(categoryID int) ([]models.StoreProductWithDetails, error)
	RetrieveStoreProductsByName(name string) ([]models.StoreProductWithDetails, error)
}

type StoreProductService struct {
	repo StoreProductRepo
}

func NewStoreProductService(repo StoreProductRepo) *StoreProductService {
	return &StoreProductService{repo: repo}
}

func (s *StoreProductService) CreateStoreProduct(sp models.StoreProductCreate) (string, error) {
	return s.repo.CreateStoreProduct(sp)
}

func (s *StoreProductService) GetStoreProductByUPC(upc string) (models.StoreProductRetrieve, error) {
	return s.repo.RetrieveStoreProductByUPC(upc)
}

func (s *StoreProductService) GetStoreProducts() ([]models.StoreProductRetrieve, error) {
	return s.repo.RetrieveStoreProducts()
}

func (s *StoreProductService) GetStoreProductsWithDetails() ([]models.StoreProductWithDetails, error) {
	return s.repo.RetrieveStoreProductsWithDetails()
}

func (s *StoreProductService) GetStoreProductsByProductID(productID int) ([]models.StoreProductRetrieve, error) {
	return s.repo.RetrieveStoreProductsByProductID(productID)
}

func (s *StoreProductService) GetPromotionalProducts() ([]models.StoreProductRetrieve, error) {
	return s.repo.RetrievePromotionalProducts()
}

func (s *StoreProductService) UpdateStoreProduct(upc string, sp models.StoreProductUpdate) error {
	return s.repo.UpdateStoreProduct(upc, sp)
}

func (s *StoreProductService) DeleteStoreProduct(upc string) error {
	return s.repo.DeleteStoreProduct(upc)
}

func (s *StoreProductService) UpdateProductQuantity(upc string, quantityChange int) error {
	return s.repo.UpdateProductQuantity(upc, quantityChange)
}

func (s *StoreProductService) CheckStockAvailability(upc string, requiredQuantity int) (bool, error) {
	return s.repo.CheckStockAvailability(upc, requiredQuantity)
}

func (s *StoreProductService) GetStoreProductsByCategory(categoryID int) ([]models.StoreProductWithDetails, error) {
	return s.repo.RetrieveStoreProductsByCategory(categoryID)
}

func (s *StoreProductService) GetStoreProductsByName(name string) ([]models.StoreProductWithDetails, error) {
	return s.repo.RetrieveStoreProductsByName(name)
}
