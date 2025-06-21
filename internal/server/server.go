package server

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/velosypedno/zlagoda/internal/config"
	"github.com/velosypedno/zlagoda/internal/ioc"
	"github.com/velosypedno/zlagoda/internal/middleware"
)

func SetupRoutes(c *ioc.HandlerContainer, cfg *config.Config) *gin.Engine {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	router.POST("/api/login", c.LoginPOSTHandler)
	router.POST("/api/register", c.RegisterPOSTHandler)
	api := router.Group("/api")
	api.Use(middleware.AuthMiddleware(cfg))
	{
		api.GET("/account", c.AccountGETHandler)

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
		api.POST("/employees/with-auth", c.EmployeeCreateWithAuthPOSTHandler)
		api.GET("/employees", c.EmployeesListGETHandler)
		api.GET("/employees/:id", c.EmployeeRetrieveGETHandler)
		api.DELETE("/employees/:id", c.EmployeeDeleteDELETEHandler)
		api.PATCH("/employees/:id", c.EmployeeUpdatePATCHHandler)

		api.POST("/receipts", c.ReceiptCreatePOSTHandler)
		api.POST("/receipts/complete", c.ReceiptCreateCompletePOSTHandler)
		api.GET("/receipts", c.ReceiptsListGETHandler)
		api.GET("/receipts/:receipt_number", c.ReceiptRetrieveGETHandler)
		api.DELETE("/receipts/:receipt_number", c.ReceiptDeleteDELETEHandler)
		api.PATCH("/receipts/:receipt_number", c.ReceiptUpdatePATCHHandler)

		api.POST("/products", c.ProductCreatePOSTHandler)
		api.GET("/products", c.ProductsListGETHandler)
		api.GET("/products/search", c.ProductsByNameGETHandler)
		api.GET("/products/by-category/:category_id", c.ProductsByCategoryGETHandler)
		api.GET("/products/:id", c.ProductRetrieveGETHandler)
		api.DELETE("/products/:id", c.ProductDeleteDELETEHandler)
		api.PATCH("/products/:id", c.ProductUpdatePATCHHandler)

		api.POST("/store-products", c.StoreProductCreatePOSTHandler)
		api.GET("/store-products", c.StoreProductsListGETHandler)
		api.GET("/store-products/details", c.StoreProductsWithDetailsListGETHandler)
		api.GET("/store-products/search", c.StoreProductsByNameGETHandler)
		api.GET("/store-products/by-category/:category_id", c.StoreProductsByCategoryGETHandler)
		api.GET("/store-products/promotional", c.PromotionalProductsGETHandler)
		api.GET("/store-products/by-product/:product_id", c.StoreProductsByProductIDGETHandler)
		api.GET("/store-products/:upc", c.StoreProductRetrieveGETHandler)
		api.DELETE("/store-products/:upc", c.StoreProductDeleteDELETEHandler)
		api.PATCH("/store-products/:upc", c.StoreProductUpdatePATCHHandler)
		api.PATCH("/store-products/:upc/quantity", c.StoreProductQuantityUpdatePATCHHandler)
		api.GET("/store-products/:upc/stock-check", c.StoreProductStockCheckGETHandler)
		api.PATCH("/store-products/:upc/delivery", c.StoreProductDeliveryPATCHHandler)

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

		api.GET("/vlad1", c.Vlad1GETHandler)
		api.GET("/vlad2", c.Vlad1GETHandler)
		api.GET("/arthur1", c.Arthur1GETHandler)
		api.GET("/arthur2", c.Arthur2GETHandler)
		api.GET("/oleksii1", c.Oleksii1GETHandler)
		api.GET("/oleksii2", c.Oleksii2GETHandler)
	}
	return router
}
