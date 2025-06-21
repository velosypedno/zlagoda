package services

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/velosypedno/zlagoda/internal/config"
	"github.com/velosypedno/zlagoda/internal/models"
	"github.com/velosypedno/zlagoda/internal/repos"
	"golang.org/x/crypto/bcrypt"
)

type RegisterService interface {
	Register(ctx context.Context, employee models.EmployeeCreate, login string, password string) (string, error)
}

type registerService struct {
	employeeRepo repos.EmployeeRepo
	cfg          *config.Config
}

func NewRegisterService(employeeRepo repos.EmployeeRepo, cfg *config.Config) RegisterService {
	return &registerService{
		employeeRepo: employeeRepo,
		cfg:          cfg,
	}
}

func (s *registerService) Register(ctx context.Context, employee models.EmployeeCreate, login string, password string) (string, error) {
	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	// Create the employee with login and hashed password
	employeeID, err := s.employeeRepo.CreateEmployeeWithAuth(employee, login, string(hashedPassword))
	if err != nil {
		return "", err
	}

	// Generate JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"employee_id": employeeID,
		"exp":         time.Now().Add(time.Hour * 24).Unix(),
	})

	return token.SignedString([]byte(s.cfg.SECRET_KEY))
} 