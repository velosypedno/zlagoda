package server

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/velosypedno/zlagoda/internal/ioc"
)

func SetupRoutes(c *ioc.HandlerContainer) *gin.Engine {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	api := router.Group("/api")
	{
		api.POST("/categories", c.CategoryCreatePOSTHandler)
		api.GET("/categories", c.CategoriesListGETHandler)
		api.GET("/categories/:id", c.CategoryRetrieveGETHandler)
		api.DELETE("/categories/:id", c.CategoryDeleteDELETEHandler)
		api.PATCH("/categories/:id", c.CategoryUpdatePATCHHandler)

		api.POST("/customer-cards", c.CustomerCardCreatePOSTHandler)
		api.GET("/customer-cards", c.CustomerCardsListGETHandler)
		api.GET("/customer-cards/:card_number", c.CustomerCardRetrieveGETHandler)
		api.DELETE("/customer-cards/:card_number", c.CustomerCardDeleteDELETEHandler)
		api.PATCH("/customer-cards/:card_number", c.CustomerCardUpdatePATCHHandler)

		api.POST("/employees", c.EmployeeCreatePOSTHandler)
		api.GET("/employees", c.EmployeesListGETHandler)
		api.GET("/employees/:id", c.EmployeeRetrieveGETHandler)
		api.DELETE("/employees/:id", c.EmployeeDeleteDELETEHandler)
		api.PATCH("/employees/:id", c.EmployeeUpdatePATCHHandler)

		api.POST("/receipts", c.ReceiptCreatePOSTHandler)
		api.GET("/receipts", c.ReceiptsListGETHandler)
		api.GET("/receipts/:receipt_number", c.ReceiptRetrieveGETHandler)
		api.DELETE("/receipts/:receipt_number", c.ReceiptDeleteDELETEHandler)
		api.PATCH("/receipts/:receipt_number", c.ReceiptUpdatePATCHHandler)

		api.POST("/products", c.ProductCreatePOSTHandler)
		api.GET("/products", c.ProductsListGETHandler)
		api.GET("/products/:id", c.ProductRetrieveGETHandler)
		api.DELETE("/products/:id", c.ProductDeleteDELETEHandler)
		api.PATCH("/products/:id", c.ProductUpdatePATCHHandler)

		api.POST("/store-products", c.StoreProductCreatePOSTHandler)
		api.GET("/store-products", c.StoreProductsListGETHandler)
		api.GET("/store-products/details", c.StoreProductsWithDetailsListGETHandler)
		api.GET("/store-products/promotional", c.PromotionalProductsGETHandler)
		api.GET("/store-products/by-product/:product_id", c.StoreProductsByProductIDGETHandler)
		api.GET("/store-products/:upc", c.StoreProductRetrieveGETHandler)
		api.DELETE("/store-products/:upc", c.StoreProductDeleteDELETEHandler)
		api.PATCH("/store-products/:upc", c.StoreProductUpdatePATCHHandler)
		api.PATCH("/store-products/:upc/quantity", c.StoreProductQuantityUpdatePATCHHandler)
		api.GET("/store-products/:upc/stock-check", c.StoreProductStockCheckGETHandler)

		api.POST("/sales", c.SaleCreatePOSTHandler)
		api.GET("/sales", c.SalesListGETHandler)
		api.GET("/sales/details", c.SalesWithDetailsListGETHandler)
		api.GET("/sales/top-products", c.TopSellingProductsGETHandler)
		api.GET("/sales/by-receipt/:receipt_number", c.SalesByReceiptGETHandler)
		api.GET("/sales/by-receipt/:receipt_number/details", c.SalesWithDetailsByReceiptGETHandler)
		api.DELETE("/sales/by-receipt/:receipt_number", c.SalesByReceiptDeleteDELETEHandler)
		api.GET("/sales/by-upc/:upc", c.SalesByUPCGETHandler)
		api.GET("/sales/stats/product/:product_id", c.SalesStatsByProductGETHandler)
		api.GET("/sales/:upc/:receipt_number", c.SaleRetrieveGETHandler)
		api.DELETE("/sales/:upc/:receipt_number", c.SaleDeleteDELETEHandler)
		api.PATCH("/sales/:upc/:receipt_number", c.SaleUpdatePATCHHandler)

		api.GET("/receipts/:receipt_number/total", c.ReceiptTotalGETHandler)
	}
	return router
}
