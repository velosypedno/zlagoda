package handlers

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/velosypedno/zlagoda/internal/models"
)

func isSalaryValid(salary float64) bool {
	var salary_str string = strconv.FormatFloat(salary, 'f', -1, 64)

	var parts []string = strings.Split(salary_str, ".")
	var before int = len(parts[0])
	var after int = 0
	if len(parts) > 1 {
		after = len(parts[1])
	}

	return before <= 13 && after <= 4
}

func isEmployeeUpdateValid(empl models.EmployeeUpdate) bool {
	if len(empl.Surname) > 50 {
		return false
	}
	if len(empl.Name) > 50 {
		return false
	}
	if len(empl.Patronymic) > 50 {
		return false
	}
	if len(empl.Role) > 10 {
		return false
	}
	if !isSalaryValid(empl.Salary) {
		return false
	}

	now := time.Now()
	eighteenYearsOld := empl.DateOfBirth.AddDate(18, 0, 0)
	// if empl.DateOfStart.After(now) {
	// 	return false
	// }
	if eighteenYearsOld.After(now) {
		return false
	}
	if eighteenYearsOld.After(empl.DateOfStart) {
		return false
	}

	if !strings.HasPrefix(empl.PhoneNumber, "+380") || len(empl.PhoneNumber) != 13 {
		return false
	}
	if len(empl.City) > 50 {
		return false
	}
	if len(empl.Street) > 50 {
		return false
	}
	if len(empl.ZipCode) > 9 {
		return false
	}
	return true
}

type employeeCreator interface {
	CreateEmployee(c models.EmployeeCreate) (string, error)
}

func NewEmployeeCreatePOSTHandler(service employeeCreator) gin.HandlerFunc {
	return func(c *gin.Context) {
		type request struct {
			Surname     string  `json:"empl_surname" binding:"required,max=50"`
			Name        string  `json:"empl_name" binding:"required,max=50"`
			Patronymic  string  `json:"empl_patronymic" binding:"max=50"`
			Role        string  `json:"empl_role" binding:"required,max=10"`
			Salary      float64 `json:"salary" binding:"required"`
			DateOfBirth string  `json:"date_of_birth" binding:"required"`
			DateOfStart string  `json:"date_of_start" binding:"required"`
			PhoneNumber string  `json:"phone_number" binding:"required,len=13,startswith=+380"`
			City        string  `json:"city" binding:"max=50"`
			Street      string  `json:"street" binding:"max=50"`
			ZipCode     string  `json:"zip_code" binding:"max=9"`
		}
		var req request
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: " + err.Error()})
			return
		}

		birthDate, err := time.Parse("2006-01-02", req.DateOfBirth)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: invalid date of birth format"})
			return
		}
		startDate, err := time.Parse("2006-01-02", req.DateOfStart)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: invalid date of start format"})
			return
		}
		now := time.Now()
		eighteenYearsOld := birthDate.AddDate(18, 0, 0)
		// if empl.DateOfStart.After(now) {
		// 	return false
		// }
		if eighteenYearsOld.After(now) || eighteenYearsOld.After(startDate) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: invalid dates"})
			return
		}

		if !isSalaryValid(req.Salary) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: invalid salary"})
			return
		}

		model := models.EmployeeCreate{
			Surname:     req.Surname,
			Name:        req.Name,
			Patronymic:  req.Patronymic,
			Role:        req.Role,
			Salary:      req.Salary,
			DateOfBirth: birthDate,
			DateOfStart: startDate,
			PhoneNumber: req.PhoneNumber,
			City:        req.City,
			Street:      req.Street,
			ZipCode:     req.ZipCode,
		}

		id, err := service.CreateEmployee(model)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create employee: " + err.Error()})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"id": id})
	}
}

type EmployeeReader interface {
	GetEmployeeById(id string) (models.EmployeeRetrieve, error)
	GetEmployees() ([]models.EmployeeRetrieve, error)
}

func NewEmployeeRetrieveGETHandler(service EmployeeReader) gin.HandlerFunc {
	return func(c *gin.Context) {
		type response struct {
			ID          string  `json:"employee_id"`
			Surname     string  `json:"empl_surname"`
			Name        string  `json:"empl_name"`
			Patronymic  string  `json:"empl_patronymic"`
			Role        string  `json:"empl_role"`
			Salary      float64 `json:"salary"`
			DateOfBirth string  `json:"date_of_birth"`
			DateOfStart string  `json:"date_of_start"`
			PhoneNumber string  `json:"phone_number"`
			City        string  `json:"city"`
			Street      string  `json:"street"`
			ZipCode     string  `json:"zip_code"`
		}
		var id string = c.Param("id")
		if len(id) != 10 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid employee ID"})
			return
		}

		employee, err := service.GetEmployeeById(id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Employee not found: " + err.Error()})
			return
		}

		resp := response{
			ID:          employee.ID,
			Surname:     employee.Surname,
			Name:        employee.Name,
			Patronymic:  employee.Patronymic,
			Role:        employee.Role,
			Salary:      employee.Salary,
			DateOfBirth: employee.DateOfBirth.Format("2006-01-02"),
			DateOfStart: employee.DateOfStart.Format("2006-01-02"),
			PhoneNumber: employee.PhoneNumber,
			City:        employee.City,
			Street:      employee.Street,
			ZipCode:     employee.ZipCode,
		}

		c.JSON(http.StatusOK, resp)
	}
}

