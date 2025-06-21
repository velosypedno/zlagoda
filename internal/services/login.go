package services

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/velosypedno/zlagoda/internal/config"
	"github.com/velosypedno/zlagoda/internal/repos"
	"golang.org/x/crypto/bcrypt"
)

type LoginService interface {
	Login(ctx context.Context, login string, password string) (string, error)
}

type loginService struct {
	employeeRepo repos.EmployeeRepo
	cfg *config.Config
}

func NewLoginService(employeeRepo repos.EmployeeRepo, cfg *config.Config) LoginService {
	return &loginService{
		employeeRepo: employeeRepo,
		cfg:          cfg,
	}
}

func (s *loginService) Login(ctx context.Context, login string, password string) (string, error) {
	employee, err := s.employeeRepo.GetByLogin(login)
	if err != nil {
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(*employee.Password), []byte(password)); err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"employee_id": *employee.ID,
		"exp":         time.Now().Add(time.Hour * 24).Unix(),
	})

	return token.SignedString([]byte(s.cfg.SECRET_KEY))
}
