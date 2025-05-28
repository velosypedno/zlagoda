package services

import "github.com/velosypedno/zlagoda/internal/models"

type CategoryRepo interface {
	CreateCategory(c models.CategoryCreate) error
	RetrieveCategoryByID(id int) (models.CategoryRetrieve, error)
}

type CategoryService struct {
	repo CategoryRepo
}

func NewCategoryService(repo CategoryRepo) *CategoryService {
	return &CategoryService{repo: repo}
}

func (s *CategoryService) CreateCategory(c models.CategoryCreate) error {
	return s.repo.CreateCategory(c)
}

func (s *CategoryService) GetCategoryByID(id int) (models.CategoryRetrieve, error) {
	return s.repo.RetrieveCategoryByID(id)
}