func NewEmployeesListGETHandler(service EmployeeReader) gin.HandlerFunc {
	type responseItem struct {
		ID          string  `json:"employee_id"`
		Surname     string  `json:"empl_surname"`
		Name        string  `json:"empl_name"`
		Patronymic  string  `json:"empl_patronymic"`
		Role        string  `json:"empl_role"`
		Salary      float64 `json:"salary"`
		DateOfBirth string  `json:"date_of_birth"`
		DateOfStart string  `json:"date_of_start"`
		PhoneNumber string  `json:"phone_number"`
		City        string  `json:"city"`
		Street      string  `json:"street"`
		ZipCode     string  `json:"zip_code"`
	}

	return func(c *gin.Context) {
		employees, err := service.GetEmployees()
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Failed to retrieve employees: " + err.Error()})
			return
		}

		var resp []responseItem
		for _, employee := range employees {
			resp = append(resp, responseItem{
				ID:          employee.ID,
				Surname:     employee.Surname,
				Name:        employee.Name,
				Patronymic:  employee.Patronymic,
				Role:        employee.Role,
				Salary:      employee.Salary,
				DateOfBirth: employee.DateOfBirth.Format("2006-01-02"),
				DateOfStart: employee.DateOfStart.Format("2006-01-02"),
				PhoneNumber: employee.PhoneNumber,
				City:        employee.City,
				Street:      employee.Street,
				ZipCode:     employee.ZipCode,
			})
		}

		c.JSON(http.StatusOK, resp)
	}
}

type EmployeeRemover interface {
	DeleteEmployee(id string) error
}

func NewEmployeeDeleteDELETEHandler(service EmployeeRemover) gin.HandlerFunc {
	return func(c *gin.Context) {
		var id string = c.Param("id")
		if len(id) != 10 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid employee ID"})
			return
		}

		err := service.DeleteEmployee(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete employee: " + err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Employee deleted successfully"})
	}
}

type EmployeeUpdater interface {
	UpdateEmployee(id string, c models.EmployeeUpdate) error
	GetEmployeeById(id string) (models.EmployeeRetrieve, error)
}

func NewEmployeeUpdatePATCHHandler(service EmployeeUpdater) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		if len(id) != 10 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid employee ID"})
			return
		}

		type request struct {
			Surname     *string  `json:"empl_surname"`
			Name        *string  `json:"empl_name"`
			Patronymic  *string  `json:"empl_patronymic"`
			Role        *string  `json:"empl_role"`
			Salary      *float64 `json:"salary"`
			DateOfBirth *string  `json:"date_of_birth"`
			DateOfStart *string  `json:"date_of_start"`
			PhoneNumber *string  `json:"phone_number"`
			City        *string  `json:"city"`
			Street      *string  `json:"street"`
			ZipCode     *string  `json:"zip_code"`
		}
		var req request
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: " + err.Error()})
			return
		}

		employeeCurrentState, err := service.GetEmployeeById(id)
		if err != nil {
			employeeCurrentState = models.EmployeeRetrieve{}
		}
		currentBirthDateStr := employeeCurrentState.DateOfBirth.Format("2006-01-02")
		currentStartDateStr := employeeCurrentState.DateOfStart.Format("2006-01-02")
		if req.Surname == nil {
			req.Surname = &employeeCurrentState.Surname
		}
		if req.Name == nil {
			req.Name = &employeeCurrentState.Name
		}
		if req.Patronymic == nil {
			req.Patronymic = &employeeCurrentState.Patronymic
		}
		if req.Role == nil {
			req.Role = &employeeCurrentState.Role
		}
		if req.Salary == nil {
			req.Salary = &employeeCurrentState.Salary
		}
		if req.DateOfBirth == nil {
			req.DateOfBirth = &currentBirthDateStr
		}
		if req.DateOfStart == nil {
			req.DateOfStart = &currentStartDateStr
		}
		if req.PhoneNumber == nil {
			req.PhoneNumber = &employeeCurrentState.PhoneNumber
		}
		if req.City == nil {
			req.City = &employeeCurrentState.City
		}
		if req.Street == nil {
			req.Street = &employeeCurrentState.Street
		}
		if req.ZipCode == nil {
			req.ZipCode = &employeeCurrentState.ZipCode
		}

		birthDate, err := time.Parse("2006-01-02", *req.DateOfBirth)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: invalid date of birth format"})
			return
		}
		startDate, err := time.Parse("2006-01-02", *req.DateOfStart)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: invalid date of start format"})
			return
		}

		model := models.EmployeeUpdate{
			Surname:     *req.Surname,
			Name:        *req.Name,
			Patronymic:  *req.Patronymic,
			Role:        *req.Role,
			Salary:      *req.Salary,
			DateOfBirth: birthDate,
			DateOfStart: startDate,
			PhoneNumber: *req.PhoneNumber,
			City:        *req.City,
			Street:      *req.Street,
			ZipCode:     *req.ZipCode,
		}
		if !isEmployeeUpdateValid(model) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: input values are out of bounds"})
			return
		}

		err = service.UpdateEmployee(id, model)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update employee: " + err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Employee updated successfully"})
	}
}
