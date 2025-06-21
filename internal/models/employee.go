package models

import "time"

type EmployeeCreate struct {
	Surname     *string
	Name        *string
	Patronymic  *string
	Role        *string
	Salary      *float64
	DateOfBirth *time.Time
	DateOfStart *time.Time
	PhoneNumber *string
	City        *string
	Street      *string
	ZipCode     *string
}

type EmployeeRetrieve struct {
	ID          *string    `json:"employee_id"`
	Login       *string    `json:"login"`
	Surname     *string    `json:"empl_surname"`
	Name        *string    `json:"empl_name"`
	Patronymic  *string    `json:"empl_patronymic"`
	Role        *string    `json:"empl_role"`
	Salary      *float64   `json:"salary"`
	DateOfBirth *time.Time `json:"date_of_birth"`
	DateOfStart *time.Time `json:"date_of_start"`
	PhoneNumber *string    `json:"phone_number"`
	City        *string    `json:"city"`
	Street      *string    `json:"street"`
	ZipCode     *string    `json:"zip_code"`
}

type EmployeeAuth struct {
	ID       *string
	Password *string
}

type EmployeeUpdate struct {
	Surname     *string
	Name        *string
	Patronymic  *string
	Role        *string
	Salary      *float64
	DateOfBirth *time.Time
	DateOfStart *time.Time
	PhoneNumber *string
	City        *string
	Street      *string
	ZipCode     *string
}
