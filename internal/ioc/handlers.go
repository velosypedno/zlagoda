package ioc

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/velosypedno/zlagoda/internal/config"
	"github.com/velosypedno/zlagoda/internal/handlers"
	"github.com/velosypedno/zlagoda/internal/repos"
	"github.com/velosypedno/zlagoda/internal/services"
)

type HandlerContainer struct {
	CategoryCreatePOSTHandler   gin.HandlerFunc
	CategoryRetrieveGETHandler  gin.HandlerFunc
	CategoriesListGETHandler    gin.HandlerFunc
	CategoryDeleteDELETEHandler gin.HandlerFunc
	CategoryUpdatePATCHHandler  gin.HandlerFunc

	CustomerCardCreatePOSTHandler   gin.HandlerFunc
	CustomerCardRetrieveGETHandler  gin.HandlerFunc
	CustomerCardsListGETHandler     gin.HandlerFunc
	CustomerCardDeleteDELETEHandler gin.HandlerFunc
	CustomerCardUpdatePATCHHandler  gin.HandlerFunc

	EmployeeCreatePOSTHandler   gin.HandlerFunc
	EmployeeRetrieveGETHandler  gin.HandlerFunc
	EmployeesListGETHandler     gin.HandlerFunc
	EmployeeDeleteDELETEHandler gin.HandlerFunc
	EmployeeUpdatePATCHHandler  gin.HandlerFunc

	ReceiptCreatePOSTHandler   gin.HandlerFunc
	ReceiptRetrieveGETHandler  gin.HandlerFunc
	ReceiptsListGETHandler     gin.HandlerFunc
	ReceiptDeleteDELETEHandler gin.HandlerFunc
	ReceiptUpdatePATCHHandler  gin.HandlerFunc
}

func BuildHandlerContainer(c *config.Config) *HandlerContainer {
	db, err := sql.Open(c.DB_DRIVER, c.DB_DSN)
	if err != nil {
		log.Fatal(err)
	}
	categoryRepo := repos.NewCategoryRepo(db)
	categoryService := services.NewCategoryService(categoryRepo)

	customerCardRepo := repos.NewCustomerCardRepo(db)
	customerCardService := services.NewCustomerCardService(customerCardRepo)

	employeeRepo := repos.NewEmployeeRepo(db)
	employeeService := services.NewEmployeeService(employeeRepo)

	receiptRepo := repos.NewReceiptRepo(db)
	receiptService := services.NewReceiptService(receiptRepo)

	return &HandlerContainer{
		CategoryCreatePOSTHandler:   handlers.NewCategoryCreatePOSTHandler(categoryService),
		CategoryRetrieveGETHandler:  handlers.NewCategoryRetrieveGETHandler(categoryService),
		CategoriesListGETHandler:    handlers.NewCategoryListGETHandler(categoryService),
		CategoryDeleteDELETEHandler: handlers.NewCategoryDeleteDELETEHandler(categoryService),
		CategoryUpdatePATCHHandler:  handlers.NewCategoryUpdatePATCHHandler(categoryService),

		CustomerCardCreatePOSTHandler:   handlers.NewCustomerCardCreatePOSTHandler(customerCardService),
		CustomerCardRetrieveGETHandler:  handlers.NewCustomerCardRetrieveGETHandler(customerCardService),
		CustomerCardsListGETHandler:     handlers.NewCustomerCardsListGETHandler(customerCardService),
		CustomerCardDeleteDELETEHandler: handlers.NewCustomerCardDeleteDELETEHandler(customerCardService),
		CustomerCardUpdatePATCHHandler:  handlers.NewCustomerCardUpdatePATCHHandler(customerCardService),

		EmployeeCreatePOSTHandler:   handlers.NewEmployeeCreatePOSTHandler(employeeService),
		EmployeeRetrieveGETHandler:  handlers.NewEmployeeRetrieveGETHandler(employeeService),
		EmployeesListGETHandler:     handlers.NewEmployeesListGETHandler(employeeService),
		EmployeeDeleteDELETEHandler: handlers.NewEmployeeDeleteDELETEHandler(employeeService),
		EmployeeUpdatePATCHHandler:  handlers.NewEmployeeUpdatePATCHHandler(employeeService),

		ReceiptCreatePOSTHandler:   handlers.NewReceiptCreatePOSTHandler(receiptService),
		ReceiptRetrieveGETHandler:  handlers.NewReceiptRetrieveGETHandler(receiptService),
		ReceiptsListGETHandler:     handlers.NewReceiptsListGETHandler(receiptService),
		ReceiptDeleteDELETEHandler: handlers.NewReceiptDeleteDELETEHandler(receiptService),
		ReceiptUpdatePATCHHandler:  handlers.NewReceiptUpdatePATCHHandler(receiptService),
	}
}
