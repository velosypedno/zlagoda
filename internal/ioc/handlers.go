package ioc

import (
	"database/sql"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/velosypedno/zlagoda/internal/config"
	"github.com/velosypedno/zlagoda/internal/handlers"
	"github.com/velosypedno/zlagoda/internal/repos"
	"github.com/velosypedno/zlagoda/internal/services"
)

type HandlerContainer struct {
	DB *sql.DB

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

	ProductCreatePOSTHandler   gin.HandlerFunc
	ProductRetrieveGETHandler  gin.HandlerFunc
	ProductsListGETHandler     gin.HandlerFunc
	ProductDeleteDELETEHandler gin.HandlerFunc
	ProductUpdatePATCHHandler  gin.HandlerFunc

	StoreProductCreatePOSTHandler          gin.HandlerFunc
	StoreProductRetrieveGETHandler         gin.HandlerFunc
	StoreProductsListGETHandler            gin.HandlerFunc
	StoreProductsWithDetailsListGETHandler gin.HandlerFunc
	StoreProductsByProductIDGETHandler     gin.HandlerFunc
	PromotionalProductsGETHandler          gin.HandlerFunc
	StoreProductUpdatePATCHHandler         gin.HandlerFunc
	StoreProductDeleteDELETEHandler        gin.HandlerFunc
	StoreProductQuantityUpdatePATCHHandler gin.HandlerFunc
	StoreProductStockCheckGETHandler       gin.HandlerFunc

	SaleCreatePOSTHandler               gin.HandlerFunc
	SaleRetrieveGETHandler              gin.HandlerFunc
	SalesByReceiptGETHandler            gin.HandlerFunc
	SalesByUPCGETHandler                gin.HandlerFunc
	SalesListGETHandler                 gin.HandlerFunc
	SalesWithDetailsListGETHandler      gin.HandlerFunc
	SalesWithDetailsByReceiptGETHandler gin.HandlerFunc
	SaleUpdatePATCHHandler              gin.HandlerFunc
	SaleDeleteDELETEHandler             gin.HandlerFunc
	SalesByReceiptDeleteDELETEHandler   gin.HandlerFunc
	ReceiptTotalGETHandler              gin.HandlerFunc
	SalesStatsByProductGETHandler       gin.HandlerFunc
	TopSellingProductsGETHandler        gin.HandlerFunc
}

// Close properly closes the database connection
func (hc *HandlerContainer) Close() error {
	if hc.DB != nil {
		return hc.DB.Close()
	}
	return nil
}

func BuildHandlerContainer(c *config.Config) (*HandlerContainer, error) {
	db, err := sql.Open(c.DB_DRIVER, c.DB_DSN)
	if err != nil {
		return nil, err
	}

	// Configure connection pool
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(5 * time.Minute)

	// Test the connection
	if err := db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	log.Println("Database connection established successfully")

	categoryRepo := repos.NewCategoryRepo(db)
	categoryService := services.NewCategoryService(categoryRepo)

	customerCardRepo := repos.NewCustomerCardRepo(db)
	customerCardService := services.NewCustomerCardService(customerCardRepo)

	employeeRepo := repos.NewEmployeeRepo(db)
	employeeService := services.NewEmployeeService(employeeRepo)

	receiptRepo := repos.NewReceiptRepo(db)
	receiptService := services.NewReceiptService(receiptRepo)

	productRepo := repos.NewProductRepo(db)
	productService := services.NewProductService(productRepo)

	storeProductRepo := repos.NewStoreProductRepo(db)
	storeProductService := services.NewStoreProductService(storeProductRepo)

	saleRepo := repos.NewSaleRepo(db)
	saleService := services.NewSaleService(saleRepo)

	return &HandlerContainer{
		DB: db,

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

		ReceiptCreatePOSTHandler:   handlers.NewReceiptCreatePOSTHandler(receiptService, c),
		ReceiptRetrieveGETHandler:  handlers.NewReceiptRetrieveGETHandler(receiptService),
		ReceiptsListGETHandler:     handlers.NewReceiptsListGETHandler(receiptService),
		ReceiptDeleteDELETEHandler: handlers.NewReceiptDeleteDELETEHandler(receiptService),
		ReceiptUpdatePATCHHandler:  handlers.NewReceiptUpdatePATCHHandler(receiptService, c),

		ProductCreatePOSTHandler:   handlers.NewProductCreatePOSTHandler(productService),
		ProductRetrieveGETHandler:  handlers.NewProductRetrieveGETHandler(productService),
		ProductsListGETHandler:     handlers.NewProductsListGETHandler(productService),
		ProductDeleteDELETEHandler: handlers.NewProductDeleteDELETEHandler(productService),
		ProductUpdatePATCHHandler:  handlers.NewProductUpdatePATCHHandler(productService),

		StoreProductCreatePOSTHandler:          handlers.NewStoreProductCreatePOSTHandler(storeProductService),
		StoreProductRetrieveGETHandler:         handlers.NewStoreProductRetrieveGETHandler(storeProductService),
		StoreProductsListGETHandler:            handlers.NewStoreProductsListGETHandler(storeProductService),
		StoreProductsWithDetailsListGETHandler: handlers.NewStoreProductsWithDetailsListGETHandler(storeProductService),
		StoreProductsByProductIDGETHandler:     handlers.NewStoreProductsByProductIDGETHandler(storeProductService),
		PromotionalProductsGETHandler:          handlers.NewPromotionalProductsGETHandler(storeProductService),
		StoreProductUpdatePATCHHandler:         handlers.NewStoreProductUpdatePATCHHandler(storeProductService),
		StoreProductDeleteDELETEHandler:        handlers.NewStoreProductDeleteDELETEHandler(storeProductService),
		StoreProductQuantityUpdatePATCHHandler: handlers.NewStoreProductQuantityUpdatePATCHHandler(storeProductService),
		StoreProductStockCheckGETHandler:       handlers.NewStoreProductStockCheckGETHandler(storeProductService),

		SaleCreatePOSTHandler:               handlers.NewSaleCreatePOSTHandler(saleService),
		SaleRetrieveGETHandler:              handlers.NewSaleRetrieveGETHandler(saleService),
		SalesByReceiptGETHandler:            handlers.NewSalesByReceiptGETHandler(saleService),
		SalesByUPCGETHandler:                handlers.NewSalesByUPCGETHandler(saleService),
		SalesListGETHandler:                 handlers.NewSalesListGETHandler(saleService),
		SalesWithDetailsListGETHandler:      handlers.NewSalesWithDetailsListGETHandler(saleService),
		SalesWithDetailsByReceiptGETHandler: handlers.NewSalesWithDetailsByReceiptGETHandler(saleService),
		SaleUpdatePATCHHandler:              handlers.NewSaleUpdatePATCHHandler(saleService),
		SaleDeleteDELETEHandler:             handlers.NewSaleDeleteDELETEHandler(saleService),
		SalesByReceiptDeleteDELETEHandler:   handlers.NewSalesByReceiptDeleteDELETEHandler(saleService),
		ReceiptTotalGETHandler:              handlers.NewReceiptTotalGETHandler(saleService),
		SalesStatsByProductGETHandler:       handlers.NewSalesStatsByProductGETHandler(saleService),
		TopSellingProductsGETHandler:        handlers.NewTopSellingProductsGETHandler(saleService),
	}, nil
}
