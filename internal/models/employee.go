package models

import "time"

type EmployeeCreate struct {
	Surname     string
	Name        string
	Patronymic  string
	Role        string
	Salary      float64
	DateOfBirth time.Time
	DateOfStart time.Time
	PhoneNumber string
	City        string
	Street      string
	ZipCode     string
}

type EmployeeRetrieve struct {
	ID          string
	Surname     string
	Name        string
	Patronymic  string
	Role        string
	Salary      float64
	DateOfBirth time.Time
	DateOfStart time.Time
	PhoneNumber string
	City        string
	Street      string
	ZipCode     string
}

type EmployeeUpdate struct {
	Surname     string
	Name        string
	Patronymic  string
	Role        string
	Salary      float64
	DateOfBirth time.Time
	DateOfStart time.Time
	PhoneNumber string
	City        string
	Street      string
	ZipCode     string
}
