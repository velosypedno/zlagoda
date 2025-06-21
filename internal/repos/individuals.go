package repos

import (
	"database/sql"

	"github.com/velosypedno/zlagoda/internal/models"
)

type IndividualsRepo struct {
	db *sql.DB
}

func NewIndividualsRepo(db *sql.DB) *IndividualsRepo {
	return &IndividualsRepo{
		db: db,
	}
}

func (r *IndividualsRepo) QueryVlad1() ([]models.EmployeeRetrieve, error) {
	query := `
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

func (r *IndividualsRepo) QueryVlad2() ([]models.EmployeeRetrieve, error) {
	query := `
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

func (r *IndividualsRepo) QueryArthur1() ([]models.EmployeeRetrieve, error) {
	query := `
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

func (r *IndividualsRepo) QueryArthur2() ([]models.EmployeeRetrieve, error) {
	query := `
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

func (r *IndividualsRepo) QueryOleksii1() ([]models.EmployeeRetrieve, error) {
	query := `
	SELECT 
		e.employee_id,
		e.empl_surname,
		e.empl_name,
		COUNT(DISTINCT cc.card_number) as high_discount_customers,
		COUNT(DISTINCT r.receipt_number) as total_receipts_high_discount,
		SUM(r.sum_total) as total_revenue_high_discount,
		AVG(r.sum_total) as avg_receipt_amount,
		AVG(cc.percent) as avg_customer_discount
	FROM employee e
	JOIN receipt r ON e.employee_id = r.employee_id
	JOIN customer_card cc ON r.card_number = cc.card_number
	WHERE cc.percent > :discount_threshold
		AND e.empl_role = 'cashier'
	GROUP BY e.employee_id, e.empl_surname, e.empl_name
	HAVING COUNT(DISTINCT cc.card_number) > 0
	ORDER BY high_discount_customers DESC, 
			total_revenue_high_discount DESC,
			e.empl_surname ASC;
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

func (r *IndividualsRepo) QueryOleksii2() ([]models.EmployeeRetrieve, error) {
	query := `
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
