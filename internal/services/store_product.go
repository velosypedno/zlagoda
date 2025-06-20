package services

import (
	"fmt"

	"github.com/velosypedno/zlagoda/internal/models"
)

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
	UpdateProductDelivery(upc string, quantityChange int, newPrice *float64) error
}

type StoreProductService struct {
	repo StoreProductRepo
}

func NewStoreProductService(repo StoreProductRepo) *StoreProductService {
	return &StoreProductService{repo: repo}
}

func (s *StoreProductService) CreateStoreProduct(sp models.StoreProductCreate) (string, error) {
	// if promotional product is set, checks whether it is valid
	if sp.UPCProm != nil {
		promotionalProduct, err := s.repo.RetrieveStoreProductByUPC(*sp.UPCProm)
		if err != nil {
			return "", err
		}
		if promotionalProduct.ProductsNumber != sp.ProductsNumber {
			return "", fmt.Errorf("promotional product's number of units does not equal nonpromotial's one")
		}
		if promotionalProduct.ProductID != sp.ProductID {
			return "", fmt.Errorf("promotional product's product ID does not equal nonpromotial's one")
		}
	}

	// supply handling
	if !sp.PromotionalProduct {
		storeProductsWithSameProductID, err := s.repo.RetrieveStoreProductsByProductID(sp.ProductID)
		if err != nil {
			return "", err
		}
		for _, storeProduct := range storeProductsWithSameProductID {
			var newProductsNumber int = storeProduct.ProductsNumber + sp.ProductsNumber
			updated := models.StoreProductUpdate{
				UPCProm:            storeProduct.UPCProm,
				ProductID:          &storeProduct.ProductID,
				ProductsNumber:     &newProductsNumber,
				PromotionalProduct: &storeProduct.PromotionalProduct,
			}
			if !storeProduct.PromotionalProduct {
				updated.SellingPrice = &sp.SellingPrice
			} else {
				var promotionalSellingPrice float64 = 0.8 * sp.SellingPrice
				updated.SellingPrice = &promotionalSellingPrice
			}
			err = s.repo.UpdateStoreProduct(storeProduct.UPC, updated)
			if err != nil {
				return "", err
			}
		}
	}
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
	storeProductCorrentState, err := s.repo.RetrieveStoreProductByUPC(upc)
	if err != nil {
		return err
	}

	var productId int = storeProductCorrentState.ProductID
	if sp.ProductID != nil {
		productId = *sp.ProductID
	}
	storeProductsWithSameProductID, err := s.repo.RetrieveStoreProductsByProductID(productId)
	if err != nil {
		return err
	}

	var newProductsNumber int = storeProductCorrentState.ProductsNumber
	if sp.ProductsNumber != nil {
		newProductsNumber = *sp.ProductsNumber
	}

	var newSellingPrice float64
	if sp.SellingPrice != nil {
		newSellingPrice = *sp.SellingPrice
	}

	for _, storeProduct := range storeProductsWithSameProductID {
		if sp.PromotionalProduct != nil && !*sp.PromotionalProduct {
			if storeProduct.UPCProm != nil {
				if *storeProduct.UPCProm == storeProductCorrentState.UPC {
					return fmt.Errorf("can not make promotional product to be nonpromotial")
				}
			}
		}

		updated := models.StoreProductUpdate{
			UPCProm:            storeProduct.UPCProm,
			ProductID:          &storeProduct.ProductID,
			ProductsNumber:     &newProductsNumber,
			PromotionalProduct: &storeProduct.PromotionalProduct,
		}

		if (sp.PromotionalProduct != nil && !*sp.PromotionalProduct) || !storeProductCorrentState.PromotionalProduct {
			if !storeProduct.PromotionalProduct {
				updated.SellingPrice = &newSellingPrice
			} else {
				var promotionalSellingPrice float64 = 0.8 * newSellingPrice
				updated.SellingPrice = &promotionalSellingPrice
			}
		}

		err = s.repo.UpdateStoreProduct(storeProduct.UPC, updated)
		if err != nil {
			return err
		}
	}

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

func (s *StoreProductService) UpdateProductDelivery(upc string, quantityChange int, newPrice *float64) error {
	return s.repo.UpdateProductDelivery(upc, quantityChange, newPrice)
}
