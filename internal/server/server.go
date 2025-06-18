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
	}
	return router
}
