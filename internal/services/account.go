package services

import (
	"context"

	"github.com/velosypedno/zlagoda/internal/models"
	"github.com/velosypedno/zlagoda/internal/repos"
)

type AccountService interface {
	GetAccount(ctx context.Context, employeeID string) (*models.EmployeeRetrieve, error)
}

type accountService struct {
	employeeRepo repos.EmployeeRepo
}

func NewAccountService(employeeRepo repos.EmployeeRepo) AccountService {
	return &accountService{
		employeeRepo: employeeRepo,
	}
}

func (s *accountService) GetAccount(ctx context.Context, employeeID string) (*models.EmployeeRetrieve, error) {
	employee, err := s.employeeRepo.RetrieveEmployeeById(employeeID)
	if err != nil {
		return nil, err
	}

	return &employee, nil
} 