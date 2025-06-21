package services

import "github.com/velosypedno/zlagoda/internal/models"

type IndividualsRepo interface {
	QueryVlad1(categoryID int, months int) ([]models.Vlad1Response, error)
	QueryVlad2() ([]models.Vlad2Response, error)
	QueryArthur1(startDate, endDate string) ([]models.Arthur1Response, error)
	QueryArthur2() ([]models.Arthur2Response, error)
	QueryOleksii1(discountThreshold int) ([]models.Oleksii1Response, error)
	QueryOleksii2() ([]models.Oleksii2Response, error)
}

type IndividualsService struct {
	repo IndividualsRepo
}

func NewIndividualsService(repo IndividualsRepo) *IndividualsService {
	return &IndividualsService{repo: repo}
}

// QueryVlad1 - Get most sold product in a category within a time period
func (s *IndividualsService) QueryVlad1(categoryID int, months int) ([]models.Vlad1Response, error) {
	return s.repo.QueryVlad1(categoryID, months)
}

// QueryVlad2 - Get employees who never sold promotional products
func (s *IndividualsService) QueryVlad2() ([]models.Vlad2Response, error) {
	return s.repo.QueryVlad2()
}

// QueryArthur1 - Get category sales statistics within date range
func (s *IndividualsService) QueryArthur1(startDate, endDate string) ([]models.Arthur1Response, error) {
	return s.repo.QueryArthur1(startDate, endDate)
}

// QueryArthur2 - Get products in store that have never been sold and are not promotional
func (s *IndividualsService) QueryArthur2() ([]models.Arthur2Response, error) {
	return s.repo.QueryArthur2()
}

// QueryOleksii1 - Get cashiers who served customers with high discount
func (s *IndividualsService) QueryOleksii1(discountThreshold int) ([]models.Oleksii1Response, error) {
	return s.repo.QueryOleksii1(discountThreshold)
}

// QueryOleksii2 - Get customers who bought from all categories in the last month
func (s *IndividualsService) QueryOleksii2() ([]models.Oleksii2Response, error) {
	return s.repo.QueryOleksii2()
}
