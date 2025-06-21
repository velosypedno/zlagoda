package repos

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/velosypedno/zlagoda/internal/models"
	"github.com/velosypedno/zlagoda/internal/utils"
)

func (r *employeeRepo) getNewEmployeeId() (string, error) {
	const maxRetries = 10

	for i := 0; i < maxRetries; i++ {
		employeeId, err := utils.GenerateSecureID(10)
		if err != nil {
			return "", fmt.Errorf("failed to generate employee ID: %w", err)
		}

		// Check if employee ID already exists
		var exists bool
		err = r.db.QueryRow("SELECT EXISTS(SELECT 1 FROM employee WHERE employee_id = $1)", employeeId).Scan(&exists)
		if err != nil {
			return "", fmt.Errorf("failed to check employee ID uniqueness: %w", err)
		}

		if !exists {
			return employeeId, nil
		}
	}

	return "", fmt.Errorf("failed to generate unique employee ID after %d attempts", maxRetries)
}

type EmployeeRepo interface {
	CreateEmployee(c models.EmployeeCreate) (string, error)
	CreateEmployeeWithAuth(c models.EmployeeCreate, login string, hashedPassword string) (string, error)
	RetrieveEmployeeById(id string) (models.EmployeeRetrieve, error)
	RetrieveEmployees() ([]models.EmployeeRetrieve, error)
	DeleteEmployee(id string) error
	UpdateEmployee(id string, c models.EmployeeUpdate) error
	GetByLogin(login string) (models.EmployeeAuth, error)
}

type employeeRepo struct {
	db *sql.DB
}

func NewEmployeeRepo(db *sql.DB) EmployeeRepo {
	return &employeeRepo{
		db: db,
	}
}

func (r *employeeRepo) CreateEmployee(c models.EmployeeCreate) (string, error) {
	query := `
		INSERT INTO employee (
			employee_id,
			empl_surname,
			empl_name,
			empl_patronymic,
			empl_role,
			salary,
			date_of_birth,
			date_of_start,
			phone_number,
			city,
			street,
			zip_code
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
		RETURNING employee_id
	`

	id, err := r.getNewEmployeeId()
	if err != nil {
		return "", err
	}
	err = r.db.QueryRow(
		query,
		id,
		c.Surname,
		c.Name,
		c.Patronymic,
		c.Role,
		c.Salary,
		c.DateOfBirth,
		c.DateOfStart,
		c.PhoneNumber,
		c.City,
		c.Street,
		c.ZipCode,
	).Scan(&id)

	return id, err
}

func (r *employeeRepo) CreateEmployeeWithAuth(c models.EmployeeCreate, login string, hashedPassword string) (string, error) {
	query := `
		INSERT INTO employee (
			employee_id,
			login,
			hashed_password,
			empl_surname,
			empl_name,
			empl_patronymic,
			empl_role,
			salary,
			date_of_birth,
			date_of_start,
			phone_number,
			city,
			street,
			zip_code
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)
		RETURNING employee_id
	`

	id, err := r.getNewEmployeeId()
	if err != nil {
		return "", err
	}
	err = r.db.QueryRow(
		query,
		id,
		login,
		hashedPassword,
		c.Surname,
		c.Name,
		c.Patronymic,
		c.Role,
		c.Salary,
		c.DateOfBirth,
		c.DateOfStart,
		c.PhoneNumber,
		c.City,
		c.Street,
		c.ZipCode,
	).Scan(&id)

	return id, err
}

func (r *employeeRepo) RetrieveEmployeeById(id string) (models.EmployeeRetrieve, error) {
	query := `
		SELECT
			employee_id,
			login,
			empl_surname,
			empl_name,
			empl_patronymic,
			empl_role,
			salary,
			date_of_birth,
			date_of_start,
			phone_number,
			city,
			street,
			zip_code
		FROM employee
		WHERE employee_id = $1
	`
	var employee models.EmployeeRetrieve
	err := r.db.QueryRow(query, id).Scan(
		&employee.ID,
		&employee.Login,
		&employee.Surname,
		&employee.Name,
		&employee.Patronymic,
		&employee.Role,
		&employee.Salary,
		&employee.DateOfBirth,
		&employee.DateOfStart,
		&employee.PhoneNumber,
		&employee.City,
		&employee.Street,
		&employee.ZipCode,
	)
	if err != nil {
		return models.EmployeeRetrieve{}, err
	}

	return employee, nil
}

func (r *employeeRepo) RetrieveEmployees() ([]models.EmployeeRetrieve, error) {
	query := `
		SELECT
			employee_id,
			login,
			empl_surname,
			empl_name,
			empl_patronymic,
			empl_role,
			salary,
			date_of_birth,
			date_of_start,
			phone_number,
			city,
			street,
			zip_code
		FROM employee
	`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var employees []models.EmployeeRetrieve
	for rows.Next() {
		var employee models.EmployeeRetrieve
		err := rows.Scan(
			&employee.ID,
			&employee.Login,
			&employee.Surname,
			&employee.Name,
			&employee.Patronymic,
			&employee.Role,
			&employee.Salary,
			&employee.DateOfBirth,
			&employee.DateOfStart,
			&employee.PhoneNumber,
			&employee.City,
			&employee.Street,
			&employee.ZipCode,
		)
		if err != nil {
			return nil, err
		}
		employees = append(employees, employee)
	}

	return employees, nil
}

func (r *employeeRepo) DeleteEmployee(id string) error {
	log.Printf("[EmployeeRepo] Executing DELETE query for employee ID: %s", id)
	query := `DELETE FROM employee WHERE employee_id = $1`
	result, err := r.db.Exec(query, id)
	if err != nil {
		log.Printf("[EmployeeRepo] Database error: %v", err)
		return err
	}
	
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("[EmployeeRepo] Error getting rows affected: %v", err)
		return err
	}
	
	if rowsAffected == 0 {
		log.Printf("[EmployeeRepo] No employee found with ID: %s", id)
		return fmt.Errorf("employee not found")
	}
	
	log.Printf("[EmployeeRepo] Successfully deleted employee with ID: %s, rows affected: %d", id, rowsAffected)
	return nil
}

func (r *employeeRepo) UpdateEmployee(id string, c models.EmployeeUpdate) error {
	query := `
		UPDATE employee
		SET
			empl_surname = $2,
			empl_name = $3,
			empl_patronymic = $4,
			empl_role = $5,
			salary = $6,
			date_of_birth = $7,
			date_of_start = $8,
			phone_number = $9,
			city = $10,
			street = $11,
			zip_code = $12
		WHERE employee_id = $1
	`
	_, err := r.db.Exec(
		query,
		id,
		c.Surname,
		c.Name,
		c.Patronymic,
		c.Role,
		c.Salary,
		c.DateOfBirth,
		c.DateOfStart,
		c.PhoneNumber,
		c.City,
		c.Street,
		c.ZipCode,
	)
	return err
}

func (r *employeeRepo) GetByLogin(login string) (models.EmployeeAuth, error) {
	query := `
		SELECT
			employee_id,
			hashed_password
		FROM employee
		WHERE login = $1
	`
	var employee models.EmployeeAuth
	err := r.db.QueryRow(query, login).Scan(
		&employee.ID,
		&employee.Password,
	)
	if err != nil {
		return models.EmployeeAuth{}, err
	}

	return employee, nil
}