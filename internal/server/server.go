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

		api.POST("/customer-card", c.CustomerCardCreatePOSTHandler)
		api.GET("/customer-cards", c.CustomerCardsListGETHandler)
		api.GET("/customer-card/:card_number", c.CustomerCardRetrieveGETHandler)
		api.DELETE("/customer-card/:card_number", c.CustomerCardDeleteDELETEHandler)
		api.PATCH("/customer-card/:card_number", c.CustomerCardUpdatePATCHHandler)
	}
	return router
}
