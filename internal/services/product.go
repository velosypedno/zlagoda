package services

import (
	"github.com/velosypedno/zlagoda/internal/models"
	"github.com/velosypedno/zlagoda/internal/repos"
)

type ProductService struct {
	repo repos.ProductRepo
}

func NewProductService(r repos.ProductRepo) *ProductService {
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

func (s *ProductService) UpdateProduct(id int, p models.ProductUpdate) error {
	return s.repo.UpdateProduct(id, p)
}

func (s *ProductService) DeleteProduct(id int) error {
	return s.repo.DeleteProduct(id)
}
