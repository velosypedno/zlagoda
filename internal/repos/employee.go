package repos

import (
	"database/sql"
	"fmt"

	"github.com/velosypedno/zlagoda/internal/models"
	"github.com/velosypedno/zlagoda/internal/utils"
)

func getNewEmployeeId(r *EmployeeRepo) (string, error) {
	const maxRetries = 10

	for i := 0; i < maxRetries; i++ {
		employeeId, err := utils.GenerateID(10)
		if err != nil {
			return "", fmt.Errorf("failed to generate employee ID: %w", err)
		}

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

type EmployeeRepo struct {
	db *sql.DB
}

func NewEmployeeRepo(db *sql.DB) *EmployeeRepo {
	return &EmployeeRepo{
		db: db,
	}
}

func (r *EmployeeRepo) CreateEmployee(c models.EmployeeCreate) (string, error) {
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

	id, err := getNewEmployeeId(r)
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

func (r *EmployeeRepo) RetrieveEmployeeById(id string) (models.EmployeeRetrieve, error) {
	query := `
		SELECT
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
		FROM employee
		WHERE employee_id = $1
	`
	var employee models.EmployeeRetrieve
	err := r.db.QueryRow(query, id).Scan(
		&employee.ID,
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

func (r *EmployeeRepo) RetrieveEmployees() ([]models.EmployeeRetrieve, error) {
	query := `
		SELECT
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

func (r *EmployeeRepo) DeleteEmployee(id string) error {
	query := `DELETE FROM employee WHERE employee_id = $1`
	_, err := r.db.Exec(query, id)
	return err
}

func (r *EmployeeRepo) UpdateEmployee(id string, c models.EmployeeUpdate) error {
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
