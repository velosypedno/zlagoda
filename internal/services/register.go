package services

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/velosypedno/zlagoda/internal/config"
	"github.com/velosypedno/zlagoda/internal/models"
	"github.com/velosypedno/zlagoda/internal/repos"
	"golang.org/x/crypto/bcrypt"
)

func validate(e models.EmployeeCreate) error {
	// required fields
	if e.Surname == nil || *e.Surname == "" {
		return errors.New("surname is required")
	}
	if e.Name == nil || *e.Name == "" {
		return errors.New("name is required")
	}
	if e.Role == nil || *e.Role == "" {
		return errors.New("role is required")
	}
	if e.City == nil || *e.City == "" {
		return errors.New("city is required")
	}
	if e.Street == nil || *e.Street == "" {
		return errors.New("street is required")
	}
	if e.ZipCode == nil || *e.ZipCode == "" {
		return errors.New("zip_code is required")
	}

	// salary ≥ 0
	if e.Salary == nil || *e.Salary < 0 {
		return errors.New("salary must be non-negative")
	}

	// phone format +380XXXXXXXXX
	if e.PhoneNumber == nil {
		return errors.New("phone_number is required")
	}
	pn := *e.PhoneNumber
	if len(pn) != 13 || !strings.HasPrefix(pn, "+380") {
		return fmt.Errorf("phone_number %q must be 13 chars and start with \"+380\"", pn)
	}

	// dates presence
	if e.DateOfBirth == nil {
		return errors.New("date_of_birth is required")
	}
	if e.DateOfStart == nil {
		return errors.New("date_of_start is required")
	}

	bd := *e.DateOfBirth
	sd := *e.DateOfStart
	now := time.Now()

	// must be ≥ 18 yo
	if now.Before(bd.AddDate(18, 0, 0)) {
		return errors.New("employee must be at least 18 years old")
	}
	// birth must come before start (already covers age check, but you can add explicitly)
	if bd.After(sd) {
		return errors.New("date_of_birth must be before date_of_start")
	}

	return nil
}

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
	if err := validate(employee); err != nil {
		return "", fmt.Errorf("validation failed: %w", err)
	}

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
