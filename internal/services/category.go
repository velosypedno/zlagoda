package services

import "github.com/velosypedno/zlagoda/internal/models"

type CategoryRepo interface {
	CreateCategory(c models.CategoryCreate) (int, error)
	RetrieveCategoryByID(id int) (models.CategoryRetrieve, error)
	RetrieveCategories() ([]models.CategoryRetrieve, error)
}

type CategoryService struct {
	repo CategoryRepo
}

func NewCategoryService(repo CategoryRepo) *CategoryService {
	return &CategoryService{repo: repo}
}

func (s *CategoryService) CreateCategory(c models.CategoryCreate) (int, error) {
	return s.repo.CreateCategory(c)
}

func (s *CategoryService) GetCategoryByID(id int) (models.CategoryRetrieve, error) {
	return s.repo.RetrieveCategoryByID(id)
}

func (s *CategoryService) GetCategories() ([]models.CategoryRetrieve, error) {
	return s.repo.RetrieveCategories()
}
