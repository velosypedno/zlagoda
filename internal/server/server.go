package server

import (
	"github.com/gin-gonic/gin"
	"github.com/velosypedno/zlagoda/internal/ioc"
)

func SetupRoutes(c *ioc.HandlerContainer) *gin.Engine {
	router := gin.Default()
	api := router.Group("/api")
	{
		api.POST("/categories", c.CategoryCreatePOSTHandler)
		api.GET("/categories", c.CategoriesListGETHandler)
		api.GET("/categories/:id", c.CategoryRetrieveGETHandler)
	}
	return router
}
