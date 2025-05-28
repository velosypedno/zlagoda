package server

import (
	"github.com/gin-gonic/gin"
	"github.com/velosypedno/zlagoda/internal/ioc"
)

func SetupRoutes(c *ioc.HandlerContainer) *gin.Engine {
	router := gin.Default()
	api := router.Group("/api")
	{
		api.POST("/category", c.CategoryCreatePOSTHandler)
		api.GET("/category/:id", c.CategoryRetrieveGETHandler)
	}
	return router
}
