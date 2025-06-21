package services

import (
	"log"
	"github.com/velosypedno/zlagoda/internal/models"
)

type EmployeeRepo interface {
	CreateEmployee(c models.EmployeeCreate) (string, error)
	CreateEmployeeWithAuth(c models.EmployeeCreate, login string, hashedPassword string) (string, error)
	RetrieveEmployeeById(id string) (models.EmployeeRetrieve, error)
	RetrieveEmployees() ([]models.EmployeeRetrieve, error)
	DeleteEmployee(id string) error
	UpdateEmployee(id string, c models.EmployeeUpdate) error
}

type EmployeeService struct {
	repo EmployeeRepo
}

func NewEmployeeService(repo EmployeeRepo) *EmployeeService {
	return &EmployeeService{repo: repo}
}

func (s *EmployeeService) CreateEmployee(c models.EmployeeCreate) (string, error) {
	return s.repo.CreateEmployee(c)
}

func (s *EmployeeService) CreateEmployeeWithAuth(c models.EmployeeCreate, login string, hashedPassword string) (string, error) {
	return s.repo.CreateEmployeeWithAuth(c, login, hashedPassword)
}

func (s *EmployeeService) GetEmployeeById(id string) (models.EmployeeRetrieve, error) {
	return s.repo.RetrieveEmployeeById(id)
}

func (s *EmployeeService) GetEmployees() ([]models.EmployeeRetrieve, error) {
	return s.repo.RetrieveEmployees()
}

func (s *EmployeeService) DeleteEmployee(id string) error {
	log.Printf("[EmployeeService] Deleting employee with ID: %s", id)
	err := s.repo.DeleteEmployee(id)
	if err != nil {
		log.Printf("[EmployeeService] Repository error: %v", err)
		return err
	}
	log.Printf("[EmployeeService] Successfully deleted employee with ID: %s", id)
	return nil
}

func (s *EmployeeService) UpdateEmployee(id string, c models.EmployeeUpdate) error {
	return s.repo.UpdateEmployee(id, c)
}
