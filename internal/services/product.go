package services

import (
	"github.com/velosypedno/zlagoda/internal/models"
)

type ProductRepo interface {
	CreateProduct(p models.ProductCreate) (int, error)
	RetrieveProductByID(id int) (models.ProductRetrieve, error)
	RetrieveProductsByCategory(categoryID int) ([]models.ProductRetrieve, error)
	RetrieveProductsByName(name string) ([]models.ProductRetrieve, error)
	RetrieveProducts() ([]models.ProductRetrieve, error)
	UpdateProduct(id int, p models.ProductUpdate) error
	DeleteProduct(id int) error
}

type ProductService struct {
	repo ProductRepo
}

func NewProductService(r ProductRepo) *ProductService {
	return &ProductService{repo: r}
}

func (s *ProductService) CreateProduct(p models.ProductCreate) (int, error) {
	return s.repo.CreateProduct(p)
}

func (s *ProductService) GetProductByID(id int) (models.ProductRetrieve, error) {
	return s.repo.RetrieveProductByID(id)
}

func (s *ProductService) GetProducts() ([]models.ProductRetrieve, error) {
	return s.repo.RetrieveProducts()
}

func (s *ProductService) GetProductsByCategory(categoryID int) ([]models.ProductRetrieve, error) {
	return s.repo.RetrieveProductsByCategory(categoryID)
}

func (s *ProductService) GetProductsByName(name string) ([]models.ProductRetrieve, error) {
	return s.repo.RetrieveProductsByName(name)
}

func (s *ProductService) UpdateProduct(id int, p models.ProductUpdate) error {
	return s.repo.UpdateProduct(id, p)
}

func (s *ProductService) DeleteProduct(id int) error {
	return s.repo.DeleteProduct(id)
}
